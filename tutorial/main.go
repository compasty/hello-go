package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/compasty/hello-go/tempconv"
)

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func incr(p *int) {
	*p++
}

func gcd(x, y int) int {
	for y != 0 {
		// 元祖复制允许同时更新多个变量的值，在赋值之前会先把右边的所有表达式先进行求值，再同一更新左边对应变量的值
		x, y = y, x%y
	}
	return x
}

func main() {
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	//     s += sep + arg
	//     sep = " "
	// }
	// fmt.Println(s)
	// 改进版本写法
	fmt.Println(strings.Join(os.Args[:], " "))
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
	i, j := 3, 5
	// 交换值
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)
	p := &i
	*p = 10
	fmt.Printf("i = %d\n", i)
	incr(p)
	fmt.Printf("i = %d\n", i)
	fmt.Printf("gcd(42, 24) = %d\n", gcd(42, 24))
	t := 12.34
	f := tempconv.Fahrenheit(t)
	fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
}
