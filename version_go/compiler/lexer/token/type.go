package token

type TokenType int

const (
	ILLEGAL TokenType = iota //非法字符
	EOF                      //结束符
	COMMENT                  //注释

	literal_beg //基本字面量开始

	IDENTITY //abc
	INT      //123
	FLOAT    //123.45
	CHAR     //'a'
	STRING   //"abc"

	literal_end //基本字面量结束

	operator_beg //操作符开始

	ADD // +
	SUB // -
	MUL // *
	DIV // /
	MOD // %

	ASSIGN // =

	EQ // ==
	NE // !=
	GT // >
	LT // <
	GE // <=
	LE // >=

	LP // (
	RP // )
	LC // {
	RC // }

	operator_end //操作符结束

	keyword_beg //关键字开始

	IF       //if
	ELSE     //else
	WHILE    //while
	FOR      //for
	RETURN   //return
	BREAK    // break
	CONTINUE //continue
	NULL     // null
	TRUE     // true
	FALSE    //false

	keyword_end //关键字结束
)

type Token struct {
	token_type TokenType //类型
	str        string    //内容
	line       int       //所在行
	column     int       //所在列
}

//是否基本字面量
func (t Token) IsLiteral() bool { return literal_beg < t.token_type && t.token_type < literal_end }

//是否操作符
func (t Token) IsOperator() bool { return operator_beg < t.token_type && t.token_type < operator_end }

//是否关键字
func (t Token) IsKeyword() bool { return keyword_beg < t.token_type && t.token_type < keyword_end }
