package num_test

import (
	"testing"

	"github.com/sk000f/metrix/internal/num"
)

func TestUnitNum(t *testing.T) {
	t.Run("test truncate float down to 2 decimal places", func(t *testing.T) {
		want := 0.25
		got := num.To2dp(0.2546)

		if got != want {
			t.Errorf("got: %v; want %v", got, want)
		}
	})

	t.Run("test truncate float up to 2 decimal places", func(t *testing.T) {
		want := 0.26
		got := num.To2dp(0.2645)

		if got != want {
			t.Errorf("got: %v; want %v", got, want)
		}
	})
}
