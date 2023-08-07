package findSubstring

// TODO: maybe some speed up?
func FindSubstring(s string, words []string) []int {
	out := make([]int, 0)
	t := make(map[string]int, len(words))
	for i := range words {
		if _, ok := t[words[i]]; ok {
			t[words[i]] = t[words[i]] + 1
		} else {
			t[words[i]] = 1
		}

	}
	size := len(words[0]) * len(words)
	sizeSingle := len(words[0])

	for i := 0; i <= len(s)-size; i++ {
		temp := make(map[string]int)
		for k, v := range t {
			temp[k] = v
		}
		for j := i; j <= len(s); j += sizeSingle {
			if len(temp) == 0 {
				out = append(out, i)
				break
			}
			if _, ok := temp[s[j:j+sizeSingle]]; ok {
				if temp[s[j:j+sizeSingle]] == 1 {
					delete(temp, s[j:j+sizeSingle])
				} else {
					temp[s[j:j+sizeSingle]] = temp[s[j:j+sizeSingle]] - 1
				}

			} else {
				break
			}
		}

	}

	return out
}
