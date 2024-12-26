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
var months = [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

func calcSum(p *[3]int) int {
	sum := 0
	for _, v := range p {
		sum += v
	}
	return sum
}

func arrayWork() {
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
	fmt.Printf("sum(a[*]) = %d\n", calcSum(&a))
}

func reverseSlice[T int | string](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 向左旋转n个单位
func rotate(s []int, n int) {
	reverseSlice(s[:n])
	reverseSlice(s[n:])
	reverseSlice(s)
}

// 删除切片中指定位置的的元素
func remove(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func sliceWork() {
	summer := months[6:9]
	// len = 3, cap = 7
	fmt.Printf("len(sum) = %d, cap(summer) = %d\n", len(summer), cap(summer))
	reverseSlice(summer)
	for _, v := range summer {
		fmt.Printf("%s\n", v)
	}

	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(s) = %d, cap(s) = %d\n", len(s), cap(s))
	rotate(s, 2)
	fmt.Println(s)
	s = remove(s, 2)
	fmt.Println(s)
}

func mapEquals(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func dedupLines(lines []string) {
	seen := make(map[string]bool)
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func mapWork() {
	m1 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	// traverse
	for name, age := range m1 {
		fmt.Printf("%s\t%d\n", name, age)
	}
	m2 := make(map[string]int, 8)
	m2["alice"] = 31
	m2["charlie"] = 34
	fmt.Println("m1 == m2:", mapEquals(m1, m2))
	// delete element
	delete(m2, "alice")
	fmt.Printf("len(m2) = %d\n", len(m2))
	// create an empty map
	m3 := map[string]int{}
	fmt.Printf("len(m3) = %d\n", len(m3))
	m3["alice"] = 22
	m3["bob"] = 23
	m3["charlie"] = 24
	lines := []string{"alice", "bob", "alice", "charles", "bob", "donk"}
	dedupLines(lines)

	addEdge("d1", "d2")
	addEdge("d1", "d4")
	addEdge("d2", "d4")
	fmt.Printf("hasEdge(d1, d3) = %t, hasEdge(d1, d4) = %t.\n", hasEdge("d1", "d3"), hasEdge("d1", "d4"))
}

func main() {
	arrayWork()
	fmt.Println(RMB, symbol[RMB])
	// slice operations
	sliceWork()
	// map operations
	mapWork()
}
