package program

// this type is the Thue program.
type Program struct {
	Rules []*Rule
	Data  []byte
}

// this type is the rule of the Thue.
type Rule struct {
	Name   []byte
	Action Rewriter
}

// this type is an rewrite action of the Thue.
type Rewriter interface {
	Rewrite(left, right []byte) ([]byte, error)
}
