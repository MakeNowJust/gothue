package parser

import (
	"github.com/MakeNowJust/gothue/program"
	"github.com/MakeNowJust/gothue/rewriter"
)

import (
	"fmt"
	"strings"
	"unicode"
)

type parseState int

const (
	parseRule parseState = iota
	parseData
)

// Parse and convert from source to program structure of Thue
func Parse(source string) (pgrm *program.Program, err error) {
	rules := []*program.Rule{}
	data := []string{}
	state := parseRule
	for row, line := range strings.Split(source, "\n") {
		switch state {
		case parseRule:
			rule := strings.SplitN(line, "::=", 2)
			switch len(rule) {
			case 2:
				name := strings.TrimFunc(rule[0], unicode.IsSpace)
				text := strings.TrimLeftFunc(rule[1], unicode.IsSpace)
				if name == "" {
					if text != "" {
						err = fmt.Errorf("parse error: malformed rule! (line at %d)", row)
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
						Text: text,
					}
				}
				rules = append(rules, &program.Rule{
					Name:   name,
					Action: action,
				})
			default:
				if strings.TrimFunc(line, unicode.IsSpace) != "" {
					err = fmt.Errorf("parse error: malformed rule! (line at %d)", row)
					return
				}
			}
		case parseData:
			data = append(data, line)
		}
	}
	pgrm = &program.Program{
		Rules: rules,
		Data:  strings.Join(data, "\n") + "\n",
	}
	return
}
