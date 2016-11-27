package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// REPLPROMPT prompt
const REPLPROMPT = ">> "

// StartREPL start repl
func StartREPL(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(REPLPROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := NewLexer(line)
		for t := l.NextToken(); t.Type != EOF; t = l.NextToken() {
			fmt.Printf("%+v\n", t)
		}
	}
}

func main() {
	StartREPL(os.Stdin, os.Stdout)
}
