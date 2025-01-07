# goroutines && channels
GO并发推荐使用“顺序通信进程”（communicating sequential processes， CSP），CSP是一种现代的并发编程模型，在这种模型中值会在不同的运行实例（goroutine）中传递。

GO语言中每一个并发的执行单元叫作一个goroutine。当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它**main goroutine**。新的goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行。而go语句本身会迅速地完成。

```go
f()    // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

主函数返回时，所有的goroutine都会被直接打断，程序退出。除了 **从主函数退出或者直接终止程序**之外，没有其它的编程方法能够让一个goroutine来打断另一个的执行。