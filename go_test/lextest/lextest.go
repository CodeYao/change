package main

import (
	"change/version_go/compiler/lexer/token"
	"fmt"
)

func main() {
	token_list := token.Lexer("./lextest")
	// for _, v := range token_list {
	// 	fmt.Printf("%+v\n", v)
	// }

	for _, v := range token_list {
		fmt.Printf("line:%d, column:%d,\ttoken_type:%d,\tstr:%s\n", v.Line, v.Column, v.Token_type, v.Str)
	}
}
