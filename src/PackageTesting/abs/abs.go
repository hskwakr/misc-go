package abs

func Abs(n int) int {
	if n < 0 {
		n = n * -1
	}
	return n
}
