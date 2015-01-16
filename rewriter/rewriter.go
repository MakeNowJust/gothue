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
func (tr *TextRewriter) Rewrite(left, right string) (string, error) {
	return left + tr.Text + right, nil
}

// implementation of Rewriter for WriterRewriter
func (wr *WriteRewriter) Rewrite(left, right string) (string, error) {
	fmt.Print(wr.Text)
	return left + right, nil
}

// implementation of Rewriter for NewlineRewriter
func (_ *NewlineRewriter) Rewrite(left, right string) (string, error) {
	fmt.Println()
	return left + right, nil
}

// implementation of Rewriter for ReadRewriter
func (_ *ReadRewriter) Rewrite(left, right string) (string, error) {
	if line, err := stdin.ReadString('\n'); err != nil {
		return left + line + right, nil
	} else {
		return "", err
	}
}
