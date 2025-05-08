package golang_base_homework_task2

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
// 1、----------------------指针--------------------------------------------------------------------------
func Add10(num *int) int {
	*num += 10
	return *num
}
func Multi(nums []*int) []int {
	result := []int{}
	for _, v := range nums {
		result = append(result, (*v)*2)
	}
	return result
}

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
// 2、----------------------Goroutine------------------------------------------------------------------------
// printOdd 函数用于打印从 1 到 10 的奇数
func PrintOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for n := 1; n <= 10; n += 2 {
		fmt.Println("奇数：", n)
	}
}

// printEven 函数用于打印从 2 到 10 的偶数
func PrintEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for n := 2; n <= 10; n += 2 {
		fmt.Println("偶数：", n)
	}
}

// --------------------设计一个任务调度器
// 定义一个无参数无返回值的函数
type Task func()

// 计算单个任务执行的时间
func singleTaskTime(task Task, taskID int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	task()
	elapsed := time.Since(start)
	// 输出任务执行时间
	fmt.Printf("任务 %d 执行耗时: %v\n", taskID, elapsed)
}

// 调度任务
func taskScheduler(tasks []Task) {
	var wg sync.WaitGroup
	//设置计数器
	wg.Add(len(tasks))
	//遍历所有任务
	for index, task := range tasks {
		go singleTaskTime(task, index, &wg)
	}
	//等待
	wg.Wait()
}

// 执行任务
func RunTask() {
	task1 := func() {
		time.Sleep(1 * time.Second)
	}
	task2 := func() {
		time.Sleep(2 * time.Second)
	}
	task3 := func() {
		time.Sleep(3 * time.Second)
	}
	tasks := []Task{task1, task2, task3}

	taskScheduler(tasks)
}

// 1. 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// - 考察点 ：接口的定义与实现、面向对象编程风格。
// 2. 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// - 考察点 ：组合的使用、方法接收者。
// 3、----------------------面向对象--------------------------------------------------------------------
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 定义矩形结构体
type Rectangle struct {
	Height float64
	Width  float64
}

// Circle 定义圆形结构体
type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("姓名: %s, 年龄: %d, 员工编号: %d\n", e.Name, e.Age, e.EmployeeID)
}

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
// 4、----------------------Channel--------------------------------------------------------------------
func Send_Chan() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 10; i++ {
			//fmt.Printf("Sent Before: %d\n", i)
			ch <- i
			fmt.Printf("Sent After: %d\n", i)
		}

	}()
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("Received: %d\n", num)
		}
	}()
	wg.Wait()
	fmt.Println("All tasks completed")
}

// producer 函数作为生产者，向缓冲通道发送 100 个整数
func Producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		fmt.Println("发送前：", i)
		ch <- i
		fmt.Println("发送后：", i)
	}
	close(ch)
}

// consumer 函数作为消费者，从缓冲通道接收整数并打印
func Consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("接收：", num)
	}
}

// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。
// 5	、----------------------锁机制--------------------------------------------------------------------
type Counter struct {
	Value int
	Mutex sync.Mutex
}

func (c *Counter) Inc() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Value++
	fmt.Println("当前值：", c.Value)

}
func (c *Counter) GetValue() int {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.Value
}
func Compute() {
	counter := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Inc()
			}
		}()
	}
	wg.Wait()
	// 输出计数器的最终值
	fmt.Println("计数器的最终值:", counter.GetValue())
}
