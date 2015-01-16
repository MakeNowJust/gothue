package types

import "fmt"

// implementation of Stringer for Program
func (p *Program) String() string {
	str := ""
	for rule := range p.Rules {
		str += rule.String() + "\n"
	}
	str += "::=\n"
	str += p.Data

	return str
}

// implementation of Stringer for Rule
func (r *Rule) String() string {
	return fmt.Sprintf("%s::=%v", r.Name, r.Action)
}

// implementation of Stringer for TextRewriter
func (tr *TextRewriter) String() string {
	return tr.Text
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
