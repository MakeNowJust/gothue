package types

// implementation of Rewriter for TextRewriter
func (tr *TextRewriter) Rewrite(left, right string) (string, error) {
	return left + tr.Text + right
}
