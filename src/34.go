package src

var leftIndex int
var rightIndex int

func searchRange(nums []int, target int) []int {

	leftIndex, rightIndex = -1, -1
	bsearch(nums, target, true)  //O(log(n))
	bsearch(nums, target, false) //O(log(n))

	return []int{leftIndex, rightIndex}
}

func Bsearch(nums []int, target int, isLeft bool) {
	start, end := 0, len(nums)-1
	mid := 0
	for start <= end {
		mid = start + (end-start)/2
		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			if isLeft {
				leftIndex = mid
				end = mid - 1
			} else {
				rightIndex = mid
				start = mid + 1
			}
		}
	}
}
