package go_vm

type TokenType int

const (
	_ TokenType = iota
	TokenType_CONST
	TokenType_ADD
	TokenType_SUB
	TokenType_MUL
	TokenType_DIV
	TokenType_JMP
	TokenType_JEQ
	TokenType_JNE
	TokenType_CALL
	TokenType_RET

	TokenType_STRING // string
	TokenType_NUMBER // float64数值

	TokenType_EOL     // 换行
	TokenType_EOF     // 结束
	TokenType_COMMENT // 注释
	// 下面是预设函数
	TokenType_PRINT
	TokenType_HALT
)

type Token struct {
	Type          TokenType // token类型
	Literal       string    // token的字面
	LineNumber    int       // 行号，报错时易于指出位置
	StartPosition int       // 开始索引
	EndPosition   int       // 结束索引
}

var TokenTypeToOpCode = map[TokenType]OpCode{
	TokenType_CONST: CONST,
	TokenType_ADD:   ADD,
	TokenType_SUB:   SUB,
	TokenType_MUL:   MUL,
	TokenType_DIV:   DIV,
	TokenType_JMP:   JMP,
	TokenType_JEQ:   JEQ,
	TokenType_JNE:   JNE,
	TokenType_CALL:  CALL,
	TokenType_RET:   RET,
	TokenType_PRINT: PRINT,
	TokenType_HALT:  HALT,
}

var StringToTokenType = map[string]TokenType{
	"CONST": TokenType_CONST,
	"ADD":   TokenType_ADD,
	"SUB":   TokenType_SUB,
	"MUL":   TokenType_MUL,
	"DIV":   TokenType_DIV,
	"JMP":   TokenType_JMP,
	"JEQ":   TokenType_JEQ,
	"JNE":   TokenType_JNE,
	"CALL":  TokenType_CALL,
	"RET":   TokenType_RET,
	"PRINT": TokenType_PRINT,
	"HALT":  TokenType_HALT,
}

var TokenTypeToString = map[TokenType]string{
	TokenType_CONST: "CONST",
	TokenType_ADD:   "ADD",
	TokenType_SUB:   "SUB",
	TokenType_MUL:   "MUL",
	TokenType_DIV:   "DIV",
	TokenType_JMP:   "JMP",
	TokenType_JEQ:   "JEQ",
	TokenType_JNE:   "JNE",
	TokenType_CALL:  "CALL",
	TokenType_RET:   "RET",
	TokenType_PRINT: "PRINT",
	TokenType_HALT:  "HALT",
}

var TokenTypeToDataType = map[TokenType]ValueType{
	TokenType_STRING: ValueType_STRING,
	TokenType_NUMBER: ValueType_NUMBER,
}
