package main

import (
	"testing"
)

func TestNewQuestion(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want Question
	}{
		// TODO: Add test cases.
		{
			"assignment question",
			args{"how much is x y z ?"},
			Question{assignmentIndicator, "x y z", nil},
		},
	}
	for _, tt := range tests {
		got := NewQuestion(tt.args.line)
		if !got.Equal(tt.want) {
			t.Errorf("NewQuestion() = %v, want %v", got, tt.want)
		}
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
