package must_test

import (
	"errors"
	"testing"

	"github.com/doomshrine/must"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDo(t *testing.T) {
	t.Parallel()

	// test different inputs
	for _, tc := range []struct {
		name      string
		input     func() (any, error)
		assertion func(assert.TestingT, assert.PanicTestFunc, ...interface{}) bool
	}{
		{
			name:      "panics",
			assertion: assert.Panics,
			input: func() (any, error) {
				return nil, errors.New("expecting panic")
			},
		},
		{
			name:      "input string",
			assertion: assert.NotPanics,
			input: func() (any, error) {
				return "test", nil
			},
		},
		{
			name:      "input int",
			assertion: assert.NotPanics,
			input: func() (any, error) {
				return 1, nil
			},
		},
		{
			name:      "input struct",
			assertion: assert.NotPanics,
			input: func() (any, error) {
				return struct{}{}, nil
			},
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.NotNil(t, tc.input)
			require.NotNil(t, tc.assertion)
			tc.assertion(t, func() {
				_ = must.Do(tc.input())
			})
		})
	}
}
