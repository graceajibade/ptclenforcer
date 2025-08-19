package dslgen

type DSLTransition struct {
	From, Label, Method, To, Condition, Named, Doc string
	Inputs, Outputs []string
	Repeat, Optional bool
}

type FSMDef struct {
	States  map[string]bool
	Finals  map[string]bool
	Transitions []DSLTransition
}
