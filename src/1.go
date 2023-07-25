package src

// TwoSum 1 twoSum
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		if v, ok := m[tmp]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
