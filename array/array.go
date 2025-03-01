package main

import (
	"fmt"
	"math/rand"
)

func randomAccess(nums []int) (randomNum int) {
	// 在区间 [0, nums.length) 中随机抽取一个数字
	randomIndex := rand.Intn(len(nums))
	// 获取并返回随机元素
	randomNum = nums[randomIndex]
	return
}

// 在数组的索引 index 处插入元素 num，丢失数组中最后一个元素
func insertNumFun1(nums []int, num int, index int) []int {
	// 把索引 index 以及之后的所有元素向后移动一位
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	// 将 num 赋给 index 处的元素
	nums[index] = num
	return nums
}

// 不丢失原来数组中的元素，通过扩容插入元素
func insertNumFun2(nums []int, num int, index int) []int {
	//增加数组长度
	newNums := make([]int, len(nums)+1)
	//将原有数组的元素从第一个到第index—1个复制到新的数组中
	copy(newNums[:index], nums[:index])
	//在第index的位置插入新元素
	newNums[index] = num
	//将第index个元素开始到最后的所有元素复制到新数组中
	copy(newNums[index+1:], nums[index:])
	return newNums
}

// 不丢失原来数组中的元素，通过扩容插入元素
func insertNumFun3(nums []int, num int, index int) []int {
	// 扩展切片长度，为新元素腾出空间
	nums = append(nums, 0) // 先追加一个 0，扩展切片
	// 将插入位置之后的元素向后移动
	copy(nums[index+1:], nums[index:])
	// 将新元素放入指定位置
	nums[index] = num
	return nums
}
func main() {

	//	所有数据结构都是基于数组、链表或二者的组合实现的
	//	基于数组可实现：栈、队列、哈希表、树、堆、图、矩阵、张量（维度 >=3 的数组）等。
	//	基于链表可实现：栈、队列、哈希表、树、堆、图等。
	// 初始化数组
	var arr [5]int // 在 Go 中，指定长度时（[5]int）为数组，不指定长度时（[]int）为切片
	fmt.Println("arr=", arr)
	// 由于 Go 的数组被设计为在编译期确定长度，因此只能使用常量来指定长度,或者[...]
	// 为了方便实现扩容 extend() 方法，以下将切片（Slice）看作数组（Array）
	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Println("nums=", nums1)
	// 访问数组元素的时间复杂度为O(1)
	// 随机访问元素
	randomNum := randomAccess(nums1)
	fmt.Println("randomNum=", randomNum)
	// 插入元素
	newNums1 := insertNumFun1(nums1, 6, 2)
	fmt.Println("newNums1=", newNums1)
	nums2 := []int{1, 2, 3, 4, 5}
	newNums2 := insertNumFun2(nums2, 6, 2)
	fmt.Println("newNums2=", newNums2)
	newNums3 := insertNumFun3(nums2, 6, 2)
	fmt.Println("newNums3=", newNums3)
}
