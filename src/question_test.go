package main

import (
	"math/rand"
	"strings"
	"testing"
	"unicode"
)

// replaces a substring with randomised case.
func randomCase(src string, toRandomise string) string {

	sb := strings.Builder{}
	parts := strings.Split(src, toRandomise)

	sb.WriteString(parts[0])
	for _, r := range toRandomise {
		if rand.Intn(2) == 0 {
			sb.WriteString(strings.ToUpper(string(r)))
		} else {
			sb.WriteString(strings.ToLower(string(r)))
		}
	}
	sb.WriteString(parts[1])

	return sb.String()
}

// replaces whitespace with an increased amount of random whitespace.
func randomPad(src string, min int, max int) string {
	sb := strings.Builder{}

	for _, r := range src {
		if unicode.IsSpace(r) {
			count := min + rand.Intn(max-min+1)
			for i := 0; i < count; i++ {
				sb.WriteRune(' ')
			}
		} else {
			sb.WriteString(string(r))
		}
	}

	return sb.String()
}

func TestNewCreditsQuestion(t *testing.T) {
	metalUnit := "Metal"
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Question
	}{
		{
			"credits question: ignore case",
			args{"how many Credits is x Metal?"},
			Question{creditsIndicator, "x", &metalUnit},
		},
		{
			"credits question: 1 var",
			args{"how many Credits is x Metal?"},
			Question{creditsIndicator, "x", &metalUnit},
		},
		{
			"credits question: 2 vars",
			args{"how many Credits is x y Metal?"},
			Question{creditsIndicator, "x y", &metalUnit},
		},
		{
			"credits question: 3 vars",
			args{"how many Credits is x y z Metal?"},
			Question{creditsIndicator, "x y z", &metalUnit},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := randomPad(randomCase(tt.args.line, "Credits"), 1, 3)
			got := NewQuestion(input)
			if !got.Equal(tt.want) {
				t.Errorf("NewQuestion = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAssignmentQuestion(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Question
	}{
		{
			"assignment question",
			args{"how much is x y z ?"},
			Question{assignmentIndicator, "x y z", nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := randomPad(tt.args.line, 1, 3)
			got := NewQuestion(input)
			if !got.Equal(tt.want) {
				t.Errorf("actual: %v, expected: %v", got, tt.want)
			}
		})
	}
}

func TestNewInvalidQuestion(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Question
	}{
		{
			"invalid question: no unit",
			args{"how many Credits is one?"},
			Question{invalidIndicator, "", nil},
		},
		{
			"invalid question: mixed assignment and credits",
			args{"how much Credits is one Metal?"},
			Question{invalidIndicator, "", nil},
		},
		{
			"invalid question: lorem ipsum",
			args{"Lorem ipsum dolor sit amet, consectetur adipiscing elit."},
			Question{invalidIndicator, "", nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := randomPad(tt.args.line, 1, 3)
			got := NewQuestion(input)
			if !got.Equal(tt.want) {
				t.Errorf("actual: %v, expected: %v", got, tt.want)
			}
		})
	}
}
func TestQuestion_calculate(t *testing.T) {
	type args struct {
		m map[string]string
		n map[string]float64
	}
	tests := []struct {
		name    string
		q       Question
		args    args
		wantSum int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := tt.q.calculate(tt.args.m, tt.args.n); gotSum != tt.wantSum {
				t.Errorf("Question.calculate() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestQuestion_answer(t *testing.T) {
	type args struct {
		numeralById    map[string]string
		unitCostByUnit map[string]float64
	}
	tests := []struct {
		name       string
		q          Question
		args       args
		wantAnswer string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := tt.q.answer(tt.args.numeralById, tt.args.unitCostByUnit); gotAnswer != tt.wantAnswer {
				t.Errorf("Question.answer() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
