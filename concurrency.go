package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
并发 concurrency
- 很多人都是冲着GO大肆宣扬的高并发而忍不住跃跃欲试，但其实从源码的解析来看，goroutine只是官方实现的超级“线程池”而已。不过话说回来，
每个实例4-5KB的栈内存占用和由于实现机制而大幅减少的创建和销魂开销，是制造GO号称的高并发的根本原因。另外，goroutine的简单易用，
也在语言层面上给予了开发者巨大的便利。

并发不是并行：Concurrency Is Not Parallelism
- 并发主要由切换时间片来实现“同时”运行，在并行则是直接利用多核实现多线程的运行，但Go可以设置使用核数，以发挥多核计算机的能力。

Goroutine奉行通过通信来共享内存，而不是共享内存来通信

*/

func main() {
	//TestGo()
	//TestChannel()
	//TestChannel2()
	//TestChannel3()
	//TestChannel4()
	//TestChannel5()
	//TestChannel6()
	//TestChannel7()
	TestSelect()
}
func TestGo() {
	//go func() {}() //也可以使用匿名函数
	go Go()
	//测试时需要下面代码堵塞一会主线程，防止主线程退出，goroutine也会退出，从而还为执行Go函数
	time.Sleep(2 * time.Second)
}
func Go() {
	fmt.Println("Go Go GO ！")
}

/*
Channel
- Channel是goroutine沟通的桥梁，大都是堵塞同步的
- 通过make创建，close关闭
- Channel是引用类型
- 可以使用for range来迭代不断操作channel
- 可以设置单向或双向通道
- 可以设置缓存大小，在未被填满前不会发生堵塞

*/
func TestChannel() {
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!")
		//存入通道c
		c <- true
	}()
	//从通道c取
	<-c
}

//可以使用for range来迭代不断操作channel
func TestChannel2() {
	//双向通道
	c := make(chan bool)
	go func() {
		fmt.Println("GO GO GO!")
		//存入通道c
		c <- true
		c <- false
		close(c) //迭代chan时需要明确关闭它，不然报错all goroutines are asleep - deadlock!
	}()
	for v := range c {
		fmt.Println("取通道的值", v)
	}

}

//channel可用设置缓存大小，不设置时默认为零就是堵塞的，如果设置了一个大小，在未被填满前都不会发生堵塞
//没有缓存的话取出操作在放入操作之前

func TestChannel3() {
	c := make(chan bool) //无缓存
	go func() {
		fmt.Println("GO GO GO!")
		//从通道c取
		<-c
	}()
	//存入通道c
	c <- true //无缓存时，设置了值但未被读出就会堵塞,会输出GO GO GO!
}
func TestChannel4() {
	c := make(chan bool, 1) //有缓存
	go func() {
		fmt.Println("GO GO GO!")
		//从通道c取
		<-c
	}()
	//存入通道c
	c <- true //有缓存时，设置了值但未被读出不会堵塞，直接结束，不会输出GO GO GO!
}

func TestChannel5() {
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go2(c, i) //go执行顺序不定，不能通过Go2中if index==9 判断 ,应用用TestChannel6中方式
	}
	//读取时堵塞
	<-c
}

//方法一  完成10个任务后结束堵塞
func TestChannel6() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //此方法不加也可以
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go3(c, i) //go执行顺序不定，不能通过Go2中if index==9 判断
	}
	//读取时堵塞
	for i := 0; i < 10; i++ {
		<-c
	}
}

//方法2  完成10个任务后结束堵塞  使用sync
func TestChannel7() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //此方法不加也可以
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go4(&wg, i) //go执行顺序不定，不能通过Go2中if index==9 判断
	}
	wg.Wait()
}

func Go2(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	if index == 9 {
		c <- true
	}
}
func Go3(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}
func Go4(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()
}

/*
Select
- 可处理一个或多个channel的发送与接收
- 同时有多个可用的channel时按随机顺序处理
- 可用空的select来堵塞main函数
- 可设置超时
*/
//一个或多个channel的发送与接收
//类似switch，只适用与channel
func TestSelect() {
	c1, c2 := make(chan int), make(chan string)
	c
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()
	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"
	close(c1)
	close(c2)

	//需要另外加一个控制channel,不然可能只输出c1 1 /n  c2 hi/n  c1 3就退出了

}
