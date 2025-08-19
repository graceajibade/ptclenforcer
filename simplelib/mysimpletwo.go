package simplelib

// Raw library
type SimpleStructTwo struct {
	
}

func (x SimpleStructTwo) MM1(choice int) bool {
	if choice != 0 {
		return true
	} else {
		return false
	}
}

func (x SimpleStructTwo) MM2() string {
	return "MM2"
}

func (x SimpleStructTwo) MM3() string {
	return "MM3"
}
