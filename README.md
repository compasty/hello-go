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
7. `%v`: 以默认格式输出变量，`%+v`: 对结构体加字段名的方式输出，`%#v`: 以 Go 语法格式化输出

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

### 字符串拼接

在go中有很多种方式进行字符串拼接：
1. 使用连接运算符：s = "Hello," + s
2. 使用`fmt.Sprintf`进行格式化拼接：`fmt.Sprintf("[name]: %s; [age]: %d", s1, s2)`
3. 使用`bytes.Buffer`，可以使用`Grow`方法来预分配内存空间的大小，减少后期的缓冲区扩容操作，`Reset`方法可以清空缓存区
4. 使用`strings.Builder`，同样有`Grow`方法

从性能角度，建议选择4。

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

slice表示变长的序列，序列中元素的类型相同，一般写作 `[]T`。slice是一个轻量级的数据结构，底层引用一个数组对象，一个slice有三个部分组成：一个slice由三个部分构成：指针、长度和容量。指针指向第一个slice元素对应的底层数组元素的地址，要注意的是 **slice的第一个元素并不一定就是数组的第一个元素**。长度对应slice中元素的数目；长度不能超过容量，容量一般是从slice的开始位置到底层数据的结尾位置。内置的len和cap函数分别返回slice的长度和容量。

多个slice之间可以共享底层数据，并且引用的数组部分区间可能重叠。

![slice-01](images/slice01.png)

```go
months := [...]string{1: "January", /* ... */, 12: "December"}

summer := months[6:9]
fmt.Println(summer[:20]) // panic: out of range

// 这是允许的，因为没有超出cap
endlessSummer := summer[:5] // extend a slice (within capacity)
fmt.Println(endlessSummer)  // "[June July August September October]"
```

> 注意：字符串的切片操作（注意是底层字节序列，不是unicode点序列）和`[]byte`的切片是类似的

因为slice值包含指向第一个slice元素的指针，因此向函数传递slice将允许在函数内部修改底层数组的元素。换句话说，复制一个slice只是对底层的数组创建了一个新的slice别名。例如借助这个实现数组的翻转：

```go
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

a := [...]int{0, 1, 2, 3, 4, 5}
reverse(a[:])
fmt.Println(a) // "[5 4 3 2 1 0]"
```

需要注意slice类型的变量和数组类型变量初始化的差异，**slice初始化的时候没有指明序列的长度**，底层会隐式创建一个合适大小的数组，然后slice的指针指向底层的数组。

```go
a := [...]int{0,1,2,3} // 数组初始化
s := []int{0,1,2,3} // 切片初始化
```

和数组不同的是，slice之间不能比较，不能使用`==`。对于字节型slice的判断可以使用高度优化过的 `bytes.Equal`函数，其他类型只能自己写函数逐个比较。slice唯一合法的比较操作是和 `nil`比较，例如：`if s == nil {}`。一个零值的slice等于nil。一个nil值的slice并没有底层数组。一个nil值的slice的长度和容量都是0，但是也有非nil值的slice的长度和容量也是0的，例如 `[]int{}`或`make([]int, 3)[3:]`。如果测试一个slice是否为空，应该使用`len(s)==0`,而不是`s == nil`。

### make函数

可以使用内置的`make`函数创建一个指定元素类型、长度和容量的slice。容量部分可以省略，在这种情况下，容量将等于长度。在底层，make创建了一个匿名的数组变量，然后返回一个slice；只有通过返回的slice才能引用底层匿名的数组变量。在第一种语句中，slice是整个数组的view。在第二个语句中，slice只引用了底层数组的前len个元素，但是容量将包含整个的数组。额外的元素是留给未来的增长用的。

```go
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
```

### append函数

内置的`append`函数可以用于向slice追加元素，其内存扩展策略比较复杂（但一般来说有足够容量时直接设值，否则进行扩容然后复制迁移），所以通常我们并不知道append调用是否导致了内存的重新分配，因此我们也不能确认新的slice和原始的slice是否引用的是相同的底层数组空间。同样，我们不能确认在原先的slice上的操作是否会影响到新的slice。因此，通常是将append返回的结果直接赋值给输入的slice变量，例如：`runes = append(runes, r)`

### copy函数

内置函数 `copy()`可以将一个切片复制到另外一个切片中，其函数签名为：`func copy(dst, src []Type) int`, 会将`src`切片中的元素复制到`dst`中,复制长度以`len(src)`和`len(dst)`的最小值为准。它返回复制的元素个数。

```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := []int{5, 4, 3}
copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
```

## map

map底层是哈希表结构，是**无序**的key/value对的集合，其中key必须是支持 `==` 比较运算符的数据类型以保证可以测试key是否相等（因此不建议使用浮点数作为key）

可以使用key对应的元素，使用`delete`函数删除对应的元素。所有这些操作是安全的，即使这些元素不在map中也没有关系；如果一个查找失败将返回value类型对应的零值，例如，即使map中不存在“bob”下面的代码也可以正常工作，因为ages["bob"]失败时将返回0。

```go
ages["bob"] = ages["bob"] + 1 // happy birthday!
```

> map中的元素不是一个变量，因此无法对map的元素进行取址操作，例如 `&ages['bob']`会报错

map的遍历&判断key是否存在的操作：

```go
// 遍历
for k, v := range m {}
// 获取
age, ok := ages["bob"]
if !ok { // bob not in map }
// 或者结合使用
if age, ok := ages["bob"]; !ok { /* ... */ }
```

map的迭代顺序是不确定的，不同的哈希实现可能导致不同的遍历顺序。如果需要强制按顺序遍历key/value对，必须显式的对key进行排序，例如使用`sort`包。

```go
import "sort"
// 预先设定大小
names := make([]string, 0, len(ages))
for name := range ages {
    names = append(names, name)
}
sort.Strings(names)
for _, name := range names {}
```

map可以和 `nil` 比较，但是不能相互间不能进行相等比较。

# 函数

## 函数基础

函数的声明语法
```go
func name(parameter-list) (result-list) {
    body
}
```

函数的类型被称为函数的签名。如果两个函数形式参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型或签名。形参和返回值的变量名不影响函数签名，也不影响它们是否可以以省略参数类型的形式表示。

在GO中，函数支持多返回值, 最常见的就是返回想要的结果和错误信息，例如: `func findLinks(url string) ([]string, error) {}`。在多返回值时，准确的变量名可以传达函数返回值的含义，尤其在返回值的类型都相同时。

```go
func Size(rect image.Rectangle) (width, height int)
func Split(path string) (dir, file string)
```

如果一个函数所有的返回值都有显式的变量名，那么该函数的return语句可以省略操作数。这称之为bare return。

```go
func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return
}
```

## 错误处理

在Go的错误处理中，错误是软件包API和应用程序用户界面的一个重要组成部分，程序运行失败仅被认为是几个预期的结果之一。如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，通常被命名为ok，例如缓存查找场景：`value, ok := cache.Lookup(key)`, 而如果导致失败的原因不止一种，尤其是对I/O操作而言，用户需要了解更多的错误信息。因此，额外的返回值不再是简单的布尔类型，而是error类型。

内置的error是接口类型

# 面向对象

一个对象其实也就是一个简单的值或者一个变量，在这个对象中会包含一些方法，而一个方法则是一个和特殊类型关联的函数。一个面向对象的程序会用方法来表达其属性和对应的操作，这样使用这个对象的用户就不需要直接去操作对象，而是借助方法来做这些事情

## 结构体基础

结构体的变量使用点操作符进行访问，点操作符也可以和指向结构体的指针一起工作。结构体类型的零值是每个成员都是零值。通常会将零值作为最合理的默认值。

```go
var e1 *Employee = &dilbert;
e1.Position = "Senior" // 与(*e1).Position = "Senior"等价

func EmployeeByID(id int) *Employee { /* ... */ }
EmployeeByID(id).Salary = 0 // fired for... no real reason
```

> 将EmployeeByID函数的返回值从*Employee指针类型改为Employee值类型，那么更新语句将不能编译通过，因为在赋值语句的左边并不确定是一个变量（ **调用函数返回的是值，并不是一个可取地址的变量** ）

如果结构体成员名字是以大写字母开头的，那么该成员就是导出的；这是Go语言导出规则决定的。一个结构体可能同时包含导出和未导出的成员。

一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身。（该限制同样适用于数组。）但是S类型的结构体可以包含`*S`指针类型的成员。

```go
type tree struct {
    value       int
    left, right *tree
}
```

### 结构体字面值

结构体字面值可以指定每个成员的值，一般有两种写法以成员结构体定义的顺序指定，需要记住记住结构体的每个成员的类型和顺序，所以一般只在定义结构体的包内部使用，或者是在较小的结构体中使用，这些结构体的成员排列比较规则，比如 `image.Point{x, y}` 或`color.RGBA{red, green, blue, alpha}`；另一种写法就是以成员名字和相应的值初始化, 这种形式下如果成员被忽略的话将默认用零值。

```go
p := Point{1, 2}
anim := gif.GIF{LoopCount: nframes}
```

### 结构体与函数

1. 结构体可以作为函数的参数和返回值，但是考虑效率的话较大的结构体通常会用指针的方式传入和返回。
2. 需要在函数内部修改结构体成员的话，也是需要指针传入

```go
func Scale(p Point, factor int) Point {
    return Point{p.X * factor, p.Y * factor}
}
func Bonus(e *Employee, percent int) int {
    return e.Salary * percent / 100
}
// modify member values
func AwardAnnualRaise(e *Employee) {
    e.Salary = e.Salary * 105 / 100
}
```

### 结构体比较

如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用==或!=运算符进行比较。相等比较运算符==将比较两个结构体的每个成员，因此下面两个比较的表达式是等价的。**对于可比较的结构体，可以作为map的key类型。**

```go
type address struct {
    hostname string
    port     int
}

hits := make(map[address]int)
```

### 匿名成员

Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。得益于匿名嵌入的特性，我们可以直接访问叶子属性而不需要给出完整的路径。

匿名成员并不是真的无法访问了。其中匿名成员Circle和Point都有自己的名字——就是命名的类型名字——但是这些名字在点操作符中是可选的。我们在访问子成员的时候可以忽略任何匿名成员部分。

```go
// before
type Point struct {
    X, Y int
}

type Circle struct {
    Center Point
    Radius int
}
c.Center.X = 5
// after
type Circle struct {
    Point
    Radius int
}

c.X = 8 // 等价于 c.Point.X = 8
```

结构体字面值并没有简短表示匿名成员的语法，必须遵循类型声明时候的结构。

```go
c = Circle{X: 8, Y: 8, Radius: 5} // compile error: unknown fields
c = Circle{
    Point:  Point{X: 8, Y: 8},
    Radius: 5,
}
```

> 不能同时包含两个类型相同的匿名成员，这会导致名字冲突

## 方法基础

在GO语言中，函数声明时在名字前放一个变量就是一个方法，这个附加的参数会将该函数附加到这种类型上相当于为这种类型定义了一个独占的方法。

```go
type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

上面代码中附加的参数p, 叫做方法的接收器（receiver）。Go语言中不像其他语言使用`this`或者`self`作为接收器，而是可以任意选择名字，但是为了一致性和简短性，通常建议是选择类型的第一个字母。

# 接口

GO语言提供了接口类型，这是一种抽

# GO并发基础

## goroutines && channels
GO并发推荐使用“顺序通信进程”（communicating sequential processes， CSP），CSP是一种现代的并发编程模型，在这种模型中值会在不同的运行实例（goroutine）中传递。 

GO语言中每一个并发的执行单元叫作一个goroutine。当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它**main goroutine**。新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。

```go
f()    // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

主函数返回时，所有的goroutine都会被直接打断，程序退出。除了 **从主函数退出或者直接终止程序**之外，没有其它的编程方法能够让一个goroutine来打断另一个的执行。

# 泛型

Go1.18之后开始支持泛型，通过引入类型形参和实参。

```go
func Add[T int | int32 | float64 | string] (a, b T) T {
  return a + b
}

// 使用
type MySlice[T int | int32 | float32] []T
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

Go语言提供了一种机制，能够在运行时更新变量和检查它们的值、调用它们的方法和它们支持的内在操作，而不需要在编译时就知道这些变量的具体类型。这种机制被称为反射。反射是由 `reflect`包提供的，定义了两个重要类型: `Type` 和 `Value`。一个`Type`表示一个Go类型。它是一个接口，有许多方法来区分类型以及检查它们的组成部分，例如一个结构体的成员或一个函数的参数等。唯一能反映 `reflect.Type` 实现的是接口的类型描述信息，也正是这个实体标识了接口值的动态类型。

函数 `reflect.TypeOf` 接受任意的 `interface{}` 类型，并以 `reflect.Type` 形式返回其动态类型


# 模块管理

## 工作区

Go在1.18+版本后开始支持工作区Workspace模式，