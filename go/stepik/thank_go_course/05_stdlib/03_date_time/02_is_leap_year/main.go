package main

import (
	"testing"
	"time"
)

// начало решения

func isLeapYear(year int) bool {
	return time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC).YearDay() == 366
}

// конец решения

func Test(t *testing.T) {
	if !isLeapYear(2020) {
		t.Errorf("2020 is a leap year")
	}
	if isLeapYear(2022) {
		t.Errorf("2022 is NOT a leap year")
	}
}
