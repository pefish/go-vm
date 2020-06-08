package go_vm

type TokenType int

const (
	_ TokenType = iota
	TokenType_CONSTQ
	TokenType_ADDQ
	TokenType_SUBQ
	TokenType_MULQ
	TokenType_DIVQ
	TokenType_JMP
	TokenType_CALL
	TokenType_RET

	TokenType_INT64   // int64数值
	TokenType_STRING  // string
	TokenType_FLOAT64 // float64数值

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
	TokenType_CONSTQ: CONSTQ,
	TokenType_ADDQ:   ADDQ,
	TokenType_SUBQ:   SUBQ,
	TokenType_MULQ:   MULQ,
	TokenType_DIVQ:   DIVQ,
	TokenType_JMP:    JMP,
	TokenType_CALL:   CALL,
	TokenType_RET:    RET,
	TokenType_PRINT:  PRINT,
	TokenType_HALT:   HALT,
}

var StringToTokenType = map[string]TokenType{
	"CONSTQ": TokenType_CONSTQ,
	"ADDQ":   TokenType_ADDQ,
	"SUBQ":   TokenType_SUBQ,
	"MULQ":   TokenType_MULQ,
	"DIVQ":   TokenType_DIVQ,
	"JMP":    TokenType_JMP,
	"CALL":   TokenType_CALL,
	"RET":    TokenType_RET,
	"PRINT":  TokenType_PRINT,
	"HALT":   TokenType_HALT,
}

var TokenTypeToString = map[TokenType]string{
	TokenType_CONSTQ: "CONSTQ",
	TokenType_ADDQ:   "ADDQ",
	TokenType_SUBQ:   "SUBQ",
	TokenType_MULQ:   "MULQ",
	TokenType_DIVQ:   "DIVQ",
	TokenType_JMP:    "JMP",
	TokenType_CALL:   "CALL",
	TokenType_RET:    "RET",
	TokenType_PRINT:  "PRINT",
	TokenType_HALT:   "HALT",
}

var TokenTypeToDataType = map[TokenType]ValueType{
	TokenType_INT64:   ValueType_INT64,
	TokenType_STRING:  ValueType_STRING,
	TokenType_FLOAT64: ValueType_FLOAT64,
}
