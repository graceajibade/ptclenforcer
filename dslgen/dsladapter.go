package dslgen

import (
	. "school/project/cmd"
)

// GenSpecDTO is a tiny, stable Data Transfer Object (DTO) the DSL side can build.
type GenSpecDTO struct {
	Pkg, ImportPath, Protocol, FileName string
	S0, F                               string
	States, Labels                      []string     // e.g., ["Check","StartDecode",...], ["More","Decode",...]
	Edges                               [][3]string  // triples: (from, label, to)
	EdgesEx                             []EdgeDTO             
	Raw                                 map[[2]string]RawSigDTO // key = [from,label]
}

type RawSigDTO struct {
	Name            string
	Inputs, Outputs []string
}

type EdgeDTO struct {
	From, Label, To string
	Outcome *bool
}

// GenerateFromSpec converts the DTO into FSM/RawTransition and writes the file.
// It uses the real types defined in types.go and the generator functions in ptclenforcer.go.
func GenerateFromSpec(spec GenSpecDTO) (string, error) {
	// Build FSM using generic Set[T] and Transition types.
	fsm := FSM{
		States: make(Set[State]),
		Labels: make(Set[Label]),
		D:      make(Transition),
		S0:     State(spec.S0),
		F:      State(spec.F),
	}

	for _, s := range spec.States {
		fsm.States[State(s)] = struct{}{}
	}
	for _, l := range spec.Labels {
		fsm.Labels[Label(l)] = struct{}{}
	}

	// Build transitions with outcome when available
	if len(spec.EdgesEx) > 0 {
		for _, e := range spec.EdgesEx {
			from, label, to := State(e.From), Label(e.Label), State(e.To)
			k := StatePair{State: from, Label: label}
			p := Pair{TransitionState: to}
			if e.Outcome != nil {
				p.Outcome = *e.Outcome
			}
			fsm.D[k] = append(fsm.D[k], p)
		}
	} else {
		// legacy path (no explicit outcome)
		for _, e := range spec.Edges {
			from, label, to := State(e[0]), Label(e[1]), State(e[2])
			k := StatePair{State: from, Label: label}
			fsm.D[k] = append(fsm.D[k], Pair{TransitionState: to})
		}
	}


	// Raw transition wiring (Method uses MethodName for Name)
	raw := make(RawTransition)
	for k, v := range spec.Raw {
		key := StatePair{State: State(k[0]), Label: Label(k[1])}
		raw[key] = Method{
			Name:    MethodName(v.Name),
			Inputs:  v.Inputs,
			Outputs: v.Outputs,
		}
	}

	// Generate API + code + file using your existing functions.
	api, err := GenerateTSAPI(fsm, raw, spec.Pkg, spec.Protocol)
	if err != nil {
		return "", err
	}
	code, err := GenerateTSString(api, spec.ImportPath)
	if err != nil {
		return "", err
	}
	if err := GenerateTSFile(code, spec.FileName); err != nil {
		return "", err
	}
	return spec.FileName, nil
}
