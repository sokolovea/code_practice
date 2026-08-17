package main

import (
	"strings"
	"testing"
)

func TestTaskExample(t *testing.T) {
	testData := struct {
		records      []string
		resultReport []string
		err          error
	}{
		[]string{
			"15.04.2022",
			"8:00 - 8:30 Завтрак",
			"8:30 - 9:30 Оглаживание кота",
			"9:30 - 10:00 Интернеты",
			"10:00 - 14:00 Напряженная работа",
			"14:00 - 14:45 Обед",
			"14:45 - 15:00 Оглаживание кота",
			"15:00 - 19:00 Напряженная работа",
			"19:00 - 19:30 Интернеты",
			"19:30 - 22:30 Безудержное веселье",
			"22:30 - 23:00 Оглаживание кота",
		},
		[]string{
			// "Мои достижения за 2022-04-15",
			"- Напряженная работа: 8h0m0s",
			"- Безудержное веселье: 3h0m0s",
			"- Оглаживание кота: 1h45m0s",
			"- Интернеты: 1h0m0s",
			"- Обед: 45m0s",
			"- Завтрак: 30m0s",
		},
		nil,
	}
	testRecords := strings.Join(testData.records, "\n")
	tasks, err := ParsePage(testRecords)
	if err != nil {
		t.Errorf("Error: want = %v, got = %v", nil, err)
	}
	for i, task := range tasks {
		if task.String() != testData.resultReport[i] {
			t.Errorf("Task error: want = %v, got = %v", testData.resultReport[i], task)
		}
	}
}
