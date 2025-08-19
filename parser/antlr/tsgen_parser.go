// Code generated from antlr/TSgen.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TSgen

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TSgenParser struct {
	*antlr.BaseParser
}

var TSgenParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tsgenParserInit() {
	staticData := &TSgenParserStaticData
	staticData.LiteralNames = []string{
		"", "'state'", "'final'", "'->'", "':'", "'('", "')'", "','", "'pkg'",
		"'import'", "'protocol'", "'file'", "'raw'", "'name'", "'returns'",
		"'.'", "'s0'", "'fstate'", "'['", "']'", "'*'", "'...'", "'when'", "'true'",
		"'false'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "LBRACK", "RBRACK", "STAR", "ELLIPSIS", "WHEN", "TRUE", "FALSE",
		"COMMENT", "FILENAME", "ID", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "stateDecl", "transitionDecl", "tArgs", "tParam",
		"docComment", "pkgDecl", "importDecl", "protocolDecl", "fileDecl", "rawDecl",
		"rawInputs", "rawParam", "rawOutputs", "typeExpr", "qualID", "s0Decl",
		"fDecl",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 173, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 5, 0, 40, 8, 0, 10, 0,
		12, 0, 43, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 57, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 62, 8, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 71, 8, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 78, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 83, 8, 4, 10, 4, 12,
		4, 86, 9, 4, 1, 5, 1, 5, 1, 5, 3, 5, 91, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 114, 8, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 5, 12, 123, 8, 12, 10, 12, 12, 12, 126,
		9, 12, 1, 13, 1, 13, 1, 13, 3, 13, 131, 8, 13, 1, 14, 1, 14, 1, 14, 5,
		14, 136, 8, 14, 10, 14, 12, 14, 139, 9, 14, 1, 15, 3, 15, 142, 8, 15, 1,
		15, 5, 15, 145, 8, 15, 10, 15, 12, 15, 148, 9, 15, 1, 15, 1, 15, 5, 15,
		152, 8, 15, 10, 15, 12, 15, 155, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 5, 16, 162, 8, 16, 10, 16, 12, 16, 165, 9, 16, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 0, 0, 19, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 0, 1, 1, 0, 23, 24, 176, 0, 41, 1,
		0, 0, 0, 2, 56, 1, 0, 0, 0, 4, 58, 1, 0, 0, 0, 6, 63, 1, 0, 0, 0, 8, 79,
		1, 0, 0, 0, 10, 90, 1, 0, 0, 0, 12, 92, 1, 0, 0, 0, 14, 94, 1, 0, 0, 0,
		16, 97, 1, 0, 0, 0, 18, 100, 1, 0, 0, 0, 20, 103, 1, 0, 0, 0, 22, 106,
		1, 0, 0, 0, 24, 119, 1, 0, 0, 0, 26, 130, 1, 0, 0, 0, 28, 132, 1, 0, 0,
		0, 30, 141, 1, 0, 0, 0, 32, 158, 1, 0, 0, 0, 34, 166, 1, 0, 0, 0, 36, 169,
		1, 0, 0, 0, 38, 40, 3, 2, 1, 0, 39, 38, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0,
		41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 44, 1, 0, 0, 0, 43, 41, 1,
		0, 0, 0, 44, 45, 5, 0, 0, 1, 45, 1, 1, 0, 0, 0, 46, 57, 3, 4, 2, 0, 47,
		57, 3, 6, 3, 0, 48, 57, 3, 12, 6, 0, 49, 57, 3, 14, 7, 0, 50, 57, 3, 16,
		8, 0, 51, 57, 3, 18, 9, 0, 52, 57, 3, 20, 10, 0, 53, 57, 3, 22, 11, 0,
		54, 57, 3, 34, 17, 0, 55, 57, 3, 36, 18, 0, 56, 46, 1, 0, 0, 0, 56, 47,
		1, 0, 0, 0, 56, 48, 1, 0, 0, 0, 56, 49, 1, 0, 0, 0, 56, 50, 1, 0, 0, 0,
		56, 51, 1, 0, 0, 0, 56, 52, 1, 0, 0, 0, 56, 53, 1, 0, 0, 0, 56, 54, 1,
		0, 0, 0, 56, 55, 1, 0, 0, 0, 57, 3, 1, 0, 0, 0, 58, 59, 5, 1, 0, 0, 59,
		61, 5, 27, 0, 0, 60, 62, 5, 2, 0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0,
		0, 0, 62, 5, 1, 0, 0, 0, 63, 64, 5, 27, 0, 0, 64, 65, 5, 3, 0, 0, 65, 66,
		5, 27, 0, 0, 66, 67, 5, 4, 0, 0, 67, 68, 5, 27, 0, 0, 68, 70, 5, 5, 0,
		0, 69, 71, 3, 8, 4, 0, 70, 69, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72,
		1, 0, 0, 0, 72, 73, 5, 6, 0, 0, 73, 74, 5, 3, 0, 0, 74, 77, 5, 27, 0, 0,
		75, 76, 5, 22, 0, 0, 76, 78, 7, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1,
		0, 0, 0, 78, 7, 1, 0, 0, 0, 79, 84, 3, 10, 5, 0, 80, 81, 5, 7, 0, 0, 81,
		83, 3, 10, 5, 0, 82, 80, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1, 0,
		0, 0, 84, 85, 1, 0, 0, 0, 85, 9, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 88,
		5, 27, 0, 0, 88, 91, 3, 30, 15, 0, 89, 91, 3, 30, 15, 0, 90, 87, 1, 0,
		0, 0, 90, 89, 1, 0, 0, 0, 91, 11, 1, 0, 0, 0, 92, 93, 5, 25, 0, 0, 93,
		13, 1, 0, 0, 0, 94, 95, 5, 8, 0, 0, 95, 96, 5, 27, 0, 0, 96, 15, 1, 0,
		0, 0, 97, 98, 5, 9, 0, 0, 98, 99, 5, 28, 0, 0, 99, 17, 1, 0, 0, 0, 100,
		101, 5, 10, 0, 0, 101, 102, 5, 27, 0, 0, 102, 19, 1, 0, 0, 0, 103, 104,
		5, 11, 0, 0, 104, 105, 5, 26, 0, 0, 105, 21, 1, 0, 0, 0, 106, 107, 5, 12,
		0, 0, 107, 108, 5, 27, 0, 0, 108, 109, 5, 27, 0, 0, 109, 110, 5, 13, 0,
		0, 110, 111, 5, 27, 0, 0, 111, 113, 5, 5, 0, 0, 112, 114, 3, 24, 12, 0,
		113, 112, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115,
		116, 5, 6, 0, 0, 116, 117, 5, 14, 0, 0, 117, 118, 3, 28, 14, 0, 118, 23,
		1, 0, 0, 0, 119, 124, 3, 26, 13, 0, 120, 121, 5, 7, 0, 0, 121, 123, 3,
		26, 13, 0, 122, 120, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0,
		0, 0, 124, 125, 1, 0, 0, 0, 125, 25, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0,
		127, 128, 5, 27, 0, 0, 128, 131, 3, 30, 15, 0, 129, 131, 3, 30, 15, 0,
		130, 127, 1, 0, 0, 0, 130, 129, 1, 0, 0, 0, 131, 27, 1, 0, 0, 0, 132, 137,
		3, 30, 15, 0, 133, 134, 5, 7, 0, 0, 134, 136, 3, 30, 15, 0, 135, 133, 1,
		0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0,
		0, 138, 29, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 142, 5, 21, 0, 0, 141,
		140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 146, 1, 0, 0, 0, 143, 145,
		5, 20, 0, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0,
		0, 0, 146, 147, 1, 0, 0, 0, 147, 153, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0,
		149, 150, 5, 18, 0, 0, 150, 152, 5, 19, 0, 0, 151, 149, 1, 0, 0, 0, 152,
		155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 156,
		1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157, 3, 32, 16, 0, 157, 31, 1, 0,
		0, 0, 158, 163, 5, 27, 0, 0, 159, 160, 5, 15, 0, 0, 160, 162, 5, 27, 0,
		0, 161, 159, 1, 0, 0, 0, 162, 165, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163,
		164, 1, 0, 0, 0, 164, 33, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 166, 167, 5,
		16, 0, 0, 167, 168, 5, 27, 0, 0, 168, 35, 1, 0, 0, 0, 169, 170, 5, 17,
		0, 0, 170, 171, 5, 27, 0, 0, 171, 37, 1, 0, 0, 0, 15, 41, 56, 61, 70, 77,
		84, 90, 113, 124, 130, 137, 141, 146, 153, 163,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TSgenParserInit initializes any static state used to implement TSgenParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTSgenParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TSgenParserInit() {
	staticData := &TSgenParserStaticData
	staticData.once.Do(tsgenParserInit)
}

// NewTSgenParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTSgenParser(input antlr.TokenStream) *TSgenParser {
	TSgenParserInit()
	this := new(TSgenParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TSgenParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TSgen.g4"

	return this
}

// TSgenParser tokens.
const (
	TSgenParserEOF      = antlr.TokenEOF
	TSgenParserT__0     = 1
	TSgenParserT__1     = 2
	TSgenParserT__2     = 3
	TSgenParserT__3     = 4
	TSgenParserT__4     = 5
	TSgenParserT__5     = 6
	TSgenParserT__6     = 7
	TSgenParserT__7     = 8
	TSgenParserT__8     = 9
	TSgenParserT__9     = 10
	TSgenParserT__10    = 11
	TSgenParserT__11    = 12
	TSgenParserT__12    = 13
	TSgenParserT__13    = 14
	TSgenParserT__14    = 15
	TSgenParserT__15    = 16
	TSgenParserT__16    = 17
	TSgenParserLBRACK   = 18
	TSgenParserRBRACK   = 19
	TSgenParserSTAR     = 20
	TSgenParserELLIPSIS = 21
	TSgenParserWHEN     = 22
	TSgenParserTRUE     = 23
	TSgenParserFALSE    = 24
	TSgenParserCOMMENT  = 25
	TSgenParserFILENAME = 26
	TSgenParserID       = 27
	TSgenParserSTRING   = 28
	TSgenParserWS       = 29
)

// TSgenParser rules.
const (
	TSgenParserRULE_program        = 0
	TSgenParserRULE_statement      = 1
	TSgenParserRULE_stateDecl      = 2
	TSgenParserRULE_transitionDecl = 3
	TSgenParserRULE_tArgs          = 4
	TSgenParserRULE_tParam         = 5
	TSgenParserRULE_docComment     = 6
	TSgenParserRULE_pkgDecl        = 7
	TSgenParserRULE_importDecl     = 8
	TSgenParserRULE_protocolDecl   = 9
	TSgenParserRULE_fileDecl       = 10
	TSgenParserRULE_rawDecl        = 11
	TSgenParserRULE_rawInputs      = 12
	TSgenParserRULE_rawParam       = 13
	TSgenParserRULE_rawOutputs     = 14
	TSgenParserRULE_typeExpr       = 15
	TSgenParserRULE_qualID         = 16
	TSgenParserRULE_s0Decl         = 17
	TSgenParserRULE_fDecl          = 18
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(TSgenParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TSgenParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&167976706) != 0 {
		{
			p.SetState(38)
			p.Statement()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
		p.Match(TSgenParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StateDecl() IStateDeclContext
	TransitionDecl() ITransitionDeclContext
	DocComment() IDocCommentContext
	PkgDecl() IPkgDeclContext
	ImportDecl() IImportDeclContext
	ProtocolDecl() IProtocolDeclContext
	FileDecl() IFileDeclContext
	RawDecl() IRawDeclContext
	S0Decl() IS0DeclContext
	FDecl() IFDeclContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) StateDecl() IStateDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateDeclContext)
}

func (s *StatementContext) TransitionDecl() ITransitionDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransitionDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransitionDeclContext)
}

func (s *StatementContext) DocComment() IDocCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDocCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDocCommentContext)
}

func (s *StatementContext) PkgDecl() IPkgDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPkgDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPkgDeclContext)
}

func (s *StatementContext) ImportDecl() IImportDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *StatementContext) ProtocolDecl() IProtocolDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProtocolDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProtocolDeclContext)
}

func (s *StatementContext) FileDecl() IFileDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFileDeclContext)
}

func (s *StatementContext) RawDecl() IRawDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRawDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRawDeclContext)
}

func (s *StatementContext) S0Decl() IS0DeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IS0DeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IS0DeclContext)
}

func (s *StatementContext) FDecl() IFDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFDeclContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TSgenParserRULE_statement)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TSgenParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.StateDecl()
		}

	case TSgenParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.TransitionDecl()
		}

	case TSgenParserCOMMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
			p.DocComment()
		}

	case TSgenParserT__7:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.PkgDecl()
		}

	case TSgenParserT__8:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(50)
			p.ImportDecl()
		}

	case TSgenParserT__9:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(51)
			p.ProtocolDecl()
		}

	case TSgenParserT__10:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(52)
			p.FileDecl()
		}

	case TSgenParserT__11:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(53)
			p.RawDecl()
		}

	case TSgenParserT__15:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(54)
			p.S0Decl()
		}

	case TSgenParserT__16:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(55)
			p.FDecl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStateDeclContext is an interface to support dynamic dispatch.
type IStateDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFrom returns the from token.
	GetFrom() antlr.Token

	// SetFrom sets the from token.
	SetFrom(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode

	// IsStateDeclContext differentiates from other interfaces.
	IsStateDeclContext()
}

type StateDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	from   antlr.Token
}

func NewEmptyStateDeclContext() *StateDeclContext {
	var p = new(StateDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_stateDecl
	return p
}

func InitEmptyStateDeclContext(p *StateDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_stateDecl
}

func (*StateDeclContext) IsStateDeclContext() {}

func NewStateDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateDeclContext {
	var p = new(StateDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_stateDecl

	return p
}

func (s *StateDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StateDeclContext) GetFrom() antlr.Token { return s.from }

func (s *StateDeclContext) SetFrom(v antlr.Token) { s.from = v }

func (s *StateDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TSgenParserID, 0)
}

func (s *StateDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitStateDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) StateDecl() (localctx IStateDeclContext) {
	localctx = NewStateDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TSgenParserRULE_stateDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(TSgenParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)

		var _m = p.Match(TSgenParserID)

		localctx.(*StateDeclContext).from = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSgenParserT__1 {
		{
			p.SetState(60)
			p.Match(TSgenParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITransitionDeclContext is an interface to support dynamic dispatch.
type ITransitionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFrom returns the from token.
	GetFrom() antlr.Token

	// GetLabel returns the label token.
	GetLabel() antlr.Token

	// GetMethod returns the method token.
	GetMethod() antlr.Token

	// GetTo returns the to token.
	GetTo() antlr.Token

	// GetOutcome returns the outcome token.
	GetOutcome() antlr.Token

	// SetFrom sets the from token.
	SetFrom(antlr.Token)

	// SetLabel sets the label token.
	SetLabel(antlr.Token)

	// SetMethod sets the method token.
	SetMethod(antlr.Token)

	// SetTo sets the to token.
	SetTo(antlr.Token)

	// SetOutcome sets the outcome token.
	SetOutcome(antlr.Token)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	TArgs() ITArgsContext
	WHEN() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsTransitionDeclContext differentiates from other interfaces.
	IsTransitionDeclContext()
}

type TransitionDeclContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	from    antlr.Token
	label   antlr.Token
	method  antlr.Token
	to      antlr.Token
	outcome antlr.Token
}

func NewEmptyTransitionDeclContext() *TransitionDeclContext {
	var p = new(TransitionDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_transitionDecl
	return p
}

func InitEmptyTransitionDeclContext(p *TransitionDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_transitionDecl
}

func (*TransitionDeclContext) IsTransitionDeclContext() {}

func NewTransitionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransitionDeclContext {
	var p = new(TransitionDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_transitionDecl

	return p
}

func (s *TransitionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TransitionDeclContext) GetFrom() antlr.Token { return s.from }

func (s *TransitionDeclContext) GetLabel() antlr.Token { return s.label }

func (s *TransitionDeclContext) GetMethod() antlr.Token { return s.method }

func (s *TransitionDeclContext) GetTo() antlr.Token { return s.to }

func (s *TransitionDeclContext) GetOutcome() antlr.Token { return s.outcome }

func (s *TransitionDeclContext) SetFrom(v antlr.Token) { s.from = v }

func (s *TransitionDeclContext) SetLabel(v antlr.Token) { s.label = v }

func (s *TransitionDeclContext) SetMethod(v antlr.Token) { s.method = v }

func (s *TransitionDeclContext) SetTo(v antlr.Token) { s.to = v }

func (s *TransitionDeclContext) SetOutcome(v antlr.Token) { s.outcome = v }

func (s *TransitionDeclContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TSgenParserID)
}

func (s *TransitionDeclContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TSgenParserID, i)
}

func (s *TransitionDeclContext) TArgs() ITArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITArgsContext)
}

func (s *TransitionDeclContext) WHEN() antlr.TerminalNode {
	return s.GetToken(TSgenParserWHEN, 0)
}

func (s *TransitionDeclContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TSgenParserTRUE, 0)
}

func (s *TransitionDeclContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TSgenParserFALSE, 0)
}

func (s *TransitionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransitionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransitionDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitTransitionDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) TransitionDecl() (localctx ITransitionDeclContext) {
	localctx = NewTransitionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TSgenParserRULE_transitionDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)

		var _m = p.Match(TSgenParserID)

		localctx.(*TransitionDeclContext).from = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(TSgenParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)

		var _m = p.Match(TSgenParserID)

		localctx.(*TransitionDeclContext).label = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(TSgenParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)

		var _m = p.Match(TSgenParserID)

		localctx.(*TransitionDeclContext).method = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(TSgenParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137625600) != 0 {
		{
			p.SetState(69)
			p.TArgs()
		}

	}
	{
		p.SetState(72)
		p.Match(TSgenParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(TSgenParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)

		var _m = p.Match(TSgenParserID)

		localctx.(*TransitionDeclContext).to = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSgenParserWHEN {
		{
			p.SetState(75)
			p.Match(TSgenParserWHEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*TransitionDeclContext).outcome = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == TSgenParserTRUE || _la == TSgenParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*TransitionDeclContext).outcome = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITArgsContext is an interface to support dynamic dispatch.
type ITArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTParam() []ITParamContext
	TParam(i int) ITParamContext

	// IsTArgsContext differentiates from other interfaces.
	IsTArgsContext()
}

type TArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTArgsContext() *TArgsContext {
	var p = new(TArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_tArgs
	return p
}

func InitEmptyTArgsContext(p *TArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_tArgs
}

func (*TArgsContext) IsTArgsContext() {}

func NewTArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TArgsContext {
	var p = new(TArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_tArgs

	return p
}

func (s *TArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *TArgsContext) AllTParam() []ITParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITParamContext); ok {
			len++
		}
	}

	tst := make([]ITParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITParamContext); ok {
			tst[i] = t.(ITParamContext)
			i++
		}
	}

	return tst
}

func (s *TArgsContext) TParam(i int) ITParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITParamContext)
}

func (s *TArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitTArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) TArgs() (localctx ITArgsContext) {
	localctx = NewTArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TSgenParserRULE_tArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.TParam()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSgenParserT__6 {
		{
			p.SetState(80)
			p.Match(TSgenParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.TParam()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITParamContext is an interface to support dynamic dispatch.
type ITParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TypeExpr() ITypeExprContext

	// IsTParamContext differentiates from other interfaces.
	IsTParamContext()
}

type TParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTParamContext() *TParamContext {
	var p = new(TParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_tParam
	return p
}

func InitEmptyTParamContext(p *TParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_tParam
}

func (*TParamContext) IsTParamContext() {}

func NewTParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TParamContext {
	var p = new(TParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_tParam

	return p
}

func (s *TParamContext) GetParser() antlr.Parser { return s.parser }

func (s *TParamContext) ID() antlr.TerminalNode {
	return s.GetToken(TSgenParserID, 0)
}

func (s *TParamContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitTParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) TParam() (localctx ITParamContext) {
	localctx = NewTParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TSgenParserRULE_tParam)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(TSgenParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.TypeExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.TypeExpr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDocCommentContext is an interface to support dynamic dispatch.
type IDocCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENT() antlr.TerminalNode

	// IsDocCommentContext differentiates from other interfaces.
	IsDocCommentContext()
}

type DocCommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocCommentContext() *DocCommentContext {
	var p = new(DocCommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_docComment
	return p
}

func InitEmptyDocCommentContext(p *DocCommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_docComment
}

func (*DocCommentContext) IsDocCommentContext() {}

func NewDocCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocCommentContext {
	var p = new(DocCommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_docComment

	return p
}

func (s *DocCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocCommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(TSgenParserCOMMENT, 0)
}

func (s *DocCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocCommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitDocComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) DocComment() (localctx IDocCommentContext) {
	localctx = NewDocCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TSgenParserRULE_docComment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(TSgenParserCOMMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPkgDeclContext is an interface to support dynamic dispatch.
type IPkgDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsPkgDeclContext differentiates from other interfaces.
	IsPkgDeclContext()
}

type PkgDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPkgDeclContext() *PkgDeclContext {
	var p = new(PkgDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_pkgDecl
	return p
}

func InitEmptyPkgDeclContext(p *PkgDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_pkgDecl
}

func (*PkgDeclContext) IsPkgDeclContext() {}

func NewPkgDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PkgDeclContext {
	var p = new(PkgDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_pkgDecl

	return p
}

func (s *PkgDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PkgDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TSgenParserID, 0)
}

func (s *PkgDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PkgDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PkgDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitPkgDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) PkgDecl() (localctx IPkgDeclContext) {
	localctx = NewPkgDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TSgenParserRULE_pkgDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(TSgenParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(TSgenParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_importDecl
	return p
}

func InitEmptyImportDeclContext(p *ImportDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_importDecl
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) STRING() antlr.TerminalNode {
	return s.GetToken(TSgenParserSTRING, 0)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitImportDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) ImportDecl() (localctx IImportDeclContext) {
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TSgenParserRULE_importDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(TSgenParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Match(TSgenParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProtocolDeclContext is an interface to support dynamic dispatch.
type IProtocolDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsProtocolDeclContext differentiates from other interfaces.
	IsProtocolDeclContext()
}

type ProtocolDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtocolDeclContext() *ProtocolDeclContext {
	var p = new(ProtocolDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_protocolDecl
	return p
}

func InitEmptyProtocolDeclContext(p *ProtocolDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_protocolDecl
}

func (*ProtocolDeclContext) IsProtocolDeclContext() {}

func NewProtocolDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtocolDeclContext {
	var p = new(ProtocolDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_protocolDecl

	return p
}

func (s *ProtocolDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtocolDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TSgenParserID, 0)
}

func (s *ProtocolDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtocolDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtocolDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitProtocolDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) ProtocolDecl() (localctx IProtocolDeclContext) {
	localctx = NewProtocolDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TSgenParserRULE_protocolDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(TSgenParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(TSgenParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFileDeclContext is an interface to support dynamic dispatch.
type IFileDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FILENAME() antlr.TerminalNode

	// IsFileDeclContext differentiates from other interfaces.
	IsFileDeclContext()
}

type FileDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileDeclContext() *FileDeclContext {
	var p = new(FileDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_fileDecl
	return p
}

func InitEmptyFileDeclContext(p *FileDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_fileDecl
}

func (*FileDeclContext) IsFileDeclContext() {}

func NewFileDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileDeclContext {
	var p = new(FileDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_fileDecl

	return p
}

func (s *FileDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FileDeclContext) FILENAME() antlr.TerminalNode {
	return s.GetToken(TSgenParserFILENAME, 0)
}

func (s *FileDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitFileDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) FileDecl() (localctx IFileDeclContext) {
	localctx = NewFileDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TSgenParserRULE_fileDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(TSgenParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(TSgenParserFILENAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRawDeclContext is an interface to support dynamic dispatch.
type IRawDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFrom returns the from token.
	GetFrom() antlr.Token

	// GetLabel returns the label token.
	GetLabel() antlr.Token

	// GetMname returns the mname token.
	GetMname() antlr.Token

	// SetFrom sets the from token.
	SetFrom(antlr.Token)

	// SetLabel sets the label token.
	SetLabel(antlr.Token)

	// SetMname sets the mname token.
	SetMname(antlr.Token)

	// Getter signatures
	RawOutputs() IRawOutputsContext
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	RawInputs() IRawInputsContext

	// IsRawDeclContext differentiates from other interfaces.
	IsRawDeclContext()
}

type RawDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	from   antlr.Token
	label  antlr.Token
	mname  antlr.Token
}

func NewEmptyRawDeclContext() *RawDeclContext {
	var p = new(RawDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_rawDecl
	return p
}

func InitEmptyRawDeclContext(p *RawDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_rawDecl
}

func (*RawDeclContext) IsRawDeclContext() {}

func NewRawDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawDeclContext {
	var p = new(RawDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_rawDecl

	return p
}

func (s *RawDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *RawDeclContext) GetFrom() antlr.Token { return s.from }

func (s *RawDeclContext) GetLabel() antlr.Token { return s.label }

func (s *RawDeclContext) GetMname() antlr.Token { return s.mname }

func (s *RawDeclContext) SetFrom(v antlr.Token) { s.from = v }

func (s *RawDeclContext) SetLabel(v antlr.Token) { s.label = v }

func (s *RawDeclContext) SetMname(v antlr.Token) { s.mname = v }

func (s *RawDeclContext) RawOutputs() IRawOutputsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRawOutputsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRawOutputsContext)
}

func (s *RawDeclContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TSgenParserID)
}

func (s *RawDeclContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TSgenParserID, i)
}

func (s *RawDeclContext) RawInputs() IRawInputsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRawInputsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRawInputsContext)
}

func (s *RawDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RawDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitRawDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) RawDecl() (localctx IRawDeclContext) {
	localctx = NewRawDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TSgenParserRULE_rawDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(TSgenParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)

		var _m = p.Match(TSgenParserID)

		localctx.(*RawDeclContext).from = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)

		var _m = p.Match(TSgenParserID)

		localctx.(*RawDeclContext).label = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(TSgenParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)

		var _m = p.Match(TSgenParserID)

		localctx.(*RawDeclContext).mname = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(TSgenParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137625600) != 0 {
		{
			p.SetState(112)
			p.RawInputs()
		}

	}
	{
		p.SetState(115)
		p.Match(TSgenParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(TSgenParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.RawOutputs()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRawInputsContext is an interface to support dynamic dispatch.
type IRawInputsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRawParam() []IRawParamContext
	RawParam(i int) IRawParamContext

	// IsRawInputsContext differentiates from other interfaces.
	IsRawInputsContext()
}

type RawInputsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawInputsContext() *RawInputsContext {
	var p = new(RawInputsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_rawInputs
	return p
}

func InitEmptyRawInputsContext(p *RawInputsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_rawInputs
}

func (*RawInputsContext) IsRawInputsContext() {}

func NewRawInputsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawInputsContext {
	var p = new(RawInputsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_rawInputs

	return p
}

func (s *RawInputsContext) GetParser() antlr.Parser { return s.parser }

func (s *RawInputsContext) AllRawParam() []IRawParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRawParamContext); ok {
			len++
		}
	}

	tst := make([]IRawParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRawParamContext); ok {
			tst[i] = t.(IRawParamContext)
			i++
		}
	}

	return tst
}

func (s *RawInputsContext) RawParam(i int) IRawParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRawParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRawParamContext)
}

func (s *RawInputsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawInputsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RawInputsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitRawInputs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) RawInputs() (localctx IRawInputsContext) {
	localctx = NewRawInputsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TSgenParserRULE_rawInputs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.RawParam()
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSgenParserT__6 {
		{
			p.SetState(120)
			p.Match(TSgenParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.RawParam()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRawParamContext is an interface to support dynamic dispatch.
type IRawParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TypeExpr() ITypeExprContext

	// IsRawParamContext differentiates from other interfaces.
	IsRawParamContext()
}

type RawParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawParamContext() *RawParamContext {
	var p = new(RawParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_rawParam
	return p
}

func InitEmptyRawParamContext(p *RawParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_rawParam
}

func (*RawParamContext) IsRawParamContext() {}

func NewRawParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawParamContext {
	var p = new(RawParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_rawParam

	return p
}

func (s *RawParamContext) GetParser() antlr.Parser { return s.parser }

func (s *RawParamContext) ID() antlr.TerminalNode {
	return s.GetToken(TSgenParserID, 0)
}

func (s *RawParamContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *RawParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RawParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitRawParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) RawParam() (localctx IRawParamContext) {
	localctx = NewRawParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TSgenParserRULE_rawParam)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Match(TSgenParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.TypeExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.TypeExpr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRawOutputsContext is an interface to support dynamic dispatch.
type IRawOutputsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypeExpr() []ITypeExprContext
	TypeExpr(i int) ITypeExprContext

	// IsRawOutputsContext differentiates from other interfaces.
	IsRawOutputsContext()
}

type RawOutputsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawOutputsContext() *RawOutputsContext {
	var p = new(RawOutputsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_rawOutputs
	return p
}

func InitEmptyRawOutputsContext(p *RawOutputsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_rawOutputs
}

func (*RawOutputsContext) IsRawOutputsContext() {}

func NewRawOutputsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawOutputsContext {
	var p = new(RawOutputsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_rawOutputs

	return p
}

func (s *RawOutputsContext) GetParser() antlr.Parser { return s.parser }

func (s *RawOutputsContext) AllTypeExpr() []ITypeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeExprContext); ok {
			len++
		}
	}

	tst := make([]ITypeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeExprContext); ok {
			tst[i] = t.(ITypeExprContext)
			i++
		}
	}

	return tst
}

func (s *RawOutputsContext) TypeExpr(i int) ITypeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *RawOutputsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawOutputsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RawOutputsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitRawOutputs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) RawOutputs() (localctx IRawOutputsContext) {
	localctx = NewRawOutputsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TSgenParserRULE_rawOutputs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.TypeExpr()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSgenParserT__6 {
		{
			p.SetState(133)
			p.Match(TSgenParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.TypeExpr()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeExprContext is an interface to support dynamic dispatch.
type ITypeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QualID() IQualIDContext
	ELLIPSIS() antlr.TerminalNode
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode
	AllLBRACK() []antlr.TerminalNode
	LBRACK(i int) antlr.TerminalNode
	AllRBRACK() []antlr.TerminalNode
	RBRACK(i int) antlr.TerminalNode

	// IsTypeExprContext differentiates from other interfaces.
	IsTypeExprContext()
}

type TypeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprContext() *TypeExprContext {
	var p = new(TypeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_typeExpr
	return p
}

func InitEmptyTypeExprContext(p *TypeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_typeExpr
}

func (*TypeExprContext) IsTypeExprContext() {}

func NewTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprContext {
	var p = new(TypeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_typeExpr

	return p
}

func (s *TypeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExprContext) QualID() IQualIDContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualIDContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualIDContext)
}

func (s *TypeExprContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(TSgenParserELLIPSIS, 0)
}

func (s *TypeExprContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(TSgenParserSTAR)
}

func (s *TypeExprContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(TSgenParserSTAR, i)
}

func (s *TypeExprContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(TSgenParserLBRACK)
}

func (s *TypeExprContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(TSgenParserLBRACK, i)
}

func (s *TypeExprContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(TSgenParserRBRACK)
}

func (s *TypeExprContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(TSgenParserRBRACK, i)
}

func (s *TypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitTypeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) TypeExpr() (localctx ITypeExprContext) {
	localctx = NewTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TSgenParserRULE_typeExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TSgenParserELLIPSIS {
		{
			p.SetState(140)
			p.Match(TSgenParserELLIPSIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSgenParserSTAR {
		{
			p.SetState(143)
			p.Match(TSgenParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSgenParserLBRACK {
		{
			p.SetState(149)
			p.Match(TSgenParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(TSgenParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.QualID()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQualIDContext is an interface to support dynamic dispatch.
type IQualIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsQualIDContext differentiates from other interfaces.
	IsQualIDContext()
}

type QualIDContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualIDContext() *QualIDContext {
	var p = new(QualIDContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_qualID
	return p
}

func InitEmptyQualIDContext(p *QualIDContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_qualID
}

func (*QualIDContext) IsQualIDContext() {}

func NewQualIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualIDContext {
	var p = new(QualIDContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_qualID

	return p
}

func (s *QualIDContext) GetParser() antlr.Parser { return s.parser }

func (s *QualIDContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TSgenParserID)
}

func (s *QualIDContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TSgenParserID, i)
}

func (s *QualIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitQualID(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) QualID() (localctx IQualIDContext) {
	localctx = NewQualIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TSgenParserRULE_qualID)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(TSgenParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TSgenParserT__14 {
		{
			p.SetState(159)
			p.Match(TSgenParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(TSgenParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IS0DeclContext is an interface to support dynamic dispatch.
type IS0DeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsS0DeclContext differentiates from other interfaces.
	IsS0DeclContext()
}

type S0DeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyS0DeclContext() *S0DeclContext {
	var p = new(S0DeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_s0Decl
	return p
}

func InitEmptyS0DeclContext(p *S0DeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_s0Decl
}

func (*S0DeclContext) IsS0DeclContext() {}

func NewS0DeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *S0DeclContext {
	var p = new(S0DeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_s0Decl

	return p
}

func (s *S0DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *S0DeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TSgenParserID, 0)
}

func (s *S0DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S0DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *S0DeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitS0Decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) S0Decl() (localctx IS0DeclContext) {
	localctx = NewS0DeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TSgenParserRULE_s0Decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(TSgenParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(TSgenParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFDeclContext is an interface to support dynamic dispatch.
type IFDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsFDeclContext differentiates from other interfaces.
	IsFDeclContext()
}

type FDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFDeclContext() *FDeclContext {
	var p = new(FDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_fDecl
	return p
}

func InitEmptyFDeclContext(p *FDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TSgenParserRULE_fDecl
}

func (*FDeclContext) IsFDeclContext() {}

func NewFDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FDeclContext {
	var p = new(FDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSgenParserRULE_fDecl

	return p
}

func (s *FDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(TSgenParserID, 0)
}

func (s *FDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TSgenVisitor:
		return t.VisitFDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TSgenParser) FDecl() (localctx IFDeclContext) {
	localctx = NewFDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TSgenParserRULE_fDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(TSgenParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(TSgenParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
