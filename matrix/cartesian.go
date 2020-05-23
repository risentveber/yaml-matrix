package matrix

func cartN(optionArrays ...[]interface{}) [][]interface{} {
	combinationsCount := 1
	for _, a := range optionArrays {
		combinationsCount *= len(a)
	}
	if combinationsCount == 0 {
		return nil
	}
	combinationLen := len(optionArrays)
	result := make([][]interface{}, combinationsCount)
	preAllocation := make([]interface{}, combinationsCount*combinationLen)
	indexAccumulator := make([]int, combinationLen)
	offset := 0
	for i := range result {
		nextOffset := offset + len(optionArrays)
		combination := preAllocation[offset:nextOffset]
		result[i] = combination
		offset = nextOffset

		for j, n := range indexAccumulator {
			combination[j] = optionArrays[j][n]
		}

		for j := len(indexAccumulator) - 1; j >= 0; j-- {
			indexAccumulator[j]++
			if indexAccumulator[j] < len(optionArrays[j]) {
				break
			}
			indexAccumulator[j] = 0
		}
	}

	return result
}
