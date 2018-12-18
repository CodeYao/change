package token

import (
	"fmt"
)

var judge_token_type map[string]TokenType

func init() {
	//初始化关键字
	judge_token_type = map[string]TokenType{
		"int":      INT,
		"char":     CHAR,
		"void":     VOID,
		"extern":   EXTERN,
		"if":       IF,
		"else":     ELSE,
		"switch":   SWITCH,
		"case":     CASE,
		"default":  DEFAULT,
		"while":    WHILE,
		"do":       DO,
		"for":      FOR,
		"return":   RETURN,
		"break":    BREAK,
		"continue": CONTINUE,
		"null":     NULL,
		"true":     TRUE,
		"false":    FALSE,
	}
}

func generate_token_list() []Token {
	var s Scanner
	var token_list []Token
	s.line = 1
	s.nextch()
	for {
		for s.ch == ' ' || s.ch == '\n' || s.ch == '\t' || s.ch == '\r' {
			s.nextch()
		}
		if s.ch == 0 {
			break
		}
		if s.ch >= 'a' && s.ch <= 'z' || s.ch >= 'A' && s.ch <= 'Z' || s.ch == '_' { //判断是关键字还是变量
			var token Token
			token.Line = s.line
			token.Column = s.column
			token.Str += string(s.ch)
			for {
				s.nextch()
				if s.ch >= 'a' && s.ch <= 'z' || s.ch >= 'A' && s.ch <= 'Z' || s.ch == '_' || s.ch >= '0' && s.ch <= '9' {
					token.Str += string(s.ch)
				} else {
					break
				}
			}
			if token_type, ok := judge_token_type[token.Str]; ok {
				token.Token_type = token_type
			} else {
				token.Token_type = IDENTITY
			}
			token_list = append(token_list, token)
		} else if s.ch == '"' { //字符串
			var token Token
			token.Line = s.line
			token.Column = s.column
			token.Token_type = STRING
			s.nextch()
			for s.ch != '"' {
				if s.ch == '\\' { //转义
					s.nextch()
					if s.ch == 'n' {
						token.Str += "\n"
					} else if s.ch == '\\' {
						token.Str += "\\"
					} else if s.ch == 't' {
						token.Str += "\t"
					} else if s.ch == '"' {
						token.Str += "\""
					} else if s.ch == '\n' { //字符串换行

					} else if s.ch == 0 { //文件结束了
						token.Token_type = ILLEGAL
						token.Column = s.column
						token.Line = s.line
						fmt.Printf("line:%d,column:%d,value:%s,message:字符串没有右引号\n", token.Line, token.Column, token.Str)
						break
					}
				} else if s.ch == '\n' || s.ch == 0 {
					token.Token_type = ILLEGAL
					token.Column = s.column
					token.Line = s.line
					fmt.Printf("line:%d,column:%d,value:%s,message:字符串没有右引号\n", token.Line, token.Column, token.Str)
					break
				} else {
					token.Str += string(s.ch)
				}
			}
			if s.ch == '"' {
				s.nextch() //读掉引号
			}
			token_list = append(token_list, token)
		} else if s.ch == '\'' { //字符
			var token Token
			token.Line = s.line
			token.Column = s.column
			token.Token_type = CHAR
			s.nextch()
			if s.ch == '\\' { //转义
				s.nextch()
				if s.ch == 'n' {
					token.Str += "\n"
				} else if s.ch == '\\' {
					token.Str += "\\"
				} else if s.ch == 't' {
					token.Str += "\t"
				} else if s.ch == '"' {
					token.Str += "\""
				} else if s.ch == '\n' || s.ch == 0 { //文件结束了
					token.Token_type = ILLEGAL
					token.Column = s.column
					token.Line = s.line
					fmt.Printf("line:%d,column:%d,value:%s,message:字符没有右单引号\n", token.Line, token.Column, token.Str)
					break
				}
			} else if s.ch == '\n' || s.ch == 0 {
				token.Token_type = ILLEGAL
				token.Column = s.column
				token.Line = s.line
				fmt.Printf("line:%d,column:%d,value:%s,message:字符串没有右引号\n", token.Line, token.Column, token.Str)
				break
			} else if s.ch == '\'' {
				token.Token_type = ILLEGAL
				token.Column = s.column
				token.Line = s.line
				fmt.Printf("line:%d,column:%d,value:%s,message:字符不可为空\n", token.Line, token.Column, token.Str)
			} else {
				token.Str += string(s.ch)
			}
			if s.ch == '\'' {
				s.nextch() //读掉引号
			}
			token_list = append(token_list, token)
		} else if s.ch >= '0' && s.ch <= '9' { //数字
			var token Token
			token.Line = s.line
			token.Column = s.column
			if s.ch != '0' { //十进制
				token.Token_type = INT
				for s.ch >= '0' && s.ch <= '9' || s.ch == '.' {
					if s.ch == '.' {
						if token.Token_type == INT { //如果之前是整数（说明之前没有.）
							token.Token_type = FLOAT
							token.Str += string(s.ch)
							s.nextch()
							if s.ch >= '0' && s.ch <= '9' {
								continue
							} else {
								token.Token_type = ILLEGAL
								token.Column = s.column
								token.Line = s.line
								fmt.Printf("line:%d,column:%d,value:%s,message:错误的数字格式\n", token.Line, token.Column, token.Str)
							}
						} else { //如果之前不是整数，说明之前已经有过一个.所以要报错
							token.Token_type = ILLEGAL
							token.Column = s.column
							token.Line = s.line
							fmt.Printf("line:%d,column:%d,value:%s,message:错误的数字格式\n", token.Line, token.Column, token.Str)
							break
						}
					}
					token.Str += string(s.ch)
					s.nextch()
				}
			} else if s.ch == '0' {
				s.nextch()
				if s.ch == 'x' { //16进制

				} else if s.ch == 'b' { //2进制

				} else if s.ch == '.' { //浮点数
					token.Token_type = FLOAT
					token.Str += "0."
					s.nextch()
					if s.ch >= '0' && s.ch <= '9' {
						for s.ch >= '0' && s.ch <= '9' {
							token.Str += string(s.ch)
							s.nextch()
						}
					} else {
						token.Token_type = ILLEGAL
						token.Column = s.column
						token.Line = s.line
						fmt.Printf("line:%d,column:%d,value:%s,message:错误的数字格式\n", token.Line, token.Column, token.Str)
					}
				} else { //0
					token.Token_type = INT
					token.Str += "0"
				}
			}
			token_list = append(token_list, token)
		} else { //界符
			var token Token
			token.Column = s.column
			token.Line = s.line
			token.Str += string(s.ch)
			s.nextch()
			switch s.lastch {
			case '+':
				if s.ch == '+' {
					token.Token_type = INC
					token.Str += string(s.ch)
					s.nextch()
				} else {
					token.Token_type = ADD
				}
			case '-':
				if s.ch == '-' {
					token.Token_type = DEC
					token.Str += string(s.ch)
					s.nextch()
				} else {
					token.Token_type = SUB
				}
			case '*':
				token.Token_type = MUL
			case '/':
				token.Token_type = DIV
			case '=':
				if s.ch == '=' {
					token.Token_type = EQ
					token.Str += string(s.ch)
					s.nextch()
				} else {
					token.Token_type = ASSIGN
				}
			case '<':
				if s.ch == '=' {
					token.Token_type = GE
					token.Str += string(s.ch)
					s.nextch()
				} else {
					token.Token_type = GT
				}
			case '>':
				if s.ch == '=' {
					token.Token_type = LE
					token.Str += string(s.ch)
					s.nextch()
				} else {
					token.Token_type = LT
				}
			case '&':
				if s.ch == '&' {
					token.Token_type = LAND
					token.Str += string(s.ch)
					s.nextch()
				} else {
					token.Token_type = AND
				}
			case '|':
				if s.ch == '|' {
					token.Token_type = LOR
					token.Str += string(s.ch)
					s.nextch()
				} else {
					token.Token_type = OR
				}
			case '(':
				token.Token_type = LPAREN
			case ')':
				token.Token_type = RPAREN
			case '{':
				token.Token_type = LBRACE
			case '}':
				token.Token_type = RBRACE
			case '[':
				token.Token_type = LBRACK
			case ']':
				token.Token_type = RBRACK
			case '.':
				token.Token_type = PERIOD
			case ',':
				token.Token_type = COMMA
			case ';':
				token.Token_type = SEMICOLON
			case ':':
				token.Token_type = COLON
			case '!':
				token.Token_type = LNOT
			default:
				token.Token_type = ILLEGAL
				token.Column = s.column
				token.Line = s.line
				fmt.Printf("line:%d,column:%d,value:%s,message:错误的符号\n", token.Line, token.Column, token.Str)
			}
			token_list = append(token_list, token)
		}
	}
	return token_list
}

func Lexer(file_path string) []Token {
	read_file(file_path)
	return generate_token_list()
}
