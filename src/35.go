package src

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	mid := 0
	for start <= end {
		mid = start + (end-start)/2
		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	if end < 0 {
		return 0
	}
	return start
}
