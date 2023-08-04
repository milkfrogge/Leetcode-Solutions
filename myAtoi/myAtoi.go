package myAtoi

import (
	"math"
	"strings"
)

func MyAtoi(s string) int {
	s = strings.Trim(s, " ")
	mul := 1

	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' {
		mul = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	res := 0

	for _, v := range s {
		if !(v >= '0' && v <= '9') {
			return res * mul
		}
		res = res*10 + int(v-'0')

		if res*mul >= math.MaxInt32 {
			return math.MaxInt32
		}
		if res*mul <= math.MinInt32 {
			return math.MinInt32
		}

	}
	return res * mul
}
