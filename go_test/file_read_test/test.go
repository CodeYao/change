package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	line := 1
// 	file_obj, _ := os.Open("testfile")
// 	defer file_obj.Close()
// 	for {
// 		get := make([]byte, 3)
// 		if i, _ := file_obj.Read(get); i == 0 {
// 			break
// 		}
// 		fmt.Printf("\n====get:[%s],len:[%d]====\n", get, len(get))
// 		for _, v := range get {
// 			if v == '\n' {
// 				fmt.Println("add line")
// 				line += 1
// 			}
// 		}

// 	}
// 	fmt.Println("line :", line)
// }

func main() {
	line := 1
	file_obj, _ := os.Open("testfile")
	defer file_obj.Close()

	reader := bufio.NewReaderSize(file_obj, 1024)
	for {
		byt, err := reader.ReadByte()
		if err != nil {
			fmt.Println(err)
			break
		}
		if byt == '\n' {
			line++
		}
		fmt.Printf("byt : %c\n", byt)
	}
	fmt.Println("line :", line)
}
