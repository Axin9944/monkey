package token

// TokenType 定义了一个字符串类型，用于表示各种标记的类型。
type TokenType string

// Token 表示一个词法单元，包含类型和具体的字面值。
type Token struct {
	// 标记的类型，例如关键字、标识符、运算符等。
	Type TokenType
	// 标记的具体字面值，例如 "let"、"123"、"+" 等
	Literal string
}

// keywords 是一个映射表，将关键字字符串映射到对应的 TokenType 常量。
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// 检查给定的表示符是否是关键字，是返回关键字常量，
// LookupIdent 检查给定的标识符是否是关键字。
// 如果是关键字，则返回对应的关键字 TokenType；
// 如果不是关键字，则返回 IDENT（表示普通标识符）
func LookupIdent(ident string) TokenType {
	// 检查标识符是否在关键字映射表中
	// 如果是关键字，返回对应的关键字 TokenType
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// 如果不是关键字，返回 IDENT
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 标识符 + 字面量
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// 运算符
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	REMAINDER = "%"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	STRING   = "STRING"
)
