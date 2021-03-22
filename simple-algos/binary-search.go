package algos

/* Find element in sorted array. */
func BinarySearch(nums [5]int, target int) int {
	left, right := 0, len(nums)-1

	for left < right { //equivalent of while loop
		cur := (left + right) / 2
		if nums[cur] < target {
			left = cur + 1
		} else if nums[cur] > target {
			right = cur - 1
		} else {
			return cur
		}
	}
	return -1
}
