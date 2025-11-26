package app

import (
	"fmt"
	"math"
	"sort"
)

// CalculatePacksNeeded calculates the minimum number of packs needed to fulfill an order.
// Uses dynamic programming to find the optimal combination of package sizes.
// Returns an error if it's not possible to fulfill the order with the given package sizes.
func (a *App) CalculatePacksNeeded(orderQuantity int, packSizes []int) (map[int]int, error) {
	if orderQuantity <= 0 {
		return nil, fmt.Errorf("order quantity must be a positive integer")
	}
	// Sort package sizes in descending order - larger packs are considered first for optimization
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	// dp stores the minimum number of packs needed to fulfill order of [i] items
	dp := make([]int, orderQuantity+1)

	// choice stores the last chosen package size to make up [i] items
	choice := make([]int, orderQuantity+1)

	for i := 1; i <= orderQuantity; i++ {
		// Initialize with large number (infinity)
		dp[i] = math.MaxInt32
		for _, pack := range packSizes {
			if i >= pack && dp[i-pack]+1 < dp[i] {
				// Update with the minimum number of packs
				dp[i] = dp[i-pack] + 1
				// Store the pack size used
				choice[i] = pack
			}
		}
	}

	if dp[orderQuantity] == math.MaxInt32 {
		return nil, fmt.Errorf("cannot fulfill order with given pack sizes")
	}

	result := make(map[int]int)
	remaining := orderQuantity
	for remaining > 0 {
		// choice - which packs were used to fulfill the order
		// remaining - track of how much of the order is still left to be packed
		pack := choice[remaining]
		result[pack]++
		remaining -= pack
	}

	return result, nil
}
