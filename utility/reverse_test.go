package utility_test

import (
	"testing"

	"github.com/data-love/neo-go-sdk/utility"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Run("Reverse()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				in          string
				out         string
			}{
				{
					description: "Empty",
					in:          "",
					out:         "",
				},
				{
					description: "Neo",
					in:          "Neo",
					out:         "oeN",
				},
				{
					description: "CityOfZion",
					in:          "CityOfZion",
					out:         "noiZfOytiC",
				},
			}
			for _, testCase := range testCases {
				t.Run(testCase.description, func(t *testing.T) {
					result := utility.Reverse(testCase.in)
					assert.Equal(t, testCase.out, result)
				})
			}
		})

		t.Run("SadCase", func(t *testing.T) {
			testCases := []struct {
				description string
				in          string
				out         string
			}{
				{
					description: "Case Insensitive",
					in:          "neo",
					out:         "oeN",
				},
				{
					description: "Wrong input",
					in:          "CityOfZion",
					out:         "noiZf0ytiC",
				},
			}
			for _, testCase := range testCases {
				t.Run(testCase.description, func(t *testing.T) {
					result := utility.Reverse(testCase.in)
					assert.NotEqual(t, testCase.out, result)
				})
			}
		})
	})
}
