package rewriter

import (
	"bufio"
	"fmt"
	"os"
)

var (
	stdin = bufio.NewReader(os.Stdin)
)

// implementation of Rewriter for TextRewriter
func (tr *TextRewriter) Rewrite(left, right []byte) ([]byte, error) {
	return append(append(left, tr.Text...), right...), nil
}

// implementation of Rewriter for WriterRewriter
func (wr *WriteRewriter) Rewrite(left, right []byte) ([]byte, error) {
	fmt.Print(wr.Text)
	return append(left, right...), nil
}

// implementation of Rewriter for NewlineRewriter
func (_ *NewlineRewriter) Rewrite(left, right []byte) ([]byte, error) {
	fmt.Println()
	return append(left, right...), nil
}

// implementation of Rewriter for ReadRewriter
func (_ *ReadRewriter) Rewrite(left, right []byte) ([]byte, error) {
	if line, err := stdin.ReadString('\n'); err == nil {
		if len(line) > 0 && line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		return append(append(left, []byte(line)...), right...), nil
	} else {
		return []byte{}, err
	}
}
