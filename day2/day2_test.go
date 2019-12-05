package day2

import "testing"

func TestDay2(t *testing.T) {
	tests := []struct {
		Name           string
		Input          []int
		ExpectedOutput int
	}{
		{
			"1,9,10,3,2,3,11,0,99,30,40,50",
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			3500,
		},
		{
			"1,0,0,0,99",
			[]int{1, 0, 0, 0, 99},
			2,
		},
		{
			"1,1,1,4,99,5,6,0,99",
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			Run(tt.Input)
		})
	}
}
