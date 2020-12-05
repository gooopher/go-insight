# 输出以下程序的结果
```go
func main() {
    defer func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println(r)
            } 
        }()
    }()
    panic(1)
}
```


















当函数调用panic()抛出异常时，函数将停止执行后续的普通语句，但是之前注册的defer()函数调用仍然保证会被正常执行，然后再返回到调用者。对于当前函数的调用者，因为处理异常状态还没有被捕获，所以和直接调用panic()函数的行为类似。
在异常发生时，如果在defer()中执行recover()调用，它可以捕获触发panic()时的参数，并且恢复到正常的执行流程。