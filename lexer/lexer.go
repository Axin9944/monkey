package lexer

import "Monkey/token"

// Lexer 是一个词法分析器，用于将输入字符串解析为一系列标记（Token）。
type Lexer struct {
	// 输入的源代码字符串
	input        string
	position     int  // 所输入字符串中的当前位置（指向当前字符）
	readPosition int  // 所输入字符串中的当前读取位置 （指向当前字符之后的一个字符串）
	ch           byte // 当前正在查看的字符
}

// New 创建并初始化一个新的 Lexer 实例。
func New(input string) *Lexer {
	// 初始化 Lexer 结构体
	l := &Lexer{input: input}
	// 读取第一个字符
	l.readChar()
	// 返回初始化后的 Lexer 实例
	return l
}

// readChar 读取输入中的下一个字符，并更新 position 和 readPosition。
func (l *Lexer) readChar() {
	// 如果读取位置超出输入长度
	if l.readPosition >= len(l.input) {
		// 将当前字符设置为 ASCII 码 0（表示 EOF）
		l.ch = 0
	} else {
		// 否则读取当前位置的字符
		l.ch = l.input[l.readPosition]
	}
	// 更新 position 为当前读取位置
	l.position = l.readPosition
	// 增加 readPosition，准备读取下一个字符
	l.readPosition += 1
}

// NextToken 返回输入中的下一个标记（Token）。
func (l *Lexer) NextToken() token.Token {
	// 定义一个 Token 变量，用于存储生成的标记
	var tok token.Token

	// 跳过空白字符
	l.skipWhitespace()

	// 根据当前字符生成不同的标记
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '%':
		tok = newToken(token.REMAINDER, l.ch)

	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// 读取下一个字符
	l.readChar()
	// 返回生成的标记
	return tok
}

// newToken 创建一个新的 Token。
func newToken(tokenType token.TokenType, ch byte) token.Token {
	// 返回一个 Token 实例
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// 读入标识符，前移词法分析器的扫描位置，直到遇见非字母字符
// readIdentifier 读取一个标识符，并返回其字符串形式。
func (l *Lexer) readIdentifier() string {
	// 记录起始位置
	position := l.position
	// 循环读取所有连续的字母字符
	for isLetter(l.ch) {
		l.readChar()
	}
	// 返回标识符的子字符串
	return l.input[position:l.position]
}

// 判断给定的参数名称是否为字母
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 跳过空白字符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 识别数字
// readNumber 读取一个数字，并返回其字符串形式。
func (l *Lexer) readNumber() string {
	// 记录起始位置
	position := l.position
	// 循环读取所有连续的数字字符
	for isDigit(l.ch) {
		l.readChar()
	}
	// 返回数字的子字符串
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// 提前查看输入中的下一个字符是什么
// peekChar 提前查看输入中的下一个字符，但不移动读取位置。
func (l *Lexer) peekChar() byte {
	// 如果读取位置超出输入长度
	if l.readPosition >= len(l.input) {
		// 返回 ASCII 码 0（表示 EOF）
		return 0
	} else {
		// 返回下一个字符
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}
