package main

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])
	if err != nil {
		return []Task{}, err
	}
	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return []Task{}, err
	}
	sortTasks(tasks)
	return tasks, nil
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	return time.Parse("02.01.2006", src)
}

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	regTask := regexp.MustCompile(`(\d+:\d+) - (\d+:\d+) (.+)`)
	taskSlice := make([]Task, 0)
	taskMap := make(map[string]Task)
	for _, record := range lines {
		currentTask := Task{}
		recordGroups := regTask.FindStringSubmatch(record)
		if len(recordGroups) != 4 {
			return []Task{}, errors.New("can't parse single task!")
		}
		currentTask.Date = date
		var err error
		startTime, err := time.Parse("15:04", recordGroups[1])
		if err != nil {
			return []Task{}, errors.New("can't parse task startTime!")
		}
		endTime, err := time.Parse("15:04", recordGroups[2])
		if err != nil {
			return []Task{}, errors.New("can't parse task endTime!")
		}
		if endTime.Sub(startTime) <= 0 {
			return []Task{}, errors.New("endTime must be > than startTime!")
		}
		currentTask.Dur = endTime.Sub(startTime)
		currentTask.Title = recordGroups[3]
		existingTask, ok := taskMap[currentTask.Title]
		if !ok {
			taskMap[currentTask.Title] = currentTask
		} else {
			existingTask.Dur += currentTask.Dur
			taskMap[currentTask.Title] = existingTask
		}
	}
	for _, value := range taskMap {
		taskSlice = append(taskSlice, value)
	}
	return taskSlice, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	slices.SortFunc(tasks, func(left Task, right Task) int {
		if left.Dur < right.Dur {
			return 1
		} else if left.Dur > right.Dur {
			return -1
		} else {
			return 0
		}
	})
}

// конец решения

const (
	fmtStrHour    = "- %v: %vh%vm%vs"
	fmtStrMinutes = "- %v: %vm%vs"
	fmtStrSeconds = "- %v: %vs"
)

func (t Task) String() string {
	if int(t.Dur.Hours()) != 0 {
		return fmt.Sprintf(fmtStrHour, t.Title, int(t.Dur.Hours()), int(t.Dur.Minutes())%60, int(t.Dur.Seconds())%60)
	}
	if int(t.Dur.Minutes()) != 0 {
		return fmt.Sprintf(fmtStrMinutes, t.Title, int(t.Dur.Minutes())%60, int(t.Dur.Seconds())%60)
	}
	return fmt.Sprintf(fmtStrSeconds, t.Title, int(t.Dur.Seconds())%60)
}
