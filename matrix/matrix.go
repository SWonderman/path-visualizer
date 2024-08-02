package matrix

func GenerateEmptyMatrix(rows int32, columns int32) *[][]byte {
	matrix := make([][]byte, rows)
	for i := range matrix {
		matrix[i] = make([]byte, columns)
	}

	for r := range rows {
		for c := range columns {
			matrix[r][c] = '-'
		}
	}

	return &matrix
}
