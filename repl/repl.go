package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/axiom-labs/monkey/lexer"
	"github.com/axiom-labs/monkey/token"
)

// PROMPT designates the REPL prompt characters to accept
// user input.
const PROMPT = ">>> "

// Start will initiate a new REPL session.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
