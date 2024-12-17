package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

var symbol = [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

func calc_sum(p *[3]int) int {
	sum := 0
	for _, v := range p {
		sum += v
	}
	return sum
}

func array_work() {
	// 默认情况下数组的元素初始化为元素类型对应的零值
	var a [3]int
	a[1] = 2
	a[2] = 4
	// 数组初始化
	var b [3]float32 = [3]float32{1.1, 2.4, 5.6}
	// 也可以指定索引进行对应初始化
	var c [4]int = [4]int{0: 1, 2: 3}
	// 遍历数组
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	for _, v2 := range b {
		fmt.Printf("%.2f\n", v2)
	}
	fmt.Printf("c[2] = %d\n", c[2])
	fmt.Printf("sum(a[*]) = %d\n", calc_sum(&a))
}

func main() {
	array_work()
	fmt.Println(RMB, symbol[RMB])
}
