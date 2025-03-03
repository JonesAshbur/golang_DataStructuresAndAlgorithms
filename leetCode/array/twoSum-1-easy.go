package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tag := make(map[int]int) // 1. 初始化一个空的 map
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]    // 2. 计算另一个需要的数
		if _, ok := tag[another]; ok { // 3. 检查 another 是否在 map 中
			return []int{tag[another], i} // 4. 如果存在，返回两个索引
		}
		tag[nums[i]] = i // 5. 如果不存在，将当前数和索引存入 map
	}
	return nil // 6. 没有找到，返回 nil
}

func main() {
	// 给定切片
	nums := []int{2, 7, 11, 15}
	// 目标值
	target := 9
	// 期望结果 [0, 1]
	res := twoSum(nums, target)
	fmt.Println("res:", res)
}
