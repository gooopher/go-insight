package view

import (
	"fmt"
	"testing"
)

func TestDeferCall(t *testing.T) {
	defer_call()
}

func TestDeferCall2(t *testing.T) {
	panic_recover()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

/*
打印后
打印中
打印前
panic: 触发异常 [recovered]
	panic: 触发异常


defer是先进后出,类似栈，逆序执行。
panic触发宕机前面的defer代码都会被执行。
panic内置函数停止当前goroutine的正常执行，当函数F调用panic时，函数F的正常执行被立即停止，然后运行所有在F函数中的defer函数，然后F返回到调用他的函数对于调用者G，F函数的行为就像panic一样，终止G的执行并运行G
中所defer函数，此过程会一直继续执行到goroutine所有的函数。

参考：https://www.jianshu.com/p/63e3d57f285f
*/

func panic_recover() {
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
	}()
	panic(1)
}

/**
recover没有被defer方法直接调用，会失效
*/
