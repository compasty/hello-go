# JSON处理

标准库提供了 `encoding/json` 进行JSON的处理。

其中: `json.Marshal`函数返回编码后的字节，`json.MarshalIndent(v any,prefix, indent string)`可以产生整齐锁紧的输出，用来设置每一行输出的前缀和每一个层级的缩进。

在编码时，默认使用Go语言结构体的成员名字作为JSON的对象（通过reflect反射技术）。只有导出的结构体成员才会被编码（**也就是大写字母开头的成员名称**）。

可以通过结构体成员tag配置