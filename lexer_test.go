package go_vm

import (
	"fmt"
)

func ExampleLexer_ParseTokens() {
	lexer := NewLexer(`
CONST "haha"  // "haha" to stack

PRINT  // print

halT  // exit

`)
	tokens := lexer.ParseTokens()
	for _, token := range tokens {
		fmt.Println(token.Literal, token.LineNumber)
	}
	// Output:
	// EOL 0
	//CONST 1
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

func ExampleLexer_ParseTokens1() {
	lexer := NewLexer(`
CONST 123

CONST 32
ADD

PRINT  // print

halT  // exit

`)
	tokens := lexer.ParseTokens()
	for _, token := range tokens {
		fmt.Println(token.Literal, token.LineNumber)
	}
	// Output:
	// EOL 0
	//CONST 1
	//123 1
	//EOL 1
	//EOL 2
	//CONST 3
	//32 3
	//EOL 3
	//ADD 4
	//EOL 4
	//EOL 5
	//PRINT 6
	//print 6
	//EOL 6
	//EOL 7
	//halT 8
	//exit 8
	//EOL 8
	//EOL 9
}