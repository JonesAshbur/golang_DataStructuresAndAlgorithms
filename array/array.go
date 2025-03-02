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

// 在数组的索引 index 处插入元素 num，丢失数组中最后一个元素,时间复杂度位O（n）n为数组长度
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

// 删除索引 index 处的元素,时间复杂度位O（n）n为数组长度
func remove(nums []int, index int) {
	// 把索引 index 之后的所有元素向前移动一位
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
		//最后的元素设为0，表示“无意义”
		nums[i+1] = 0
	}
	fmt.Println(nums)
}

// 遍历数组,并且对数组元素进行求和
func traverse(nums []int) (int, int, int) {
	// 遍历时对数组元素进行求和
	count1 := 0
	// 通过索引遍历数组
	for i := 0; i < len(nums); i++ {
		count1 += nums[i]
	}
	count2 := 0
	// 直接遍历数组元素
	for _, num := range nums {
		count2 += num
	}
	// 同时遍历数据索引和元素
	count3 := 0
	for i, num := range nums {
		count3 += nums[i]
		// or
		count3 += num
	}
	return count1, count2, count3 / 2
}

// 在数组中查找指定元素第一次出现的位置,时间复杂度为O（n），索引查找的时间复杂度为O（1）
func find(nums []int, target int) (index int) {
	index = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			index = i
			break
		}
	}
	return
}

// 扩展数组长度 时间复杂度为O（n），开辟新的更大的连续空间，将原数组拷贝到新数组
func extend(nums []int, enlarge int) []int {
	// 初始化一个扩展长度后的数组
	res := make([]int, len(nums)+enlarge)
	// 将原数组中的所有元素复制到新数组
	for i, num := range nums {
		res[i] = num
	}
	// 返回扩展后的新数组
	return res
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

	// 删除index位置的元素
	nums3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	remove(nums3, 0)

	//遍历数组
	num4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res1, res2, res3 := traverse(num4)
	fmt.Printf("通过索引遍历数组进行求和：%v, 直接遍历数组元素进行求和：%v,同时遍历数据索引和元素进行求和：%v\n", res1, res2, res3)

	//查找元素
	nums5 := []int{2, 4, 5, 2, 5, 4, 7}
	findRes := find(nums5, 2)
	fmt.Println("该元素第一次出现的位置：", findRes)

	//扩容数组
	nums6 := []int{2, 4, 5, 2, 5, 4, 7}
	extendRes := extend(nums6, 4)
	fmt.Printf("扩容后数组为：%v，扩容后数组长度为：%v", extendRes, len(extendRes))

}
