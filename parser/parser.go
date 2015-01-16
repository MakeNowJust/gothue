package parser

import (
	"github.com/MakeNowJust/gothue/program"
	"github.com/MakeNowJust/gothue/rewriter"
)

import (
	"fmt"
	"strings"
)

type parseState int

const (
	parseRule parseState = iota
	parseData
)

// Parse and convert from source to program structure of Thue
func Parse(source string) (pgrm *program.Program, err error) {
	if len(source) > 0 && source[len(source)-1] == '\n' {
		source = source[:len(source)-1]
	}

	rules := []*program.Rule{}
	data := []string{}
	state := parseRule
	for row, line := range strings.Split(source, "\n") {
		switch state {
		case parseRule:
			rule := strings.SplitN(line, "::=", 2)
			switch len(rule) {
			case 2:
				name := rule[0]
				text := rule[1]
				if strings.TrimSpace(name) == "" {
					if strings.TrimSpace(text) != "" {
						err = fmt.Errorf("parse error: malformed rule (line at %d)", row+1)
						return
					}
					state = parseData
					break
				}

				var action program.Rewriter
				if text[0] == '~' { // write
					if len(text) == 1 { // newline
						action = &rewriter.NewlineRewriter{}
					} else {
						action = &rewriter.WriteRewriter{
							Text: text[1:],
						}
					}
				} else if text == ":::" { // read
					action = &rewriter.ReadRewriter{}
				} else { // replace text
					action = &rewriter.TextRewriter{
						Text: []byte(text),
					}
				}
				rules = append(rules, &program.Rule{
					Name:   []byte(name),
					Action: action,
				})
			default:
				if strings.TrimSpace(line) != "" {
					err = fmt.Errorf("parse error: malformed rule (line at %d)", row+1)
					return
				}
			}
		case parseData:
			data = append(data, line)
		}
	}

	if state != parseData {
		err = fmt.Errorf("parse error: not terminated a rulebase")
		return
	}

	pgrm = &program.Program{
		Rules: rules,
		Data:  []byte(strings.Join(data, "\n")),
	}
	return
}
