package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculatePacksNeeded(t *testing.T) {
	app := &App{}
	tests := []struct {
		orderQuantity int
		packSizes     []int
		expected      map[int]int
		expectError   bool
	}{
		{
			orderQuantity: 4,
			packSizes:     []int{5, 10},
			expected:      nil,
			expectError:   true,
		},
		{
			orderQuantity: 0,
			packSizes:     []int{3, 5, 7},
			expected:      nil,
			expectError:   true,
		},
		{
			orderQuantity: -5,
			packSizes:     []int{3, 5, 7},
			expected:      nil,
			expectError:   true,
		},
		{
			orderQuantity: 500000,
			packSizes:     []int{23, 31, 53},
			expected:      map[int]int{23: 2, 31: 7, 53: 9429},
			expectError:   false,
		},
	}

	for _, tt := range tests {
		result, err := app.CalculatePacksNeeded(tt.orderQuantity, tt.packSizes)

		if tt.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.Equal(t, tt.expected[53], result[53])
		}
	}
}
