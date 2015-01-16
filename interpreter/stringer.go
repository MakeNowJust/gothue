package interpreter

func (cm ChoiceMode) String() (str string) {
	switch cm {
	case ChoiceRandom:
		str = "random"
	case ChoiceLeft:
		str = "left"
	case ChoiceRight:
		str = "right"
	}
	return
}
