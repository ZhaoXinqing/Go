func minimunTotal(tringle [][]int) int {
	for i, v := range triangle {
		if i == 0 {
			continue
		}
		for idx, val := range v {
			if idx == len(v)-1 {
				triangle[i][idx] = val + triangle[i-1][idx-1]
				continue
			}
		}
	}

}