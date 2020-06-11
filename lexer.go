package go_vm

import (
	"fmt"
	"strings"
)

type Lexer struct {
	currentChar      byte   // 当前读出来的字符
	input            string // 输入的源码
	nextReadPosition int    // 下一次要读的位置
	currentPosition  int    // 当前读到的位置（读过了）
	currentLine      int    // 当前行号
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		currentChar:      input[0], // 第一个字符立马读出来
		input:            input,
		nextReadPosition: 1,
		currentPosition:  0,
		currentLine:      0,
	}
}

// 跳过所有空格或者tab
func (lexer *Lexer) skipSpace() {
	for lexer.currentChar == ' ' || lexer.currentChar == '\t' {
		lexer.readChar()
	}
}

// 读出下一个字符
func (lexer *Lexer) readChar() {
	if lexer.nextReadPosition >= len(lexer.input) {
		lexer.currentChar = 0
	} else {
		lexer.currentChar = lexer.input[lexer.nextReadPosition]
	}
	lexer.currentPosition++
	lexer.nextReadPosition++
}

// 识别出下一个token
func (lexer *Lexer) NextToken() Token {
	//fmt.Println(lexer.currentPosition, fmt.Sprintf("--%c--", rune(lexer.currentChar)))
	var token Token
	lexer.skipSpace()

	if lexer.currentChar == '"' || lexer.currentChar == '\'' {
		lexer.readChar()
		token.StartPosition = lexer.currentPosition
		token.Literal = lexer.readString()
		token.Type = TokenType_STRING
		token.EndPosition = lexer.currentPosition
		token.LineNumber = lexer.currentLine
	} else if lexer.currentChar == '/' && lexer.nextChar() == '/' {
		lexer.readChar() // 读之前要读两次，把//读掉
		lexer.readChar()
		lexer.skipSpace()
		token.StartPosition = lexer.currentPosition
		token.Literal = lexer.readComment()
		token.Type = TokenType_COMMENT
		token.EndPosition = lexer.currentPosition
		token.LineNumber = lexer.currentLine
	} else if lexer.currentChar == '\r' || lexer.currentChar == '\n' {
		token.StartPosition = lexer.currentPosition
		token.Literal = "EOL"
		token.Type = TokenType_EOL
		token.EndPosition = lexer.currentPosition
		token.LineNumber = lexer.currentLine
		lexer.readChar()
		lexer.currentLine++
	} else if isLetter(lexer.currentChar) {
		token.StartPosition = lexer.currentPosition
		token.Literal = lexer.readInstruction()
		tokenType, ok := StringToTokenType[strings.ToUpper(token.Literal)]
		if !ok {
			panic(fmt.Errorf("illegal keyword - line: %d, position: %d, keyword: %s", lexer.currentLine, lexer.currentPosition, token.Literal))
		}
		token.Type = tokenType
		token.EndPosition = lexer.currentPosition
		token.LineNumber = lexer.currentLine
	} else if isDigit(lexer.currentChar) {
		token.StartPosition = lexer.currentPosition
		token.Literal = lexer.readDigit()
		token.Type = TokenType_NUMBER
		token.EndPosition = lexer.currentPosition
		token.LineNumber = lexer.currentLine
	} else if lexer.currentChar == 0 {
		token.StartPosition = lexer.currentPosition
		token.Literal = "EOF"
		token.Type = TokenType_EOF
		token.EndPosition = lexer.currentPosition
		token.LineNumber = lexer.currentLine
	} else {
		panic(fmt.Errorf("illegal char - line: %d, position: %d, char: %c", lexer.currentLine, lexer.currentPosition, lexer.currentChar))
	}

	//fmt.Printf("找到token: %v\n", token)
	return token
}

func (lexer *Lexer) readComment() string {
	currentPosition := lexer.currentPosition
	for lexer.currentChar != '\n' {
		lexer.readChar()
	}
	result := lexer.input[currentPosition:lexer.currentPosition]
	//lexer.readChar() // 不能读掉\n，他是token
	return result
}

func (lexer *Lexer) readString() string {
	currentPosition := lexer.currentPosition
	for lexer.currentChar != '"' && lexer.currentChar != '\'' && lexer.currentChar != '\n' {
		lexer.readChar()
	}
	result := lexer.input[currentPosition:lexer.currentPosition]
	lexer.readChar() // 读掉"或者\n
	return result
}

func (lexer *Lexer) readDigit() string {
	currentPosition := lexer.currentPosition
	for isDigit(lexer.currentChar) {
		lexer.readChar()
	}
	return lexer.input[currentPosition:lexer.currentPosition]
}

func (lexer *Lexer) nextChar() byte {
	if lexer.nextReadPosition >= len(lexer.input) {
		return 0
	}
	return lexer.input[lexer.nextReadPosition]
}

func (lexer *Lexer) readInstruction() string {
	currentPosition := lexer.currentPosition
	for isLetter(lexer.currentChar) {
		lexer.readChar()
	}
	return lexer.input[currentPosition:lexer.currentPosition]
}

func (lexer *Lexer) ParseTokens() []Token {
	var tokens []Token
	for token := lexer.NextToken(); token.Type != TokenType_EOF; token = lexer.NextToken() {
		tokens = append(tokens, token)
	}
	return tokens
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9') || ch == '.'
}
