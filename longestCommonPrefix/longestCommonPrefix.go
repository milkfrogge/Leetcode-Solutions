package longestCommonPrefix

func LongestCommonPrefix(strs []string) string {

	end := false
	out := ""

	for i := range strs[0] {
		if end {
			break
		}

		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				end = true
				break
			}

			temp := strs[0][i] ^ strs[j][i]
			if temp != 0 {
				end = true
				break
			}

		}

		if !end {
			out += string(strs[0][i])
		}

	}
	return out
}
