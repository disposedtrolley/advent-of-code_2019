package day2

import "testing"

import "github.com/stretchr/testify/assert"

func TestDay2(t *testing.T) {
	tests := []struct {
		Name           string
		Input          string
		ExpectedOutput int
	}{
		{
			"1,9,10,3,2,3,11,0,99,30,40,50",
			"1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50",
			3500,
		},
		{
			"1,0,0,0,99",
			"1, 0, 0, 0, 99",
			2,
		},
		{
			"1,1,1,4,99,5,6,0,99",
			"1, 1, 1, 4, 99, 5, 6, 0, 99",
			30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			result := Run(tt.Input)

			assert.Equal(t, tt.ExpectedOutput, result)
		})
	}
}
