package rewriter

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
