package golang_base_homework_task1

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
)

/*---------------------136-只出现一次的数字---------------------力扣-----------*/
func SingleNumber(nums []int) int {
	/*
		给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
		你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
	*/
	/*
		按位异或（^）运算有以下特性：
		任何数与 0 异或，结果是其本身，例如a ^ 0 = a。
		任何数与自身异或，结果为 0，即a ^ a = 0。！！！！！！！！很重要
	*/
	a := 0
	for _, n := range nums {
		a = a ^ n
	}
	return a
}

/*---------------------198. 打家劫舍---------------------------力扣-----------*/
func Rob(nums []int) int {
	/*你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
	影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
	给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。*/
	result_list := []int{0, 0} //第一列是奇数，第二列是偶数

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			fmt.Printf("%d 是偶数列\n", i)
			result_list[0] += nums[i]
		} else {
			fmt.Printf("%d 是奇数列\n", i)
			result_list[1] += nums[i]
		}
	}
	max_data := slices.Max(result_list)
	return max_data
}

/*---------------------21. 合并两个有序链表---------------------------力扣-----------*/
//type ListNode struct {
//
//}

//func MergeTwoLists(list1 *ListNode, list2 *) {
//	/*将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 */
//	if list1 == nil {
//		return list2
//	}
//	if list2 == nil {
//		return list1
//	}
//	if list1.Val < list2.Val {
//		list1.Next = MergeTwoLists(list1.Next, list2)
//		return list1
//	} else {
//		list2.Next = MergeTwoLists(list1, list2.Next)
//		return list2
//	}
//

func MergeTwoLists(list1 []int, list2 []int) {
	list := append(list1, list2...)
	sort.Sort(sort.IntSlice(list))
	fmt.Println(list)
}

/*---------------------46. 全排列--------------------------力扣-----------*/

//func permute(nums []int) [][]int {
//	var result [][]int
//	var current []int
//	used := make([]bool, len(nums))
//	backtrack(nums, current, used, &result)
//	return result
//}
//func backtrack(nums []int, current []int, used []bool, result *[][]int) {
//	/*给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。*/
//	if len(current) == len(nums) {
//
//	}
//
//	for _, v := range nums {
//		current_element := v //当前元素
//	}
//}
/*---------------------344. 反转字符串-------------------------力扣-----------*/
func ReverseString(s []byte) {
	/*编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
	不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。*/
	/*输入：s = ["h","e","l","l","o"]
	输出：["o","l","l","e","h"]*/
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

/*---------------------69. x 的平方根 -------------------------力扣-----------*/
func MySqrt(x int) int {
	/*给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
	由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
	注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
	*/
	/*输入：x = 4
	输出：2*/
	return int(math.Pow(float64(x), 0.5))
}

/*---------------------26. 删除有序数组中的重复项 -------------------------力扣-----------*/
func RemoveDuplicates(nums []int) []int {
	/*给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
	元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
	考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
	更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
	返回 k 。*/
	result := []int{}
	for _, v := range nums {
		if !slices.Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

/*---------------------56. 合并区间-------------------------力扣-----------*/
func Merge(intervals [][]int) [][]int {
	/*以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
	请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。*/
	/*
		输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
		输出：[[1,6],[8,10],[15,18]]
		解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
	*/
	return nil
}

/*---------------------430. 扁平化多级双向链表------------------------力扣-----------*/
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func Flatten(head string) {
	/*你会得到一个双链表，其中包含的节点有一个下一个指针、一个前一个指针和一个额外的 子指针 。
	这个子指针可能指向一个单独的双向链表，也包含这些特殊的节点。
	这些子列表可以有一个或多个自己的子列表，以此类推，以生成如下面的示例所示的 多层数据结构 。
	给定链表的头节点 head ，将链表 扁平化 ，以便所有节点都出现在单层双链表中。
	让 curr 是一个带有子列表的节点。子列表中的节点应该出现在扁平化列表中的 curr 之后 和 curr.next 之前 。
	返回 扁平列表的 head 。列表中的节点必须将其 所有 子指针设置为 null 。*/
	/*输入：head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
	输出：[1,2,3,7,8,11,12,9,10,4,5,6]*/
	//node := Node{Val: 1}
	head = strings.Trim(head, "[]")      // 去除首尾的方括号
	headList := strings.Split(head, ",") // 按逗号分割字符串
	dataList := make([][]string, 0)      // 存储分段后的数据
	stempList := make([]string, 0)       // 临时存储当前分段
	//currentIndex := 0                    // 初始化当前分段起始索引
	for _, v := range headList {
		stempList = append(stempList, v)
		if v == "null" {
			dataList = append(dataList, stempList)
			stempList = []string{}

		}
	}
}
func Flatten2(root *Node) {

}
