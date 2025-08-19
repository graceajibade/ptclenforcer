package ptclenforcer

/* 
Q: what is the tool your making?
A: FSM -> typestate API (in Go)   // "design" 

Q: How will this tool be done"
A: func generateTS(fsm FSM, rawLogic rawTransition, pkg, protocol string) (apiStruct API, error) {}
A: func generateTSString(apiStruct API) (string, error) {}
A: func generateTSFile(apiString, fileName string) error {}

*/

// FSM represents a typestate finite state machine
type FSM struct {
	Labels Set[Label]  // Action labels
	States Set[State]  // States
	S0     State       // Initial state
	F      State       // Final states
	D      Transition  // Transition map
}

// ------------------------------
// FSM Structure/Types
// ------------------------------
type Label string

type State string

// Generic set implementation
type Set[T comparable] map[T]struct{}

// (state, label) → [](outcome conditions, next states)
type Transition map[StatePair][]Pair

// StatePair represents the current state 
// and a method that can be called on it
type StatePair struct {
	State State // current state
	Label Label // triggering method
}

// Pair represents a transition from one state 
// to either another state or an interface
type Pair struct {
	Outcome         bool   // used to encode branching
	TransitionState State // next state     
}

// ------------------------------
// Method Logic Structures
// ------------------------------
type RawTransition map[StatePair]Method

type MethodName string

type Method struct {
	Name    MethodName   // raw Go method name
	Inputs  []string     // parameter signatures
	Outputs []string     // return value types
}

// ------------------------------
// Typestate API Representations
// ------------------------------
type API struct {
	Typedefs   []TD // Generated state structs
	Methoddefs []MD // Method transitioning between those structs
	IFStates   IFD // Initial and Final States of the api
}

/*
The relationship between Typedefs and Methoddefs is as follow:
methoddef = func (x A) mehodName(inputs) (outputs) {}
A is the type of x and A is a member of Typedefs
*/
type TD interface { 
	TypeName() string
	Fields() []FD
}

type TDStruct struct { // Implements type TD 
	Name       string
	FieldsList []FD
}

type FD struct {
	fieldName string
	fieldType string
}

type IFD struct {
	IState State
	FState State
}

type MD struct {
	Receiver string   // The state (struct name) that owns this method
    Name     string   // Name of the method (the label in the FSM)
	RawName  string   // Name of the method of the raw api.
    Inputs   []string // List of input parameter types (e.g., "int", "string")
    Outputs  []string // List of output types returned by the method (e.g., "bool", "error")
    Returns  []Returns// The next state (struct name) or interface returned after this method executes
}

type Returns struct {
	ItfName string
	Outcome bool
	State string
}

func (t TDStruct) TypeName() string {
	return t.Name
}

func (t TDStruct) Fields() []FD {
	return t.FieldsList
}