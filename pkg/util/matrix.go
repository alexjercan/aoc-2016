package util

func Transpose[K any](matrix [][]K) [][]K {
	result := make([][]K, len(matrix[0]))
	for i := range result {
		result[i] = make([]K, len(matrix))
	}

	for i := range matrix {
		for j := range matrix[i] {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}
