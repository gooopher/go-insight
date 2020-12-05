# view

- 对已经关闭的chan进行读写，会如何？
```shell script
go test skill/view/chan_closed_test.go
```

- defer、panic、recover问题
- [defer、panic用法](https://github.com/gooopher/go-insight/blob/master/skill/view/panic_defer_test.go)
    - panic后defer
    ```go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	defer_call()
    }
    
    func defer_call() {
    	defer func() { fmt.Println("打印前") }()
    	defer func() { fmt.Println("打印中") }()
    	defer func() { fmt.Println("打印后") }()
    
    	panic("触发异常")
    }
    ```
    ```shell script
    go test skill/view/panic_defer_test.go -test.run TestDeferCall
    ```
    - panic后defer中recover
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
  
    ```shell script
    go test skill/view/panic_defer_test.go -test.run TestDeferCall2
    ```
  
- go中struct能不能做比较
struct的结构不一致就是单独的类型，不同的类型不能做比较，同一类型的实例值可以做比较，但实例不可以做比较，因为其是指针类型。

- context包的用途

- client实现长连接

- 主协程如何等其余协程完再操作

- map如何实现顺序读
使用slice指定顺序取map值

- 下面这段代码输出什么，说明原因
```go
func main() {

     slice := []int{0,1,2,3}
     m := make(map[int]*int)

     for key,val := range slice {
         m[key] = &val
     }

    for k,v := range m {
        fmt.Println(k,"->",*v)
    }
}
```

结果
```
0 -> 3
1 -> 3
2 -> 3
3 -> 3
```

引用地址非值复制