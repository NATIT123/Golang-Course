package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for index, value := range nums {
		number := target - value
		v, ok := hashMap[number]
		if ok {
			return []int{index, v}
		}
		hashMap[value] = index
	}
	return []int{}
}
