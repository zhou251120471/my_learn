package main

import (
	"fmt"
	"golang_learn/function/golang_1_jc"
	"golang_learn/function/golang_2_jj"
)

// golang作业
func golang_homework() {
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
	l1 := []int{1, 2, 4}
	l2 := []int{1, 3, 4}
	golang_1_jc.MergeTwoLists(l1, l2)
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
	creditCard := &golang_2_jj.CreditCard{Balance: 0, Limit: 1000}
	debitCard := &golang_2_jj.DebitCard{Balance: 500}

	fmt.Println("使用信用卡购买:")
	golang_2_jj.PurchaseItem(creditCard, 800)

	fmt.Println("\n使用借记卡购买:")
	golang_2_jj.PurchaseItem(debitCard, 300)

	fmt.Println("\n再次使用借记卡购买:")
	golang_2_jj.PurchaseItem(debitCard, 300)
}

func main() {
	/*---------------作业----------------------------*/
	golang_homework()
	/*---------------练习----------------------------*/
	exercise()
	test()
}
