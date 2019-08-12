package phone

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPhoneNumber(t *testing.T) {

	t.Run("Validate phone number ", func(t *testing.T) {
		tt := []struct {
			input    string
			expected bool
		}{
			{"351960000000", true},
			{"00351961111111", true},
			{"+00351961111111", false},
			{"+ 351961111111", false},
			{"+351 96111 1111", true},
			{"351t61111111", false},
			{"351+61111111", false},
			{"351611p1111", false},
			{"+123", true},
			{"+00123", false},
			{"0012", false},
		}

		for _, tc := range tt {
			_, b := Sanitize(tc.input)
			assert.Equal(t, b, tc.expected,
				fmt.Sprintf("number %s got result %v expected %v", tc.input, b, tc.expected),
			)
		}
	})
}
