package num

import "math"

func To2dp(f float64) float64 {
	return math.Floor(f*100) / 100
}
