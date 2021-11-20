package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// 竞争状态：同一时间访问同一个共享资源的状态
	mutex  sync.Mutex
	number int64
	wg     sync.WaitGroup
)

func main() {
	// ========================================================================
	// Go 使用并发只需要在函数前使用 go 关键字

	// fmt.Println("start")
	// go func() {
	//     fmt.Println("run goroutine") // 不会被执行，因为主线程已经提前结束
	// }()
	//
	// fmt.Println("end")

	// ========================================================================
	// 但单独使用 goroutine 会导致主线程提前终止，导致 goroutine 中的代码无法被执行到，所以需要结合 sync.WaitGroup 使用。

	// fmt.Println("start")
	// // 计数器加一
	// wg.Add(1)
	//
	// go func() {
	//     // 当前函数返回时，计数器减一
	//     defer wg.Done()
	//     fmt.Println("run goroutine")
	// }()
	//
	// fmt.Println("waiting")
	// // 阻塞主进程，直到所有 goroutine 运行完毕（计数器为0时）
	// wg.Wait()
	// fmt.Println("end")

	// ========================================================================
	// 单个逻辑处理器运行多个goroutine

	// fmt.Println("开始")
	// // 计数器加2，表示需要等待2个goroutine执行完毕
	// wg.Add(2)
	// // 设置当前使用逻辑处理器的数量
	// // runtime.GOMAXPROCS(1)
	//
	// // 查看当前使用的逻辑处理器数量
	// fmt.Println(runtime.GOMAXPROCS(0)) // 1
	//
	// go printPrime("A")
	// go printPrime("B")
	//
	// fmt.Println("等待中")
	// wg.Wait()
	// fmt.Println("结束")

	// 多次运行，打印的结果顺序可能会不一样，因为并发的时间是不确定的
	// 结果大概如下：
	// 开始
	// 等待中
	// B: 2
	// ...
	// B:3907
	// A: 2
	// ...
	// A: 4999
	// B: 3911
	// ...
	// 结束

	// ========================================================================
	// 多个逻辑处理器处理多个 goroutine

	// fmt.Println("开始")
	// // 设置逻辑处理去使用当前机器的逻辑处理器数量, 默认就是使用当前机器支持的最大逻辑处理器数量
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// // 可以尝试不同逻辑处理器比较运行结果
	// runtime.GOMAXPROCS(100)
	//
	// // 查看当前逻辑处理器数量
	// // fmt.Println(runtime.GOMAXPROCS(0))
	//
	// wg.Add(2)
	//
	// go printLetter('a')
	// go printLetter('A')
	//
	// fmt.Println("等待")
	// wg.Wait()
	//
	// fmt.Println("\n结束")

	// 开始
	// 等待
	// A B C D E F a b c d e f g h i j k l m n o p q r s t u G v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q H I J K L M N O P Q R S T U V W X Y r s t u v w x y z Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
	// 结束

	// ========================================================================
	// 竞争状态

	//     fmt.Println("开始")
	//     wg.Add(2)
	//     // 查看当前逻辑处理器数量
	//     // fmt.Println(runtime.GOMAXPROCS(0))
	//
	//     go incCounter1()
	//     go incCounter1()
	//
	//     fmt.Println("等待")
	//     wg.Wait()
	//
	//     fmt.Println("\n结束", number)

	// ========================================================================
	// 原子函数 atomic.AddInt64()

	// fmt.Println("开始")
	// wg.Add(2)
	// // 查看当前逻辑处理器数量
	// //fmt.Println(runtime.GOMAXPROCS(0))
	//
	// go incCounter2()
	// go incCounter2()
	//
	// fmt.Println("等待")
	// wg.Wait()
	//
	// fmt.Println("\n结束", number)

	// ========================================================================
	// 原子函数 atomic.StoreInt64() 和 atomic.LoadInt64()

	// wg.Add(2)
	//
	// go doWork("A")
	// go doWork("B")
	//
	// time.Sleep(1 * time.Second)
	// fmt.Println("Shutdown now!")
	// // 设置 shutdown 标志
	// atomic.StoreInt64(&shutdown, 1)
	// wg.Wait()

	// ========================================================================
	// 互斥锁

	wg.Add(2)
	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Println("\n结束", number)
}
func incCounter() {
	defer wg.Done()
	for n := 0; n < 2; n++ {
		// 加锁
		mutex.Lock()
		value := number

		// 当前 goroutine 从线程退出，并放回队列
		runtime.Gosched()

		value++

		number = value
		// 解锁
		mutex.Unlock()
	}
}

var shutdown int64

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("%s Doing\n", name)
		time.Sleep(250 * time.Millisecond)
		// 获取 shutdown 值
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s down\n", name)
			break
		}
	}
}

func incCounter2() {
	defer wg.Done()
	for n := 0; n < 2; n++ {
		atomic.AddInt64(&number, 1)
		// 当前 goroutine 从线程退出，并放回队列
		runtime.Gosched()
	}
}

func incCounter1() {
	defer wg.Done()
	for n := 0; n < 2; n++ {
		value := number

		// 当前 goroutine 从线程退出，并放回队列
		runtime.Gosched()

		value++

		number = value
	}
}

// 打印字母
func printLetter(firstLetter rune) {
	defer wg.Done()
	for count := 0; count < 3; count++ {
		for char := firstLetter; char < firstLetter+26; char++ {
			fmt.Printf("%c ", char)
		}
	}
}

// 查找素数
func printPrime(prefix string) {
	defer wg.Done()
next:
	for n := 2; n < 5000; n++ {
		for m := 2; m < n; m++ {
			if n%m == 0 {
				continue next
			}
		}
		fmt.Printf("%s: %d\n", prefix, n)
	}
}
