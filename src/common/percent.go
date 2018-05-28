package common

import (
	"fmt"
)

func Percent(divisorA float64, divisorB float64) string {
	v := divisorA / divisorB
	if v < 1 {
		v = 1 - v
		v = v * 100
		return "-" + Decimal(v)
	}
	v = v - 1
	return Decimal(v * 100)
}

func Decimal(v float64) string {
	//return math.Trunc(v*1e2) * 1e-2
	return fmt.Sprintf("%.2f", v)
}
