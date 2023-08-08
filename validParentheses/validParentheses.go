package validParentheses

var closes = map[uint8]uint8{
	')': '(',
	'}': '{',
	']': '[',
}

func IsValid(str string) bool {
	s := []uint8(str)
	stack := make([]uint8, 0)

	for _, symbol := range s {
		opening, closing := closes[symbol]

		if !closing {
			stack = append(stack, symbol)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		lastOpening := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if lastOpening != opening {
			return false
		}
	}

	return len(stack) == 0
}
