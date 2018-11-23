package token

var judge_token_type map[string]TokenType

func init() {
	judge_token_type = map[string]TokenType{
		"+":        ADD,
		"-":        SUB,
		"*":        MUL,
		"/":        DIV,
		"%":        MOD,
		"=":        ASSIGN,
		"==":       EQ,
		"!=":       NE,
		">":        GT,
		"<":        LT,
		"<=":       GE,
		">=":       LE,
		"(":        LP,
		")":        RP,
		"{":        LC,
		"}":        RC,
		"if":       IF,
		"else":     ELSE,
		"while":    WHILE,
		"for":      FOR,
		"return":   RETURN,
		"break":    BREAK,
		"continue": CONTINUE,
		"null":     NULL,
		"true":     TRUE,
		"false":    FALSE,
	}
}

func generate_token_list(file_path string) []Token {
	var s Scanner
	var token_list []Token
	for {

		for {
			s.nextch()
			//跳过空白符
			if s.ch != ' ' || s.ch != '\n' || s.ch != '\t' {
				break
			}
		}
		if s.ch >= 'a' && s.ch <= 'z' || s.ch >= 'A' && s.ch <= 'Z' || s.ch == '_' {
			var token Token
			token.line = s.line
			token.column = s.column
			token.str += string(s.ch)
			for {
				s.nextch()
				if s.ch >= 'a' && s.ch <= 'z' || s.ch >= 'A' && s.ch <= 'Z' || s.ch == '_' {
					token.str += string(s.ch)
				} else {
					break
				}
			}
			if token_type, ok := judge_token_type[token.str]; ok {
				token.token_type = token_type
			} else {
				token.token_type = IDENTITY
			}
			token_list = append(token_list, token)
		}
	}
}