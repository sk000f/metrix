package num

import "math"

func Trunc2dp(f float64) float64 {
	return math.Floor(f*100) / 100
}
