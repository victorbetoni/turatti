package repl

import (
	"bufio"
	"fmt"
	"io"
	"turatti/lexer"
	"turatti/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {

		fmt.Printf(">> ")

		if !scanner.Scan() {
			return
		}

		text := scanner.Text()
		l := lexer.New(text)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
