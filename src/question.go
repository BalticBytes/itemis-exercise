package main

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	assignmentIndicator = "much"
	creditsIndicator    = "many Credits"
	invalidIndicator    = "not a question"
)

// There is no "true" enum type so keyword is just a string
type Question struct {
	keyword string // much | many
	amount  string
	unit    *string // optional
}

func NewQuestion(line string) *Question {
	// replace all occurences of 1..N whitespaces with single whitespace
	line = strings.Join(WhitespaceRegex.Split(line, -1), " ")

	if matches := HowMuchRegex.FindStringSubmatch(line); matches != nil {

		amount := strings.TrimSpace(matches[1])
		return &Question{assignmentIndicator, amount, nil}

	} else if matches = HowManyRegex.FindStringSubmatch(line); matches != nil {

		amount, unit := strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2])
		return &Question{creditsIndicator, amount, &unit}

	} else {
		// handle unknown question type
		return &Question{invalidIndicator, "", nil}
	}
}

func (q Question) Equal(other interface{}) bool {
	return reflect.DeepEqual(q, other)
}

func (q Question) String() string {
	unit := "nil"
	if q.unit != nil {
		unit = *q.unit
	}
	// return fmt.Sprintf("{How %s is %s %s ?}", q.keyword, q.amount, unit)
	return fmt.Sprintf("Question{keyword:'%s', amount:'%s', unit:'%s'}", q.keyword, q.amount, unit)
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
	case assignmentIndicator:
		// Calculate Roman Numeral
		_, amt := translate(numeralById, q.amount)
		return fmt.Sprintf("%s is %d\n", q.amount, amt)
	case creditsIndicator:
		if q.unit != nil {
			// Calculate credits for unit
			credits := q.calculate(numeralById, unitCostByUnit)
			if 0 < credits {
				return fmt.Sprintf("%s %s is %d Credits\n", q.amount, *q.unit, credits)
			}
		}
		// well-formed questions about non-existent units are unanswerable
		fallthrough
	default:
		return "I have no idea what you are talking about\n"
	}
}
