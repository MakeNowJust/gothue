package types

// this type is the Thue program.
type Program struct {
	Rules []Rule
	Data  string
}

// this type is the rule of the Thue.
type Rule struct {
	Name   string
	Action Rewiter
}

// this type is an rewrite action of the Thue.
type Rewiter interface {
	Replace(left, right string) string
}

// this type is a basic rewriter, such as `rule::=text`.
type TextRewiter struct {
	Text string
}

// this type is a writer, such as `rule::=~text`.
type WriteRewiter struct {
	Text string
}

// this type is a newline writer, such as `rule::=~`.
type NewlineRewiter struct{}

// this type is a reader, such as `rule::=:::`.
type ReadRewiter struct{}
