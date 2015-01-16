package types

// this type is the Thue program.
type Program struct {
	Rules []Rule
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

// this type is a basic rewriter, such as `rule::=text`.
type TextRewriter struct {
	Text string
}

// this type is a writer, such as `rule::=~text`.
type WriteRewriter struct {
	Text string
}

// this type is a newline writer, such as `rule::=~`.
type NewlineRewriter struct{}

// this type is a reader, such as `rule::=:::`.
type ReadRewriter struct{}
