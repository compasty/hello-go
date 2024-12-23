# 基础拾遗

## 打印

`fmt.Printf`格式化打印：

1. `%T`: 用于显示一个值对应的数据类型
2. `%t`: 是用于打印布尔型数据
3. `%x`: 以十六进制的格式打印数组或slice全部的元素

## 变量声明

1. `var`用来变量声明，变量会在声明时直接初始化，如果变量没有显式初始化，则被隐式地赋予其类型的 零值（zero value），数值类型是 0，字符串类型是空字符串 `""`。
2. 符号`:=`是 短变量声明 的一部分，是定义一个或多个变量并根据它们的初始值为这些变量赋予适当类型的语句，短变量声明智能用在函数的内部，不能用于包变量。

# 复合数据类型

常用的复合类型有数组、slice、maph和结构体。

## 数组

数组是由固定长度的特定类型元素组成的序列，默认情况下数组的每个元素都被初始化为元素类型对应的零值。

数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。

数组初始化方式:

```go
// 数组初始化
var b [3]float32 = [3]float32{1.1, 2.4, 5.6}
// 也可以指定索引进行对应初始化
var c [4]int = [4]int{0: 1, 2: 3}
r := [...]int{99: -1}
```

数组相等比较：只有当两个数组的长度相等，且所有元素都是相等的时候数组才是相等的。

数组作为参数传值时，GO语言会复制一个数组赋值给函数内部的参数变量，而不是传递指针或者引用。我们可以显式地传入一个数组指针，那样的话函数通过指针对数组的任何修改都可以直接反馈到调用者。

## slice

# GO并发基础

## goroutines && channels
GO并发推荐使用“顺序通信进程”（communicating sequential processes， CSP），CSP是一种现代的并发编程模型，在这种模型中值会在不同的运行实例（goroutine）中传递。 

GO语言中每一个并发的执行单元叫作一个goroutine。当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它**main goroutine**。新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。

```go
f()    // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

# 常用包

## os包

os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。
1. 程序的命令行参数可从 os 包的 Args 变量获取, `os.Args` 变量是字符串的切片（slice），其中 `os.Args[0]`是命令本身的名字，其他的元素是程序启动时传递的参数。

## strings包

1. `strings.Join`函数：连接数组或者切片形成新的字符串

## bufio包

这个包的核心作用是利用缓冲区减少IO操作次数，提升读写性能。

1. `bufio.NewScanner`: 


# 模块管理

## 工作区

Go在1.18+版本后开始支持工作区Workspace模式，