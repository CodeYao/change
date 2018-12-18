package token

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const BUFF_SIZE = 512 //缓冲区大小

var reader *bufio.Reader //文件内容

type Scanner struct {
	line   int  // 行号
	column int  //列号
	lastch byte //上一个字符
	ch     byte // 当前字符
}

func read_file(file_path string) {
	if file, err := os.Open(file_path); err == nil {
		// defer file.Close()
		reader = bufio.NewReaderSize(file, 4096)
	} else {
		log.Fatalf("read file err :[%v]", err)
	}

}

func (s *Scanner) nextch() {
	if reader == nil {
		log.Fatalf("no file!!!")
	}
	s.lastch = s.ch
	if s.lastch == '\n' {
		s.line += 1
		s.column = 0
	}
	byt, err := reader.ReadByte()
	if err != nil {
		fmt.Println(err)
		fmt.Println("END")
	}
	// fmt.Printf("get char : [%c]\n", byt)
	s.column += 1
	s.ch = byt
}
