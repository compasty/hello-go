package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func checkReader() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 2048)
	// 读取最多 10 个字节数据到字节切片 b 中
	b := make([]byte, 10)
	n, err := reader.Read(b)
	if err != nil {
		panic(err)
	}
	// 输出
	fmt.Printf("Read %d bytes: %s\n", n, b)
	// 返回缓存中现有的可读取的字节数
	fmt.Printf("Unread %d bytes\n", reader.Buffered())

	// 读取一个字节, 如果读取不成功会返回 Error
	c, err := reader.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read 1 byte: %c\n", c)

	// 还原最近一次读取操作读出的最后一个字节
	err = reader.UnreadByte()
	if err != nil {
		panic(err)
	}

	// 读取一个字符, 如果读取不成功会返回 Error
	r, size, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read 1 character, %d bytes：%c\n", size, r)

	// 还原前一次 ReadRune 操作读取的字符
	err = reader.UnreadRune()
	if err != nil {
		panic(err)
	}

	// 读取到分隔符，包含分隔符，返回字节切片
	b, err = reader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read bytes: %s", b)

	// 读取到分隔符，包含分隔符，返回字符串
	str, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Printf("Read string: %s", str)
}

func checkReadAsLines() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 2048)
	var lines []string
	for {
		// 读取出来的字符串是包含换行符\n
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// err 为 io.EOF 时，str 返回可能为 ""，也可能非 ""
				if len(str) != 0 {
					lines = append(lines, str)
				}
				break
			}
			panic(err)
		}
		lines = append(lines, str)
	}
	for idx, line := range lines {
		fmt.Printf("Line %d: %s", idx, line)
	}
}

func checkPeek() {
	file, err := os.Open("test2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 512)
	// 读取下 5 个字节的数据，文件读偏移量不动，返回的字节切片是内部 buf 的引用
	b, err := reader.Peek(3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Peeked at 3 bytes: %s\n", b)

	b, err = reader.Peek(12)
	// 由于文件长度不够12个字符，所以err会返回io.EOF
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Printf("Peeked at 12 bytes: %s\n", b)

	// 每两个字符循环读取
	bytes := make([]byte, 2)
	for {
		n, err := reader.Read(bytes)
		if err != nil {
			if err == io.EOF {
				if n != 0 {
					fmt.Printf("Read %d bytes: %s\n", n, bytes[0:n])
				}
				break
			} else {
				panic(err)
			}
		} else {
			fmt.Printf("Read %d bytes: %s\n", n, bytes[0:n])
		}
	}
}

func checkWriter() {
	file, err := os.Create("test3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 默认缓存区大小为 4096 字节
	writer := bufio.NewWriter(file)

	// 写字节切片到缓存
	n, err := writer.Write([]byte{65, 66, 67})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes written: %d\n", n)

	// 检查缓存中的已使用的字节大小
	unFlushedNum := writer.Buffered()
	fmt.Printf("Bytes buffered: %d\n", unFlushedNum)

	// 检查缓存中未使用的字节大小
	availableNum := writer.Available()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Available buffer: %d\n", availableNum)

	// 写单个字节到缓存
	err = writer.WriteByte('!')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes written: 1\n")

	// 写单个字符到缓存
	n, err = writer.WriteRune('您')
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes written: %d\n", n)

	// 写字符串到缓存
	n, err = writer.WriteString("Hello World")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bytes written: %d\n", n)

	// 将缓存数据刷新到底层的 io.Writer 对象
	writer.Flush()
}

var lines = []string{
	"Hello World",
	"你好，世界！",
	"欢迎来到Go语言的世界",
}

func checkWriteLines() {
	file, err := os.Create("test4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriterSize(file, 1024)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}

func checkScanner() {

}

func main() {
	checkReader()
	fmt.Println("==========")
	checkReadAsLines()
	fmt.Println("==========")
	checkPeek()
	fmt.Println("==========")
	checkWriter()
	fmt.Println("==========")
	checkWriteLines()
	fmt.Println("==========")
	checkScanner()
}
