package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/compasty/hello-go/popcount"
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
		// 元组复制允许同时更新多个变量的值，在赋值之前会先把右边的所有表达式先进行求值，再同一更新左边对应变量的值
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
	// 引入包执行操作
	t := 12.34
	f := tempconv.Fahrenheit(t)
	fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
	fmt.Printf("bits of %d = %d\n", 3, popcount.PopCount(3))
	fmt.Printf("bits of %d = %d,%d,%d\n", 300, popcount.PopCount(300), popcount.PopCount2(300), popcount.PopCount3(300))
	// 字符串操作
	s1 := "hello, world"
	fmt.Println(len(s1))      // "12"
	fmt.Println(s1[0], s1[7]) // "104 119" ('h' and 'w')
	fmt.Println(s1[:5])       // "hello"
	// 字符串unicode编码
	s2 := "hello, 世界"
	fmt.Println(len(s2))                    // "13"
	fmt.Println(utf8.RuneCountInString(s2)) // "9"
	for i := 0; i < len(s2); {
		r, size := utf8.DecodeRuneInString(s2[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
