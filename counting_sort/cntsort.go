package cntsort

// CountingSort performs a linear-time in-place sorting of not-so-large integers
func CountingSort(values []int) {
	n := len(values)
	max := 0
	min := 0
	tmp := n
	for tmp != 0 {
		tmp--
		if values[tmp] >= max {
			max = values[tmp]
		}
		if values[tmp] <= min {
			min = values[tmp]
		}
	}
	arr := make([]int, max-min+1)
	tmp = n
	for tmp != 0 {
		tmp--
		arr[values[tmp]-min] += 1
	}
	tmp = 0
	n = 0
	for n <= max-min {
		for arr[n] != 0 {
			arr[n]--
			values[tmp] = n + min
			tmp++
		}
		n++
	}
}
