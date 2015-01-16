package interpreter

// this type is a type of choice mode
type ChoiceMode int

const (
	ChoiceRandom ChoiceMode = iota
	ChoiceLeft
	ChoiceRight
)
