package types

// this type is the Thue program.
type Program struct {
	Rules []Rule
	Data  string
}

// this type is the rule of the Thue.
type Rule struct {
	Name   string
	Action Replacer
}

// this type is an replace action of the Thue.
type Replacer interface {
	Replace(left, right string) string
}

// this type is a basic replacer, such as `rule::=text`.
type TextReplacer struct {
	Text string
}

// this type is a writer, such as `rule::=~text`.
type WriteReplacer struct {
	Text string
}

// this type is a newline writer, such as `rule::=~`.
type NewlineReplacer struct{}

// this type is a reader, such as `rule::=:::`.
type ReadReplacer struct{}
