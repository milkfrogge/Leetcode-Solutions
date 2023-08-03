package romanToInt

func RomanToInt(s string) int {
	out := int(0)
	symbols := map[uint8]int{73: 1, 86: 5, 88: 10, 76: 50, 67: 100, 68: 500, 77: 1000}
	for i := 0; i < len(s); {
		if i == len(s)-1 {
			out = out + symbols[s[i]]
			break
		}
		left, right := s[i], s[i+1]
		if symbols[left] >= symbols[right] {
			out = out + symbols[left]
			i++
		} else {
			out = out + symbols[right] - symbols[left]
			i += 2
		}
	}
	return out
}
