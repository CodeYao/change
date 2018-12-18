package main

import (
	"change/change/version_go/compiler/lexer/token"
	"fmt"
)

var judge_token_type map[token.TokenType]string

func init() {
	//初始化关键字
	judge_token_type = map[token.TokenType]string{
		token.INT_T:    "INT_T",
		token.BOOL_T:   "BOOL_T",
		token.BYTE_T:   "BYTE_T",
		token.INT8_T:   "INT8_T",
		token.INT16_T:  "INT16_T",
		token.INT32_T:  "INT32_T",
		token.INT64_T:  "INT64_T",
		token.UINT_T:   "UINT_T",
		token.UINT8_T:  "UINT8_T",
		token.UINT16_T: "UINT16_T",
		token.UINT32_T: "UINT32_T",
		token.UINT64_T: "UINT64_T",
		token.FLOAT_T:  "FLOAT_T",
		token.DOUBLE_T: "DOUBLE_T",
		token.STRING_T: "STRING_T",
		token.MAP:      "MAP",
		token.VOID:     "VOID",
		token.EXTERN:   "EXTERN",
		token.STRUCT:   "STRUCT",
		token.ENUM:     "ENUM",

		token.PUBLIC:  "PUBLIC",
		token.PRIVATE: "PRIVATE",

		token.IMPORT:   "IMPORT",
		token.CONTRACT: "CONTRACT",

		token.IF:       "IF",
		token.ELSE:     "ELSE",
		token.SWITCH:   "SWITCH",
		token.CASE:     "CASE",
		token.DEFAULT:  "DEFAULT",
		token.WHILE:    "WHILE",
		token.DO:       "DO",
		token.FOR:      "FOR",
		token.RETURN:   "RETURN",
		token.BREAK:    "BREAK",
		token.CONTINUE: "CONTINUE",
		token.NULL:     "NULL",
		token.TRUE:     "TRUE",
		token.FALSE:    "FALSE",
	}
}
func main() {
	token_list := token.Lexer("./lextest")
	// for _, v := range token_list {
	// 	fmt.Printf("%+v\n", v)
	// }

	for _, v := range token_list {
		fmt.Printf("line:%d, column:%d,\ttoken_type:%s,\tstr:%s\n", v.Line, v.Column, judge_token_type[v.Token_type], v.Str)
	}
}
