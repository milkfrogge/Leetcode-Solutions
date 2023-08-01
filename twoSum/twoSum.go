package twoSum

func TwoSum(nums []int, target int) []int {
	passed := make(map[int]int) // значение - индекс в массиве
	for i := 0; i < len(nums); i++ {
		if idx, ok := passed[(target - nums[i])]; ok {
			return []int{idx, i}
		} else {
			passed[nums[i]] = i
		}
	}
	return nil
}
