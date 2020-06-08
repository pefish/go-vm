package go_vm

import (
	"fmt"
)

func ExampleLexer_ParseTokens() {
	lexer := NewLexer(`
CONSTQ "haha"  // "haha" to stack

PRINT  // print

halT  // exit

`)
	tokens := lexer.ParseTokens()
	for _, token := range tokens {
		fmt.Println(token.Literal, token.LineNumber)
	}
	// Output:
	// EOL 0
	//CONSTQ 1
	//haha 1
	//"haha" to stack 1
	//EOL 1
	//EOL 2
	//PRINT 3
	//print 3
	//EOL 3
	//EOL 4
	//halT 5
	//exit 5
	//EOL 5
	//EOL 6
}