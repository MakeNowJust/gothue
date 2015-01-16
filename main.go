package main

import (
	"github.com/MakeNowJust/gothue/interpreter"
	"github.com/MakeNowJust/gothue/parser"
)

import (
	"github.com/mattn/go-colorable"
)

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// command line options
var (
	debug  = false
	choice = interpreter.ChoiceRandom
	file   string
	seed   = time.Now().UTC().UnixNano()
)

// colorable output
var (
	stderr = colorable.NewColorableStderr()
)

// parse commnd line option
func init() {
	for _, arg := range os.Args[1:] {
		if arg[0] == '-' {
			if len(arg) >= 3 && arg[:3] == "-s=" {
				if i, err := strconv.ParseInt(arg[3:], 10, 64); err == nil {
					seed = i
				} else {
					errorPrint("command line error: malform seed number")
					os.Exit(1)
				}
			} else {
				for _, f := range arg[1:] {
					switch f {
					case 'd':
						debug = true
					case 'l':
						choice = interpreter.ChoiceLeft
					case 'r':
						choice = interpreter.ChoiceRight
					case 'h':
						help()
						os.Exit(0)
					default:
						errorPrint("command line error: unknown flag '%c'", f)
						help()
						os.Exit(1)
					}
				}
			}
		} else {
			if file == "" {
				file = arg
			} else {
				errorPrint("command line error: file name is already set")
				help()
				os.Exit(1)
			}
		}
	}

	if file == "" {
		errorPrint("command line error: please take a file name")
		help()
		os.Exit(1)
	}
}

func errorPrint(format string, args ...interface{}) {
	fmt.Fprintf(stderr, "\u001b[31m"+format+"\u001b[0m\n", args...)
}

func debugPrint(format string, args ...interface{}) {
	if debug {
		fmt.Fprintf(os.Stderr, "\u001b[32m"+format+"\u001b[0m\n", args...)
	}
}

func help() {
	fmt.Println("gothue - Thue interpreter written in Go")
	fmt.Println()
	fmt.Println("usage:")
	fmt.Printf("  %s [-s=<seed>] [-lrdh] <file name>\n", os.Args[0])
	fmt.Println()
	fmt.Println("option:")
	fmt.Println("  -l         execute leftmost matches first")
	fmt.Println("  -r         execute rightmost matches first")
	fmt.Println("  -d         debug mode")
	fmt.Println("  -s=<seed>  set a seed (default is current unix time)")
	fmt.Println("  -h         show this help")
}

func main() {
	var err error
	rand.Seed(seed)
	debugPrint("seed: %d", seed)
	if source, e := ioutil.ReadFile(file); e == nil {
		if pgrm, e := parser.Parse(string(source)); e == nil {
			debugPrint("program:\n%s", pgrm)
			if _, e := interpreter.Interpret(pgrm, choice, debug); e != nil {
				err = e
				goto exit
			}
		} else {
			errorPrint(e.Error())
			os.Exit(1)
		}
	} else {
		err = e
		goto exit
	}

	return

exit:
	errorPrint("error: %s", err.Error())
	os.Exit(1)
}
