package letterCombinations

func x(digits string, a map[uint8]string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	out := make([]string, 0)

	if len(digits) == 1 {
		for _, k := range a[digits[0]] {
			out = append(out, string(k))
		}
		return out
	}

	if len(digits) == 2 {
		for _, k := range a[digits[0]] {
			for _, l := range a[digits[1]] {
				out = append(out, string(k)+string(l))
			}
		}
		return out
	} else {
		temp := x(digits[1:], a)
		for _, k := range a[digits[0]] {
			for _, v := range temp {
				out = append(out, string(k)+v)
			}
		}

	}

	return out
}

func LetterCombinations(digits string) []string {

	a := make(map[uint8]string)
	a[50] = "abc"
	a[51] = "def"
	a[52] = "ghi"
	a[53] = "jkl"
	a[54] = "mno"
	a[55] = "pqrs"
	a[56] = "tuv"
	a[57] = "wxyz"
	aa := x(digits, a)
	return aa
}
