package app

import (
	"fmt"
	"math"
	"sort"
)

//	This function calculates the minimum number of packs needed to fulfill order(its quantity metric).
//
// If it's not possible to calculate order using the given package sizes -> error would be returned.
func (a *App) CalculatePacksNeeded(orderQuantity int, packSizes []int) (map[int]int, error) {
	if orderQuantity <= 0 {
		return nil, fmt.Errorf("order quantity must be a positive integer")
	}
	// The package sizes are sorted in DESC order => larger packs are considered first.
	// This is for optimisation strategy
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	//stores the min number of packs needed to fulfill order of [i]-items.
	dp := make([]int, orderQuantity+1)

	//stores the last chosen package size to make up [i] items.
	choice := make([]int, orderQuantity+1)

	for i := 1; i <= orderQuantity; i++ {

		//initialize with large number (infinity)
		dp[i] = math.MaxInt32
		for _, pack := range packSizes {
			if i >= pack && dp[i-pack]+1 < dp[i] {

				// update with the minimum number of packs
				dp[i] = dp[i-pack] + 1
				// s tore the pack size used
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

		//'choice' -  which packs were used to fulfill the order
		//'remaining' - track of how much of the order is still left to be packed
		pack := choice[remaining]
		result[pack]++
		remaining -= pack
	}

	return result, nil
}
