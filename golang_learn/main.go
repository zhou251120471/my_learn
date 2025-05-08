package main

import (
	"fmt"
	"golang_learn/function/golang_base_homework_task2"
)

// golang作业
func golang_homework1() {
	//----------------136-只出现一次的数字----------------
	//nums := []int{2, 2, 1}
	//nums := []int{4, 1, 2, 1, 2}
	//nums := []int{1}
	//result1 := golang_1_jc.SingleNumber(nums)
	//fmt.Println(result1)
	/*---------------------198. 打家劫舍---------------------------*/
	//nums := []int{1, 2, 3, 1}
	//nums := []int{2, 7, 9, 3, 1}
	//result2 := golang_1_jc.Rob(nums)
	//fmt.Println("最大值是：", result2)
	/*---------------------21. 合并两个有序链表--------------------------*/
	//l1 := []int{1, 2, 4}
	//l2 := []int{1, 3, 4}
	//golang_base_homework_task1.MergeTwoLists(l1, l2)
	/*---------------------344. 反转字符串-------------------------力扣-----------*/
	//s := []byte{'h', 'e', 'l', 'l', 'o'}
	//golang_1_jc.ReverseString(s)
	//fmt.Println(string(s))
	//fmt.Println(s)
	/*---------------------69. x 的平方根 -------------------------力扣-----------*/
	//s := 4
	//result := golang_1_jc.MySqrt(s)
	//fmt.Println(result)
	/*---------------------26. 删除有序数组中的重复项 -------------------------力扣-----------*/
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//result := golang_1_jc.RemoveDuplicates(nums)
	//fmt.Println(result)
	/*---------------------430. 扁平化多级双向链表------------------------力扣-----------*/
	//head := "[1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]"
	//golang_jc.Flatten(head)
}
func golang_homework2() {
	// 1、----------------------指针--------------------------
	//param_num := 10
	//result := golang_base_homework_task2.Add10(&param_num)
	//fmt.Println(result)
	//num1 := 1
	//num2 := 2
	//num3 := 3
	//num4 := 4
	//num5 := 5
	//param_num2 := []*int{&num1, &num2, &num3, &num4, &num5}
	//result2 := golang_base_homework_task2.Multi(param_num2)
	//fmt.Println(result2)
	// 2、----------------------Goroutine--------------------------
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go golang_base_homework_task2.PrintOdd(&wg)
	//go golang_base_homework_task2.PrintEven(&wg)
	//wg.Wait()
	//golang_base_homework_task2.RunTask()
	// 3、----------------------面向对象--------------------------
	//r := golang_base_homework_task2.Rectangle{Height: 10, Width: 10}
	//c := golang_base_homework_task2.Circle{Radius: 5}
	//fmt.Println("矩形的面积：", r.Area())
	//fmt.Println("矩形的周长", r.Perimeter())
	//fmt.Println("圆形的面积：", c.Area())
	//fmt.Println("圆形的周长", c.Perimeter())

	//e := golang_base_homework_task2.Employee{
	//	Person: golang_base_homework_task2.Person{
	//		Name: "张三",
	//		Age:  18,
	//	},
	//	EmployeeID: 1,
	//}
	//e.PrintInfo()
	// 4、----------------------Channel--------------------------------------------------------------------
	//golang_base_homework_task2.Send_Chan()
	// 创建一个容量为 10 的缓冲通道
	//ch := make(chan int, 10)
	//var wg sync.WaitGroup
	//
	//wg.Add(2)
	//// 启动生产者协程
	//go golang_base_homework_task2.Producer(ch, &wg)
	//// 启动消费者协程
	//go golang_base_homework_task2.Consumer(ch, &wg)
	//
	//// 等待生产者和消费者协程完成
	//wg.Wait()
	// 5、----------------------锁机制--------------------------------------------------------------------
	golang_base_homework_task2.Compute()
}
func exercise() {
	/*1、--------------------------------------------------------------*/
	//c := golang_1_jc.Equal(1, 2)
	//fmt.Println(c)

	/**2---------------------------------------------*/
	//var data []string = []string{"hello", "-", "world"}
	//result := golang_1_jc.Join(data)
	//fmt.Println(result)

	/*3-----------------------------------------------------------------*/
	//result := golang_1_jc.Count("小明的英文名叫jack")
	//fmt.Println(result)
	/*4-----------------------------------------------------------------*/
	//result := golang_1_jc.IsPalindrome(121)
	//fmt.Println(result)
	/*5-----------------------------------------------------------------*/
	//golang_1_jc.Formatstr(121)
	/*6-----------------------------------------------------------------*/
	//function.Sum("12", "34")
	/*7-----------------------------------------------------------------*/
	//golang_1_jc.BackList()
}

func test() {
	creditCard := &golang_base_homework_task2.CreditCard{Balance: 0, Limit: 1000}
	debitCard := &golang_base_homework_task2.DebitCard{Balance: 500}

	fmt.Println("使用信用卡购买:")
	golang_base_homework_task2.PurchaseItem(creditCard, 800)

	fmt.Println("\n使用借记卡购买:")
	golang_base_homework_task2.PurchaseItem(debitCard, 300)

	fmt.Println("\n再次使用借记卡购买:")
	golang_base_homework_task2.PurchaseItem(debitCard, 300)
}

func main() {
	/*---------------作业----------------------------*/
	golang_homework1()
	golang_homework2()
	/*---------------练习----------------------------*/
	exercise()
	//test()
}
