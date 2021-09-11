package decimal

import "github.com/shopspring/decimal"

func RoundFloat64(value float64, places int32) (f float64, exact bool) {
	return decimal.NewFromFloat(value).Round(places).Float64()
}
