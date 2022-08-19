package lc

func twoSum(nums []int, target int) []int {
	_hash := make(map[int]int, 0)

	for i, num := range nums {
		temp := target - num
		_, ok := _hash[temp]
		if ok {
			return []int{i, _hash[temp]}
		}
		_hash[num] = i
	}
	return []int{}
}
