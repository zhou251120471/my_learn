package golang_1_jc

import (
	"fmt"
	"strconv"
)

/*---------------------牛客---------------------------*/
func Equal(a int, b int) []bool {
	/*1、--------------------------------------------------------------
	给定两个变量a,b，判断两个变量的地址，值（a，b的地址取得）是否相等,将结果依次存入切片，并返回。
	知识点：
	Go语言中的指针操作非常简单，只需要记住两个符号：&（取地址）和*（根据地址取值）。
	每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
	Go语言中使用&字符放在变量前面对变量进行“取地址”操作。
	Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等。*/
	// write code here
	c := make([]bool, 0)
	c = append(c, &a == &b)
	c = append(c, a == b)
	return c
}
func Join(s []string) string {
	/*
		给定一个字符串数组，将其拼接成一个字符串。
		知识点：
		1,for循环遍历切片
		2,"+"可以拼接字符串
		输入：["hello","-","world"]
		返回值：	"hello-world"
	*/
	// write code here
	var result string
	for i := range s {
		result += s[i]
	}
	return result
}
func Count(s string) int {
	// write code here
	/*
		给定一个字符串，统计其中的字符个数（一个中文算一个）。
		知识点：
		1,汉字是采用unicode编码，占三个字节
		2,字符传转化为rune数组
		2,rune是int32的别名（-231~231-1），对比byte（-128～127），可表示的字符更多
		3,len()可以求出rune数组的长度

		"小明的英文名叫jack"
		11
	*/
	length := []rune(s)
	return len(length)
}
func IsPalindrome(x int) bool {
	// write code here
	/*
		给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
		例如，121 是回文，而 123 不是。
		整型转化为字符串
		字符串的遍历

		121
		true
	*/
	input := strconv.Itoa(x)
	result := ""
	for i := 0; i < len(input); i++ {
		result += string(input[len(input)-1-i])
	}
	if result == input {
		return true
	} else {
		return false
	}
}
func Formatstr(num int) {
	// write code here
	/*
			给定一个正整数，将其转换为字符串类型。

		知识点：
		Go 使用 import 关键字来导入包
		Go 可以使用 fmt.Sprintf 来格式化字符串，fmt.Sprintf(格式化样式, 参数列表…),格式化样式如下：
		    %v   按值的本来值输出
		    %+v  在 %v 基础上，对结构体字段名和值进行展开
		    %#v  输出 Go 语言语法格式的值
		    %T   输出 Go 语言语法格式的类型和值
		    %%   输出 % 本体
		    %b   整型以二进制方式显示
		    %o   整型以八进制方式显示
		    %d   整型以十进制方式显示
		    %x   整型以十六进制方式显示
		    %X   整型以十六进制、字母大写方式显示
		    %U   Unicode 字符
		    %f   浮点数
		    %p   指针，十六进制方式显示
	*/
	type Number struct {
		Num int
	}
	number := Number{
		Num: num,
	}
	str := strconv.Itoa(num)
	// %v 按值的本来值输出
	formattedV := fmt.Sprintf("按值的本来值输出: %v", num)
	fmt.Println(formattedV)
	// %+v 在 %v 基础上，对结构体字段名和值进行展开
	formattedPlusV := fmt.Sprintf("展开结构体: %+v", number)
	fmt.Println(formattedPlusV)
	// %#v 输出 Go 语言语法格式的值
	formattedSharpV := fmt.Sprintf("Go 语言语法格式的值: %#v", number)
	fmt.Println(formattedSharpV)
	// %T 输出 Go 语言语法格式的类型和值
	formattedT := fmt.Sprintf("类型和值: %T", str)
	fmt.Println(formattedT)
	// %% 输出 % 本体
	formattedPercent := fmt.Sprintf("输出 %% 本体: %v%%", num)
	fmt.Println(formattedPercent)
	// %b 整型以二进制方式显示
	formattedBinary := fmt.Sprintf("二进制: %b", num)
	fmt.Println(formattedBinary)
	// %o 整型以八进制方式显示
	formattedOctal := fmt.Sprintf("八进制: %o", num)
	fmt.Println(formattedOctal)
	// %d 整型以十进制方式显示
	formattedDecimal := fmt.Sprintf("十进制: %d", num)
	fmt.Println(formattedDecimal)
	// %x 整型以十六进制方式显示
	formattedHex := fmt.Sprintf("十六进制: %x", num)
	fmt.Println(formattedHex)
	// %X 整型以十六进制、字母大写方式显示
	formattedHexUpper := fmt.Sprintf("十六进制（大写）: %X", num)
	fmt.Println(formattedHexUpper)
	// %U Unicode 字符
	formattedUnicode := fmt.Sprintf("Unicode 字符: %U", str)
	fmt.Println(formattedUnicode)
	// %f 浮点数
	formattedFloat := fmt.Sprintf("浮点数: %f", float32(num))
	fmt.Println(formattedFloat)
	// %p 指针，十六进制方式显示
	formattedPtr := fmt.Sprintf("指针: %p", &num)
	fmt.Println(formattedPtr)
}
func Sum(a string, b string) {
	// write code here
	/*
		给定两个字符串类型的数字，求这两个数字之和，并返回字符串形式。
		知识点：
		golang strconv.ParseInt 是将字符串转换为数字的函数,
		参数1 数字的字符串形式,
		参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制,
		参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	*/
	int_a, a_err := strconv.Atoi(a)
	int_b, b_err := strconv.Atoi(b)
	if a_err == nil && b_err == nil {
		result := strconv.Itoa(int_a + int_b)
		fmt.Println("返回相加的和：")
		fmt.Println(result)
	}
}
func Perimeter(a int, b int) {
	// write code here
	/*
		已知一个长方形的长，宽，求这个长方形的周长，周长=2*(长+宽)。
		知识点：
		golang中, '算术运算符 '*','+',分别表示乘法，加法
		golang中，括号()的运算优先级高于 * ，+ ，而*的运算优先级又高于+
	*/
	fmt.Println((a + b) * 2)
}
func BackList() {
	// write code here
	/*
		小明投了5次保龄球，每次的分数分别为2，5，4，6，5,用一个数组记录这5次分数，然后输出这个数组。
		知识点 ：
		数组：是同一种数据类型的固定长度的序列。
		数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
		长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
	*/
	data := []int{2, 5, 4, 6, 5}
	fmt.Println(data)
}

//func Makeslice(length int, capacity int) []int {
//	// write code here
//}
