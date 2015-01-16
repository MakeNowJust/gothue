package program

import "fmt"

// implementation of Stringer for Program
func (p *Program) String() string {
	str := ""
	for _, rule := range p.Rules {
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
