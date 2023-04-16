package main

import (
	"fmt"
	"strings"
)

// There is no "true" enum type so keyword is just a string
type Question struct {
	keyword string // much | many
	amount  string
	unit    *string // optional
}

func NewQuestion(line string) *Question {
	if matches := howMuchRegex.FindStringSubmatch(line); matches != nil {

		amount := strings.TrimSpace(matches[1])
		return &Question{"much", amount, nil}

	} else if matches = howManyRegex.FindStringSubmatch(line); matches != nil {

		amount, unit := strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2])
		return &Question{"many Credits", amount, &unit}

	} else {
		// handle unknown question type
		return &Question{"not a question", "", nil}
	}
}

func (q Question) String() string {
	unit := "N/A"
	if q.unit != nil {
		unit = *q.unit
	}
	return fmt.Sprintf("{How %s is %s %s ?}", q.keyword, q.amount, unit)
}

// returns the sum of an intergalactic
func (q Question) calculate(m map[string]string, n map[string]float64) (sum int) {
	if q.unit != nil {
		_, i := translate(m, q.amount)

		return int(float64(i) * n[*q.unit])
	}
	return 0
}

// returns an answer to the question
func (q Question) answer(numeralById map[string]string, unitCostByUnit map[string]float64) (answer string) {
	switch q.keyword {
	case "much":
		// Calculate Roman Numeral
		_, amt := translate(numeralById, q.amount)
		return fmt.Sprintf("%s is %d\n", q.amount, amt)
	case "many Credits":
		if q.unit != nil {
			// Calculate credits for unit
			credits := q.calculate(numeralById, unitCostByUnit)
			return fmt.Sprintf("%s %s is %d Credits\n", q.amount, *q.unit, credits)
		}
		fallthrough
	default:
		return "I have no idea what you are talking about"
	}
}
