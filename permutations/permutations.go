package permutations

func permutate(nums []int, size int, out *[][]int) {
	if size == 1 {
		copyArray := make([]int, len(nums))
		copy(copyArray, nums)
		*out = append(*out, copyArray)
		return
	}

	for i := 0; i < size; i++ {

		permutate(nums, size-1, out)

		if size%2 == 1 {
			nums[0], nums[size-1] = nums[size-1], nums[0]
		} else {
			nums[i], nums[size-1] = nums[size-1], nums[i]
		}
	}
	return
}

func Permute(nums []int) [][]int {
	a := make([][]int, 0)
	permutate(nums, len(nums), &a)
	return a
}
