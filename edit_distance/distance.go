package distance

func Min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

// EditDistance computes the minimal number of operations "add", "delete", and "replace" to transform `s1` into `s2`
func EditDistance(s1, s2 string) int {
	len1 := len(s1)
	len2 := len(s2)
	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}
	prev := make([]int, len2+1)
	for j := 0; j <= len2; j++ {
		prev[j] = j
	}
	curr := make([]int, len2+1)
	for i := 1; i <= len1; i++ {
		curr[0] = i
		for j := 1; j <= len2; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost += 1
			}
			curr[j] = Min(prev[j]+1, curr[j-1]+1, prev[j-1]+cost)
		}
		prev, curr = curr, prev
	}
	return prev[len2]
}
