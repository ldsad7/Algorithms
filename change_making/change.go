package tree

import (
	"math"
)

// Change amount with coins
func Change(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	var arr []int
	arr = append(arr, 0)
	for i := 1; i <= amount; i++ {
		arr = append(arr, math.MaxUint32)
	}
	for i := 1; i <= amount; i++ {
		for j, coin := range coins {
			if coin <= i && coin > 0 {
				if arr[i-coins[j]] != math.MaxUint32 {
					tmp := arr[i-coins[j]] + 1
					if arr[i] > tmp {
						arr[i] = tmp
					}
				}
			}
		}
	}
	if arr[amount] < math.MaxUint32 {
		return arr[amount]
	}
	return -1
}
