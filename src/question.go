package main

import "fmt"

// There is no "true" enum type so keyword is just a string
type Question struct {
	keyword string // much | many
	amount  string
	unit    *string // optional
}

func (q Question) String() string {
	unit := "N/A"
	if q.unit != nil {
		unit = *q.unit
	}
	return fmt.Sprintf("{How %s is %s %s ?}", q.keyword, q.amount, unit)
}

// returns the sum of an intergalactic
func (q Question) calculate(m map[string]string, n map[string]float64) int {
	_, i := translate(m, q.amount)

	return int(float64(i) * n[*q.unit])
}
