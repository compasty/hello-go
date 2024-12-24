# 基础拾遗

## 类型

1. 整型: int8/int16/int32/int64, uint8/uint16/uint32/uint64; int和uint类型是对应特定CPU平台机器字大小，长度不定。`rune`是和`int32`等价的类型，通常用于表示一个Unicode码点，`byte`是`uint8`的等价类型。

> 很多时候即使数值本身不可能出现负数，也倾向于使用有符号的int类型，就像数组的长度一样。否则当我们逆序遍历时`for i := len(arr); i >= 0; i--`条件会永远为true。所以一般无符号整数一般用于位操作、哈希和加密操作等

2. 浮点型：建议使用float64, 而不是float32, 因为float32的表示范围和精度相对较差

## 打印

`fmt.Printf`格式化打印：

1. `%T`: 用于显示一个值对应的数据类型
2. `%t`: 是用于打印布尔型数据
3. `%x`: 以十六进制的格式打印数组或slice全部的元素，`%#x`会增加`0x`前缀
4. `%o`: 以八进制的格式打印, `%#o`会增加`0`前缀
5. `%g`: 以紧凑的形式打印浮点数
6. `%e`,`%f`: 以指数形式、小数形式打印浮点数，`%8.3f`表示打印宽度为8，精确为3

## 声明

1. Go语言主要有四种类型的声明语句：`var`、`const`、`type`和`func`，分别对应变量、常量、类型和函数实体对象的声明。
2. `var`用来变量声明，变量会在声明时直接初始化，如果变量没有显式初始化，则被隐式地赋予其类型的 零值（zero value），数值类型是 0，字符串类型是空字符串 `""`, 布尔类型是false，接口或者引用类型(包含slice、指针、map、chan等)对应的是`nil`。
3. 符号`:=`是 短变量声明 的一部分，是定义一个或多个变量并根据它们的初始值为这些变量赋予适当类型的语句，短变量声明只能用在函数的内部，不能用于包变量
4. 一般使用`type 类型名字 底层类型`进行类型声明，这样会创建一个新的类型名称，和现有类型具有相同的底层结构，新类型可以用来分隔不同概念的类型，这样即使底层类型相同也是不兼容的


## 可见性

1. 在函数内部定义的变量仅在函数内部有效，在函数外部定义的变量在当前包的所有文件中的都可以访问
2. 名字的开头字母大小写决定了名字在包外的可见性，大写字母开头的将是导出的，也就是可以被外部的包访问，例如`fmt`包的`Println`，包本身的名字一般总是用小写字母
3. 推荐使用 **驼峰式** 的命名方式


## new函数
表达式 `new(T)`会创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址（`*T`类型）

## 函数习惯

在GO语言中，经常会出现返回多个值的函数，通常这类函数会用额外的返回值来表达某种错误类型，一般来说返回一个error类型的错误，或者返回一个布尔值通常称为ok

```go
f, err = os.Open("foo.txt")
v, ok = m[key] // map lookup
v, ok = x.(T)  // type assertion
_, ok = m[key] 
```

## 包

一个包的源代码保存在一个或多个以.go为文件后缀名的源文件中，通常一个包所在目录路径的后缀是包的导入路径。每个包都对应一个独立的名字空间。例如，在`image`包中的`Decode`函数和在`unicode/utf16`包中的`Decode`函数是不同的。要在外部引用该函数，必须显式使用`image.Decode`或`utf16.Decode`形式访问。

在每个源文件的包声明前紧跟着的注释是包注释（§10.7.4）。通常，包注释的第一句应该先是包的功能概要说明。一个包通常只有一个源文件有包注释（译注：如果有多个包注释，目前的文档工具会根据源文件名的先后顺序将它们链接为一个包注释）。如果包注释很大，通常会放到一个独立的doc.go文件中。

包的初始化：包的初始化首先是解决包级变量的依赖顺序，然后按照包级变量声明出现的顺序依次初始化。其次如果包中含有多个.go源文件，它们将按照发给编译器的顺序进行初始化，Go语言的构建工具首先会将.go文件根据文件名排序，然后依次调用编译器编译。

特殊的init函数：有些包级别的变量需要比较复杂的初始化，此时可以使用特殊的`init`函数进行初始化，

## 字符串

一个字符串是一个不可改变的字节序列。字符串可以包含任意的数据，包括byte值0，但是通常是用来包含人类可读的文本。文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。

字符串是不可变的，所以尝试修改字符串内部结构的操作是被禁止的，执行`s[0]=L`会提示编译报错。不变性意味着如果两个字符串共享相同的底层数据的话也是安全的，这使得复制任何长度的字符串代价是低廉的。

一个原生的字符串面值形式是`` `...` ``，使用反引号代替双引号。在原生的字符串面值中，没有转义操作；全部的内容都是字面的意思，包含退格和换行，因此一个程序中的原生字符串面值可能跨越多行（译注：在原生字符串面值内部是无法直接写`` ` ``字符的，可以用八进制或十六进制转义或`` +"`" ``连接字符串常量完成）。唯一的特殊处理是会删除回车以保证在所有平台上的值都是一样的，包括那些把回车也放入文本文件的系统（译注：Windows系统会把回车和换行一起放入文本文件中）。

```go
const GoUsage = `Go is a tool for managing Go source code.

Usage:
    go command [arguments]
...`
```

### Unicode编码

目前Unicode标准里收集了超过120,000个字符，涵盖超过100多种语言。这些在计算机程序和数据中是如何体现的呢？通用的表示一个Unicode码点的数据类型是int32，也就是Go语言中rune对应的类型；它的同义词`rune`符文正是这个意思。我们可以将一个符文序列表示为一个int32序列。这种编码方式叫UTF-32或UCS-4，每个Unicode码点都使用同样大小的32bit来表示。这种方式比较简单统一，但是它会浪费很多存储空间。UTF8是一个将Unicode码点编码为字节序列的变长编码, 可以很好的节省空间，兼容ASCII编码，但是无法直接通过索引访问第n个字符。

Go语言的源文件采用UTF8编码，并且Go语言处理UTF8编码的文本也很出色。unicode包提供了诸多处理rune字符相关功能的函数（比如区分字母和数字，或者是字母的大写和小写转换等），unicode/utf8包则提供了用于rune字符序列的UTF8编码和解码的功能。例如，下面的字符串面值都表示相同的值

```
"世界"
"\xe4\xb8\x96\xe7\x95\x8c"
"\u4e16\u754c"
"\U00004e16\U0000754c"
```

`len`函数得到的是字符串占用的字节长度，如果需要关注实际的字符数量可以使用`utf8.RuneCountInString`。我们也可以借助 `unicode/utf8` 进行对应字符的遍历，当然最简单的方式是直接借助`range`循环，底层会自动隐式解码UTF8字符串。

```go
import "unicode/utf8"

s := "Hello, 世界"
fmt.Println(len(s))                    // "13"
fmt.Println(utf8.RuneCountInString(s)) // "9"
// 借助unicode/utf8包进行遍历
for i := 0; i < len(s); {
    // DecodeRuneInString函数都返回一个r和长度，r对应字符本身，长度对应r采用UTF8编码后的编码字节数目
    r, size := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%d\t%c\n", i, r)
    i += size
}
for i, r := range s {
    fmt.Printf("%d\t%q\t%d\n", i, r, r)
}
```

UTF8字符串作为交换格式的时候十分方便，但是在程序中很多时候转换为rune序列更加方便，因为rune大小一致，支持数组索引和方便切割。

```go
s := "プログラム"
//% x中用于在每个十六进制数字前插入空格
fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
// []rune和string直接相互转换
r := []rune(s)
fmt.Printf("%x\n", r)  // "[30d7 30ed 30b0 30e9 30e0]"
s2 = string(r)
```

### 字符串处理

标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。

strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。

bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。因为字符串是只读的，因此逐步构建字符串会导致很多分配和复制。在这种情况下，使用bytes.Buffer类型将会更有效。

strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。

unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。每个函数有一个单一的rune类型的参数，然后返回一个布尔值。而像ToUpper和ToLower之类的转换函数将用于rune字符的大小写转换。所有的这些函数都是遵循Unicode标准定义的字母、数字等分类规范。strings包也有类似的函数，它们是ToUpper和ToLower，将原始字符串的每个字符都做相应的转换，然后返回新的字符串。

字符串是只读的，一旦创建不可改变，但是字节slice或者rune slice则可以自由的修改。字符串和字节slice之间的转换：

```go
s := "abc"
b := []byte(s)
s2 := string(b)
```

从概念上讲，一个[]byte(s)转换是 **分配了一个新的字节数组用于保存字符串数据的拷贝**，然后引用这个底层的字节数组。编译器的优化可以避免在一些场景下分配和复制字符串数据，但总的来说需要确保在变量b被修改的情况下，原始的s字符串也不会改变。将一个字节slice转换到字符串的string(b)操作则是构造一个字符串拷贝，以确保s2字符串是只读的。

为了避免转换中不必要的内存分配，`bytes`包和`strings`包提供了很多使用的函数。

```go
// strings包中
func Contains(s, substr string) bool
func Count(s, sep string) int
func Fields(s string) []string
func HasPrefix(s, prefix string) bool
func Index(s, sep string) int
func Join(a []string, sep string) string
// bytes包中的
func Contains(b, subslice []byte) bool
func Count(s, sep []byte) int
func Fields(s []byte) [][]byte
func HasPrefix(s, prefix []byte) bool
func Index(s, sep []byte) int
func Join(s [][]byte, sep []byte) []byte
```

`bytes`包还提供了Buffer类型用于字节slice的缓存。一个Buffer开始是空的，但是随着string、byte或[]byte等类型数据的写入可以动态增长，一个bytes.Buffer变量并不需要初始化，因为零值也是有效的。

### 字符串和数字的转换

字符串和数字之间的转换一般由 `strconv`包提供。

将一个整数转为字符串，一种方法是用`fmt.Sprintf`返回一个格式化的字符串；另一个方法是用 `strconv.Itoa()`。`strconv`包的`FormatInt`和`FormatUint`函数可以用不同的进制来格式化数字

```go
x := 123
fmt.Sprintf("%d", x)
strconv.Itoa(x)

fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
// 当然用%b,%o,%x更方便
s := fmt.Sprintf("x=%b", x) // "x=1111011"
```

如果要将一个字符串解析为整数，可以使用`strconv`包的`Atoi`或`ParseInt`函数，还有用于解析无符号整数的`ParseUint`函数。ParseInt函数的第三个参数是用于指定整型数的大小；例如16表示int16，0则表示int。在任何情况下，返回的结果y总是int64类型，你可以通过强制类型转换将它转为更小的整数类型。

```go
x, err := strconv.Atoi("123")             // x is an int
y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
```

## 常量

常量都是在编译期计算，而不是运行期。

常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。在一个const声明语句中，在第一个声明的常量所在的行，iota将会被置为0，然后在每一个有常量声明的行加一。

```go
type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

type Flags uint

// 也可以在常量表达式中使用iota
const (
    FlagUp Flags = 1 << iota // is up
    FlagBroadcast            // supports broadcast access capability
    FlagLoopback             // is a loopback interface
    FlagPointToPoint         // belongs to a point-to-point link
    FlagMulticast            // supports multicast access capability
)
```

# 复合数据类型

常用的复合类型有数组、slice、map和结构体。

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

# 函数

## 函数基础

函数的声明语法
```go
func name(parameter-list) (result-list) {
    body
}
```

函数的类型被称为函数的签名。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型或签名。形参和返回值的变量名不影响函数签名，也不影响它们是否可以以省略参数类型的形式表示。

在GO中，函数支持多返回值。

# 面向对象

一个对象其实也就是一个简单的值或者一个变量，在这个对象中会包含一些方法，而一个方法则是一个和特殊类型关联的函数。一个面向对象的程序会用方法来表达其属性和对应的操作，这样使用这个对象的用户就不需要直接去操作对象，而是借助方法来做这些事情。

## 方法基础

在GO语言中，函数声明时在名字前放一个变量就是一个方法，这个附加的参数会将该函数附加到这种类型上相当于为这种类型定义了一个独占的方法。

```go
type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

上面代码中附加的参数p, 叫做方法的接收器（receiver）。Go语言中不像其他语言使用`this`或者`self`作为接收器，而是可以任意选择名字，但是为了一致性和简短性，通常建议是选择类型的第一个字母。





# GO并发基础

## goroutines && channels
GO并发推荐使用“顺序通信进程”（communicating sequential processes， CSP），CSP是一种现代的并发编程模型，在这种模型中值会在不同的运行实例（goroutine）中传递。 

GO语言中每一个并发的执行单元叫作一个goroutine。当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它**main goroutine**。新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。

```go
f()    // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

# 测试

## go test命令

`go test`命令是一个按照一定的约定和组织来测试代码的程序，在包目录内所有以 `_test.go`为后缀的文件在执行`go build`时不会被构建成包的一部分，而是`go test`的一部分。

在 `*_test.go`文件中，有三种类型的函数：测试函数、基准测试（benchmark）函数、示例函数。一个测试函数是以Test为函数名前缀的函数，用于测试程序的一些逻辑行为是否正确；go test命令会调用这些测试函数并报告测试结果是PASS或FAIL。基准测试函数是以Benchmark为函数名前缀的函数，它们用于衡量一些函数的性能；go test命令会多次运行基准测试函数以计算一个平均的执行时间。示例函数是以Example为函数名前缀的函数，提供一个由编译器保证正确性的示例文档

go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，生成一个临时的main包用于调用相应的测试函数，接着构建并运行、报告测试结果，最后清理测试中生成的临时文件。

`go test -v`命令可用于打印每个函数的名字和运行时间
`go test -run=`命令可以指定一个正则，只有测试名被正确匹配的测试函数才会被执行

## 测试函数编写

所有的测试函数必须导入`testing`包。

```go
func TestName(t *testing.T) {}
```

通常我们会在测试函数中使用 `t.Error`, `t.Errorf`, `t.Fatal`, `t.Fatalf`等函数描述错误信息。两者的区别是：`t.Errorf`调用也没有引起panic异常或停止测试的执行。即使前面的数据导致了测试的失败，依然会运行后续的测试, 而使用`t.Fatal`或`t.Fatalf`会停止当前测试函数。它们必须在和测试函数同一个goroutine内调用。

测试失败的信息一般形式是“f(x) = y, want z”，其中f(x)解释了失败的操作和对应的输入，y是实际的运行结果，z是期望的正确的结果

示例：
```go
func TestPalindrome(t *testing.T) {
    if !IsPalindrome("detartrated") {
        t.Error(`IsPalindrome("detartrated") = false`)
    }
    if IsPalindrome("palindrome") {
        t.Error(`IsPalindrome("palindrome") = true`)
    }
}
```

建议使用表格驱动的测试，容器添加测试数据
```go
func TestIsPalindrome(t *testing.T) {
    var tests = []struct {
        input string
        want  bool
    }{
        {"", true},
        {"a", true},
        {"aa", true},
        {"A man, a plan, a canal: Panama", true},
        {"Evil I did dwell; lewd did I live.", true},
        {"Able was I ere I saw Elba", true},
        {"été", true},
        {"Et se resservir, ivresse reste.", true},
        {"palindrome", false}, // non-palindrome
        {"desserts", false},   // semi-palindrome
    }
    for _, test := range tests {
        if got := IsPalindrome(test.input); got != test.want {
            t.Errorf("IsPalindrome(%q) = %v", test.input, got)
        }
    }
}
```

# 反射

Go语言提供了一种机制，能够在运行时更新变量和检查它们的值、调用它们的方法和它们支持的内在操作，而不需要在编译时就知道这些变量的具体类型。这种机制被称为反射。



# 常用包

## os包

os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。
1. 程序的命令行参数可从 os 包的 Args 变量获取, `os.Args` 变量是字符串的切片（slice），其中 `os.Args[0]`是命令本身的名字，其他的元素是程序启动时传递的参数。

## strings包

1. `strings.Join`函数：连接数组或者切片形成新的字符串

## bufio包

这个包的核心作用是利用缓冲区减少IO操作次数，提升读写性能。

1. `bufio.NewScanner`: 

## flag包

## time包

## math包

### 随机数

初始化一个随机数生成器, 以时间作为seed

```go
seed := time.Now().UTC().UnixNano()
rng := rand.New(rand.NewSource(seed))
rng.Intn(100) // 生成[0,100)的随机整数，
```

## net/http包

`net/http`包提供了HTTP客户端和服务端的实现

## net/url包


# 模块管理

## 工作区

Go在1.18+版本后开始支持工作区Workspace模式，