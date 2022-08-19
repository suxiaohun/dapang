package lc

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	index := 0
	count := 0
	p1 := startTime[index]
	p2 := endTime[index]
	for {
		if queryTime >= p1 && queryTime <= p2 {
			count += 1
		}
		index += 1
		if index > len(startTime)-1 {
			break
		}
		p1 = startTime[index]
		p2 = endTime[index]
	}
	return count
}
