package main

import "fmt"

func removeDuplicates(nums []int) (int, []int) {
	//如果数组长度为0意味着数组为空，直接返回
	if len(nums) == 0 {
		return 0, []int{}
	}
	//采用双指针比较，开始默认都指向第一个元素
	i, j := 0, 0
	//循环判断，只要没有指向最后一个元素就继续处理
	for i < len(nums)-1 {
		//如果当前两个指针指向的元素相等，一个指针就向后移动
		for nums[j] == nums[i] {
			j++
			//如果指针指向最后一个元素就直接返回数组长度和新数组
			if j == len(nums) {
				return i + 1, nums[:i+1]
			}
		}
		//如果两个指针指向的数组元素不相等，将j指向的元素放到i指向元素的下一个位置上，然后i向后移动
		nums[i+1] = nums[j]
		i++
	}
	//返回新数组与其长度
	return i + 1, nums[:i+1]
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 5, 6, 7, 7, 7, 8, 9}
	length, newNums := removeDuplicates(nums)
	fmt.Printf("新数组：%v,新数组长度：%v", newNums, length)
}
