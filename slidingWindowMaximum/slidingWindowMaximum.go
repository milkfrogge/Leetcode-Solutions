package slidingWindowMaximum

func MaxSlidingWindow(nums []int, k int) []int {
	var queue []int
	var result = make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {

		// Remove all elements smaller than or equal to nums[i]
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		// Add i to queue
		queue = append(queue, i)

		// Skip reducing the window size if i is less than window size
		if i < k-1 {
			continue
		}

		// Delete all elements whose index is out of window range
		for len(queue) > 0 && queue[0] < i-k+1 {
			queue = queue[1:]
		}
		// First element is the largest element
		result[i-k+1] = nums[queue[0]]
	}
	return result
}
