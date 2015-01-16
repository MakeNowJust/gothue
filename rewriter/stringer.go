package rewriter

// implementation of Stringer for TextRewriter
func (tr *TextRewriter) String() string {
	return string(tr.Text)
}

// implementation of Stringer for WriteRewriter
func (wr *WriteRewriter) String() string {
	return "~" + wr.Text
}

// implementation of Stringer for NewlineRewriter
func (_ *NewlineRewriter) String() string {
	return "~"
}

// implementation of Stringer for ReadRewriter
func (_ *ReadRewriter) String() string {
	return ":::"
}
