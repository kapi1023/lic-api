package levenshtein

// Funkcja obliczająca odległość Levenshteina między dwoma ciągami znaków
func distance(s1, s2 string) int {
	m, n := len(s1), len(s2)
	// Inicjalizacja macierzy odległości
	matrix := make([][]int, m+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
	}

	// Inicjalizacja pierwszego wiersza i pierwszej kolumny
	for i := 0; i <= m; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= n; j++ {
		matrix[0][j] = j
	}

	// Obliczanie odległości Levenshteina
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}
			matrix[i][j] = min(min(matrix[i-1][j]+1, matrix[i][j-1]+1), matrix[i-1][j-1]+cost)
		}
	}

	return matrix[m][n]
}
