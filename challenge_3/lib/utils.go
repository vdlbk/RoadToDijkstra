package lib

func IsMatrixEqual(m1, m2 [][]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for i, _ := range m1 {
		if !isArrayEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func isArrayEqual(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i, _ := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func CreateMatrix(items ...[]int) [][]int {
	a := make([][]int, 0, len(items))
	return append(a, items...)
}
