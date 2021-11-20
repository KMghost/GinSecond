package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg     sync.WaitGroup
	number int
	// 定义 goroutine 数量
	goroutineCount = 4
	// 任务数
	taskCount = 10
)

func main() {

	// ========================================================================
	// 无缓冲 channel
	// count := make(chan int)
	// wg.Add(2)
	// go player("A", count)
	// go player("B", count)
	//
	// count <- 1
	// wg.Wait()

	// ========================================================================
	// 有缓冲 channel
	// 有缓冲通道 类型为 string, 缓冲数 taskCount 为 10
	tasks := make(chan string, taskCount)
	wg.Add(goroutineCount)

	for gr := 1; gr <= goroutineCount; gr++ {
		// 开启多个 goroutine 执行任务
		go worker(tasks, gr)
	}

	// 存放任务到任务通道中
	for task := 1; task <= taskCount; task++ {
		tasks <- fmt.Sprintf("Task %d\n", task)
	}
	// 执行完所有的 goroutine 后，关闭通道
	close(tasks)

	// 阻塞主线程，等到所有 goroutine 执行完毕在往下执行
	wg.Wait()
}
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 从任务通道中取任务 和 通道 是否关闭的状态
		task, ok := <-tasks
		// 通道已经关闭
		if !ok {
			fmt.Printf("Worker %d: shutdown\n", worker)
			return
		}
		// 模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker %d started %s", worker, task)
		// 完成工作
		fmt.Printf("Worker %d completed %s\n", worker, task)
	}
}

func player(name string, count chan int) {
	defer wg.Done()
	for {
		ball, ok := <-count
		// ok 为 false 则表示通道已经关闭
		if !ok {
			fmt.Printf("Player %s won\n", name)
			return
		}
		// 随机数
		n := rand.Intn(100)
		// 模拟接不到球
		if n%13 == 0 {
			fmt.Printf("Player %s missed\n", name)
			close(count)
			return
		}
		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++
		count <- ball
	}
}
