# JSON处理

标准库提供了 `encoding/json` 进行JSON的处理。

其中: `json.Marshal`函数返回编码后的字节，`json.MarshalIndent(v any,prefix, indent string)`可以产生整齐锁紧的输出，用来设置每一行输出的前缀和每一个层级的缩进。

在编码时，默认使用Go语言结构体的成员名字作为JSON的对象（通过reflect反射技术）。只有导出的结构体成员才会被编码（**也就是大写字母开头的成员名称**）。

可以通过结构体成员tag控制 `encoding/json` 包的编码和解码行为，成员Tag中json对应值的第一部分用于指定JSON对象的名字，比如将Go语言中的TotalCount成员对应到JSON中的total_count对象 。Color成员的Tag还带了一个额外的`omitempty`选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象（这里false为零值）。

```go
Year  int  `json:"released"`
Color bool `json:"color,omitempty"`
```