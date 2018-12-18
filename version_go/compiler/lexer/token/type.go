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

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	ADD_ASSIGN     // +=
	SUB_ASSIGN     // -=
	MUL_ASSIGN     // *=
	DIV_ASSIGN     // /=
	MOD_ASSIGN     // %=
	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=

	INC // ++
	DEC // --

	ASSIGN // =

	LAND // &&
	LOR  // ||
	EQ   // ==
	NE   // !=
	GT   // >
	LT   // <
	GE   // >=
	LE   // <=
	LNOT // !

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :

	operator_end //操作符结束

	keyword_beg //关键字开始

	INT_T    //int
	BOOL_T   //bool
	BYTE_T   //byte
	INT8_T   //int8
	INT16_T  //int16
	INT32_T  //int32
	INT64_T  //int64
	UINT_T   //uint
	UINT8_T  //uint8
	UINT16_T //uint16
	UINT32_T //uint32
	UINT64_T //uint64
	FLOAT_T  //float
	DOUBLE_T //double
	STRING_T //string
	MAP      //map
	VOID     //void
	EXTERN   //extern
	STRUCT   //struct
	ENUM     //enum

	PUBLIC  //public
	PRIVATE //private

	IMPORT   //import
	CONTRACT //contract

	IF       //if
	ELSE     //else
	SWITCH   //switch
	CASE     //case
	DEFAULT  //default
	WHILE    //while
	DO       //do
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
	Token_type TokenType //类型
	Str        string    //内容
	Line       int       //所在行
	Column     int       //所在列
}

//是否基本字面量
func (t Token) IsLiteral() bool { return literal_beg < t.Token_type && t.Token_type < literal_end }

//是否操作符
func (t Token) IsOperator() bool { return operator_beg < t.Token_type && t.Token_type < operator_end }

//是否关键字
func (t Token) IsKeyword() bool { return keyword_beg < t.Token_type && t.Token_type < keyword_end }
