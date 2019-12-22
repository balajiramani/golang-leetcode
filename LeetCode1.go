package main

func twoSum(nums []int, target int) []int {
	var result []int
	var m = make(map[int]int)
	for i:=0;i<len(nums);i++ {
		var num = nums[i]
		if val, ok := m[target - num]; ok {
			result = []int{val,i}
			break
		}
		m[num] = i
	}
	return result
}
