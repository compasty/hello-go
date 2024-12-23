package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	//     s += sep + arg
	//     sep = " "
	// }
	// fmt.Println(s)
	// 改进版本写法
	fmt.Println(strings.Join(os.Args[:], " "))
}
