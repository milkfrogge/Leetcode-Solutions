package isPalindrome

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	c := 0
	prevPow := 1
	for i := 1; true; i++ {
		pow := prevPow * 10

		if prevPow > x {
			break
		}
		c = c*10 + (x%pow)/prevPow
		prevPow = pow
	}
	return c == x
}
