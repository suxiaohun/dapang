package lc

import "math"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	strHash := make(map[string]int, 0)
	start := 0

	for i, c := range s {
		v, ok := strHash[string(c)]
		if ok && v >= start {
			tempLen := i - start
			if maxLen < tempLen {
				maxLen = tempLen
			}
			start = v + 1
		}
		strHash[string(c)] = i
	}
	tempLen := len(s) - start
	if maxLen < tempLen {
		maxLen = tempLen
	}
	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	start := float64(0)
	uniqCount := float64(0)
	last := make(map[string]float64, 0)

	for i := 0; i < len(s); i++ {
		last_index, ok := last[s[i:i+1]]
		if !ok {
			last_index = -1
		}
		start = math.Max(start, last_index+1)

		uniqCount = math.Max(uniqCount, float64(i)-start+1)
		last[s[i:i+1]] = float64(i)
	}
	return int(uniqCount)
}
