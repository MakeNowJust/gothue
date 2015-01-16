package interpreter

import (
	"github.com/MakeNowJust/gothue/program"
)

import (
	"github.com/cznic/mathutil"
)

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
)

type matchData struct {
	index int
	rule  *program.Rule
}

func debugPrint(debug bool, format string, args ...interface{}) {
	if debug {
		fmt.Fprintf(os.Stderr, "\u001b[32m"+format+"\u001b[0m\n", args...)
	}
}

func Interpret(pgrm *program.Program, choice ChoiceMode, debug bool) (data []byte, err error) {
	data = pgrm.Data
	for {
		debugPrint(debug, ">>> data: %s", data)
		dataLen := len(data)
		leftIndex := mathutil.MaxInt
		rightIndex := mathutil.MinInt
		left := []*program.Rule{}
		right := []*program.Rule{}
		matches := []*matchData{}
		for _, rule := range pgrm.Rules {
			nameLen := len(rule.Name)
			for i := range data {
				if i+nameLen <= dataLen {
					if bytes.Equal(data[i:i+nameLen], rule.Name) {
						matches = append(matches, &matchData{
							index: i,
							rule:  rule,
						})
						if i == leftIndex {
							left = append(left, rule)
						}
						if i == rightIndex {
							right = append(right, rule)
						}
						if i < leftIndex {
							leftIndex = i
							left = []*program.Rule{rule}
						}
						if i > rightIndex {
							rightIndex = i
							right = []*program.Rule{rule}
						}
					}
				}
			}
		}

		// no match...
		if len(matches) == 0 {
			return
		}

		// match!
		var match *matchData
		switch choice {
		case ChoiceLeft:
			match = &matchData{
				index: leftIndex,
				rule:  left[rand.Intn(len(left))],
			}
		case ChoiceRight:
			match = &matchData{
				index: rightIndex,
				rule:  right[rand.Intn(len(right))],
			}
		case ChoiceRandom:
			match = matches[rand.Intn(len(matches))]
		}

		debugPrint(debug, ">>> match (%d..%d): %s -> %s", match.index, match.index+len(match.rule.Name), match.rule.Name, match.rule.Action)

		if data, err = match.rule.Action.Rewrite(data[:match.index], data[match.index+len(match.rule.Name):]); err != nil {
			return
		}
	}
}
