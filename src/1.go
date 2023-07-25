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

// LengthOfLongestSubstring 2 lengthOfLongestSubstring
func LengthOfLongestSubstring(s string) int {
    
    if len(s) == 0 {
        return 0
    }
    freq := make([]int, 255)
    ans := 1
    freq[s[0]] = 1
    max := 1
    for start , end := 0, 1; end < len(s); end++ {
        if freq[s[end]] > 0 && freq[s[end]] >= start {
            start = freq[s[end]]
            freq[s[end]] = end + 1
            max = end - start + 1
            ans = maxi(ans, max)
            continue
        }
        freq[s[end]] = end+1
        max++
        ans = maxi(ans, max)
        
    }
    
    return ans
}

func maxi(x,y int) int {
    if x > y {
        return x
    }
    return y
}
