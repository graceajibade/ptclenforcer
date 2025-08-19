package dslgen

import (
	"fmt"
	"sort"
	"strings"

	parser "school/project/parser/antlr"
	"github.com/antlr4-go/antlr/v4"
)

type edgeItem struct {
	from, label, to string
	outcome         *bool
}

type Builder struct {
	parser.BaseTSgenVisitor
	FSM *FSMDef
	doc string

	// config for generation
	pkgShort   string
	importPath string
	protocol   string
	fileName   string

	// collected for generation
	states map[string]struct{}
	labels map[string]struct{}
	s0     string
	f      string
	edges  [][3]string // (from, label, to)
	edgesEx []edgeItem  // (from, label, to) with outcome

	// raw logic (collected then folded)
	rawList []rawItem

	// to reconstruct exact input text
	ts *antlr.CommonTokenStream
}

type rawItem struct {
	from, label, name string
	inputs, outputs   []string
}

func NewBuilder() *Builder {
	return &Builder{
		FSM: &FSMDef{
			States:      make(map[string]bool),
			Finals:      make(map[string]bool),
			Transitions: []DSLTransition{},
		},
		states: make(map[string]struct{}),
		labels: make(map[string]struct{}),
	}
}

// Ensure traversal uses our visitor
func (b *Builder) VisitChildren(node antlr.RuleNode) interface{} {
	var res interface{}
	for i := 0; i < node.GetChildCount(); i++ {
		if pt, ok := node.GetChild(i).(antlr.ParseTree); ok {
			res = pt.Accept(b)
		}
	}
	return res
}

// program : statement* EOF ;
func (b *Builder) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, s := range ctx.AllStatement() {
		s.Accept(b)
	}
	return nil
}

// statement : stateDecl | transitionDecl | docComment | pkgDecl | importDecl |
//             protocolDecl | fileDecl | rawDecl | s0Decl | fDecl ;
func (b *Builder) VisitStatement(ctx *parser.StatementContext) interface{} {
	if x := ctx.StateDecl(); x != nil { return x.Accept(b) }
	if x := ctx.TransitionDecl(); x != nil { return x.Accept(b) }
	if x := ctx.DocComment(); x != nil { return x.Accept(b) }
	if x := ctx.PkgDecl(); x != nil { return x.Accept(b) }
	if x := ctx.ImportDecl(); x != nil { return x.Accept(b) }
	if x := ctx.ProtocolDecl(); x != nil { return x.Accept(b) }
	if x := ctx.FileDecl(); x != nil { return x.Accept(b) }
	if x := ctx.RawDecl(); x != nil { return x.Accept(b) }
	if x := ctx.S0Decl(); x != nil { return x.Accept(b) }
	if x := ctx.FDecl(); x != nil { return x.Accept(b) }
	// safe fallback
	return b.VisitChildren(ctx)
}

// stateDecl : 'state' from=ID ('final')? ;
func (b *Builder) VisitStateDecl(ctx *parser.StateDeclContext) interface{} {
	// labeled token if present, else single ID()
	var id string
	type withFrom interface{ GetFrom() antlr.Token }
	if w, ok := any(ctx).(withFrom); ok && w.GetFrom() != nil {
		id = w.GetFrom().GetText()
	} else if tok := ctx.ID(); tok != nil {
		id = tok.GetText()
	}
	if id == "" { return nil }

	b.FSM.States[id] = true
	b.states[id] = struct{}{}

	// detect "final"
	final := false
	if getter, ok := any(ctx).(interface{ GetFinal() antlr.Token }); ok && getter.GetFinal() != nil {
		final = true
	} else if containsWord(ctx.GetText(), "final") {
		final = true
	}
	if final {
		b.FSM.Finals[id] = true
	}
	return nil
}

// transitionDecl : from=ID '->' label=ID ':' method=ID '(' tArgs? ')' '->' to=ID ( WHEN outcome=(TRUE|FALSE) )?
func (b *Builder) VisitTransitionDecl(ctx *parser.TransitionDeclContext) interface{} {
    fromTok := ctx.GetFrom()
    labelTok := ctx.GetLabel()
    methodTok := ctx.GetMethod()
    toTok := ctx.GetTo()
    if fromTok == nil || labelTok == nil || methodTok == nil || toTok == nil {
        return nil
    }
    from  := fromTok.GetText()
    label := labelTok.GetText()
    method:= methodTok.GetText()
    to    := toTok.GetText()

    // preview AST (keeps your original test intact)
    b.FSM.Transitions = append(b.FSM.Transitions, DSLTransition{
        From: from, Label: label, Method: method, To: to, Doc: b.doc,
    })

    // for generation
    b.states[from] = struct{}{}
    b.states[to]   = struct{}{}
    b.labels[label]= struct{}{}
    b.edges = append(b.edges, [3]string{from, label, to})

    // explicit outcome (optional)
    var outPtr *bool
    if o := ctx.GetOutcome(); o != nil {
        v := (o.GetText() == "true")
        outPtr = &v
    }
    b.edgesEx = append(b.edgesEx, edgeItem{
        from: from, label: label, to: to, outcome: outPtr,
    })

    b.doc = "" // consume any doc
    return nil
}

// docComment : COMMENT ;
func (b *Builder) VisitDocComment(ctx *parser.DocCommentContext) interface{} {
	if t := ctx.COMMENT(); t != nil {
		b.doc = t.GetText()
	}
	return nil
}

// --- config visitors ---

func (b *Builder) VisitPkgDecl(ctx *parser.PkgDeclContext) interface{} {
	if id := ctx.ID(); id != nil {
		b.pkgShort = id.GetText()
	}
	return nil
}

func (b *Builder) VisitImportDecl(ctx *parser.ImportDeclContext) interface{} {
	if s := ctx.STRING(); s != nil {
		b.importPath = strings.Trim(s.GetText(), `"`)
	}
	return nil
}

func (b *Builder) VisitProtocolDecl(ctx *parser.ProtocolDeclContext) interface{} {
	if id := ctx.ID(); id != nil {
		b.protocol = id.GetText()
	}
	return nil
}

func (b *Builder) VisitFileDecl(ctx *parser.FileDeclContext) interface{} {
	if f := ctx.FILENAME(); f != nil {
		b.fileName = f.GetText()
	}
	return nil
}

// rawDecl : 'raw' from=ID label=ID 'name' mname=ID '(' rawInputs? ')' 'returns' rawOutputs
func (b *Builder) VisitRawDecl(ctx *parser.RawDeclContext) interface{} {
    from := ctx.GetFrom().GetText()
    label := ctx.GetLabel().GetText()
    name := ctx.GetMname().GetText()

    var inputs []string
    if ri := ctx.RawInputs(); ri != nil {
        // Reconstruct exact text (preserves spaces and ellipsis)
        start := ri.GetStart().GetTokenIndex()
        stop  := ri.GetStop().GetTokenIndex()
        seg := b.ts.GetTextFromInterval(antlr.NewInterval(start, stop))
        // seg is like: "v any" or "args ...string" or "t json.Token, n int"
        for _, part := range strings.Split(seg, ",") {
            s := strings.TrimSpace(part)
            if s != "" {
                inputs = append(inputs, s)
            }
        }
    }

    var outputs []string
    // Keep outputs logic, but make it robust to dotted types as well
    if ro := ctx.RawOutputs(); ro != nil {
        s := b.ts.GetTextFromInterval(antlr.NewInterval(ro.GetStart().GetTokenIndex(), ro.GetStop().GetTokenIndex()))
        for _, part := range strings.Split(s, ",") {
            p := strings.TrimSpace(part)
            if p != "" {
                outputs = append(outputs, p)
            }
        }
    }

    b.rawList = append(b.rawList, rawItem{
        from: from, label: label, name: name, inputs: inputs, outputs: outputs,
    })
    b.states[from] = struct{}{}
    b.labels[label] = struct{}{}
    return nil
}

// s0Decl : 's0' ID ;
func (b *Builder) VisitS0Decl(ctx *parser.S0DeclContext) interface{} {
	if id := ctx.ID(); id != nil {
		b.s0 = id.GetText()
		b.states[b.s0] = struct{}{}
	}
	return nil
}

// fDecl : 'fstate' ID ;
func (b *Builder) VisitFDecl(ctx *parser.FDeclContext) interface{} {
	if id := ctx.ID(); id != nil {
		b.f = id.GetText()
		b.states[b.f] = struct{}{}
		b.FSM.Finals[b.f] = true
	}
	return nil
}

// ---------- public: parser-only entrypoint (keeps your original test) ----------
func ParseANTLRDSL(input string) (*FSMDef, error) {
	is := antlr.NewInputStream(input)
	lex := parser.NewTSgenLexer(is)
	tok := antlr.NewCommonTokenStream(lex, 0)
	p := parser.NewTSgenParser(tok)
	p.BuildParseTrees = true

	tree := p.Program()
	b := NewBuilder()
	b.ts = tok 
	tree.Accept(b)
	return b.FSM, nil
}

// ---------- public: produce a generation spec ----------
func (b *Builder) Spec() GenSpecDTO {
	// fold raw list to map[[2]string]RawSigDTO
	raw := make(map[[2]string]RawSigDTO)
	for _, it := range b.rawList {
		raw[[2]string{it.from, it.label}] = RawSigDTO{
			Name:    it.name,
			Inputs:  it.inputs,
			Outputs: it.outputs,
		}
	}
	// stable ordering for States/Labels
	sts := make([]string, 0, len(b.states))
	for s := range b.states { sts = append(sts, s) }
	sort.Strings(sts)
	lbls := make([]string, 0, len(b.labels))
	for l := range b.labels { lbls = append(lbls, l) }
	sort.Strings(lbls)

	var edgesEx []EdgeDTO
	for _, e := range b.edgesEx {
		edgesEx = append(edgesEx, EdgeDTO{
			From: e.from, Label: e.label, To: e.to, Outcome: e.outcome,
		})
	}

	return GenSpecDTO{
		Pkg:        b.pkgShort,
		ImportPath: b.importPath,
		Protocol:   b.protocol,
		FileName:   b.fileName,
		S0:         b.s0,
		F:          b.f,
		States:     sts,
		Labels:     lbls,
		Edges:      b.edges,
		EdgesEx:    edgesEx,
		Raw:        raw,
	}
}

// ---------- convenience: parse + generate in one shot ----------
func GenerateFromDSL(input string) (string, error) {
	is := antlr.NewInputStream(input)
	lex := parser.NewTSgenLexer(is)
	tok := antlr.NewCommonTokenStream(lex, 0)
	p := parser.NewTSgenParser(tok)
	p.BuildParseTrees = true

	tree := p.Program()
	b := NewBuilder()
	b.ts = tok 
	tree.Accept(b)

	spec := b.Spec()
	// simple validation
	if spec.Pkg == "" || spec.ImportPath == "" || spec.Protocol == "" || spec.FileName == "" {
		return "", fmt.Errorf("DSL config missing: require pkg, import, protocol, and file")
	}
	if spec.S0 == "" && len(spec.States) > 0 {
		spec.S0 = spec.States[0]
	}
	if spec.F == "" && len(spec.States) > 0 {
		spec.F = spec.States[len(spec.States)-1]
	}

	return GenerateFromSpec(spec)
}

// ---------- helpers ----------
func collectIDs(t antlr.Tree) []antlr.TerminalNode {
	var out []antlr.TerminalNode
	var walk func(antlr.Tree)
	walk = func(n antlr.Tree) {
		switch nn := n.(type) {
		case antlr.TerminalNode:
			if nn.GetSymbol() != nil && nn.GetSymbol().GetText() != "" {
				out = append(out, nn)
			}
		case antlr.RuleNode:
			for i := 0; i < nn.GetChildCount(); i++ {
				walk(nn.GetChild(i))
			}
		}
	}
	walk(t)
	return out
}

func containsWord(s, w string) bool {
	if len(s) < len(w) { return false }
	for i := 0; i+len(w) <= len(s); i++ {
		if s[i:i+len(w)] == w { return true }
	}
	return false
}

func compactSpaces(s string) string {
	var b strings.Builder
	space := false
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			space = true
			continue
		}
		if space && b.Len() > 0 {
			b.WriteByte(' ')
		}
		space = false
		b.WriteRune(r)
	}
	return b.String()
}
