package program

// this type is the Thue program.
type Program struct {
	Rules []*Rule
	Data  string
}

// this type is the rule of the Thue.
type Rule struct {
	Name   string
	Action Rewriter
}

// this type is an rewrite action of the Thue.
type Rewriter interface {
	Rewrite(left, right string) (string, error)
}
