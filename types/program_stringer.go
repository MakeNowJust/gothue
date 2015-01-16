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

// implementation of Stringer for TextReplacer
func (tr *TextReplacer) String() string {
	return tr.Text
}

// implementation of Stringer for WriteReplacer
func (wr *WriteReplacer) String() string {
	return "~" + wr.Text
}

// implementation of Stringer for NewlineReplacer
func (_ *NewlineReplacer) String() string {
	return "~"
}

// implementation of Stringer for ReadReplacer
func (_ *ReadReplacer) String() string {
	return ":::"
}
