package arrays

func plusOne(digits []int) []int {
	y := 1
	integer := make([]int, len(digits), len(digits)+1)
	for i := len(digits) - 1; i >= 0; i-- {
		x := (digits[i] + y) % 10
		y = (digits[i] + y) / 10
		integer[i] = x
	}
	if y > 0 {
		integer = append([]int{y}, integer...)
	}
	return integer
}
