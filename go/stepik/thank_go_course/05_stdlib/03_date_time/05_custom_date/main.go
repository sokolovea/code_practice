package main

import (
	"errors"
	"fmt"
	"time"
)

// начало решения

// TimeOfDay описывает время в пределах одного дня
type TimeOfDay struct {
	hour     int
	minute   int
	second   int
	location *time.Location
}

// Hour возвращает часы в пределах дня
func (t TimeOfDay) Hour() int {
	return t.hour
}

// Minute возвращает минуты в пределах часа
func (t TimeOfDay) Minute() int {
	return t.minute
}

// Second возвращает секунды в пределах минуты
func (t TimeOfDay) Second() int {
	return t.second
}

func (t TimeOfDay) getCurrentSecondCounter() int {
	return t.Hour()*3600 + t.Minute()*60 + t.Second()
}

// String возвращает строковое представление времени
// в формате чч:мм:сс TZ (например, 12:34:56 UTC)
func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02v:%02v:%02v %v", t.Hour(), t.Minute(), t.Second(), t.location.String())
}

// Equal сравнивает одно время с другим.
// Если у t и other разные локации - возвращает false.
func (t TimeOfDay) Equal(other TimeOfDay) bool {
	if t.location.String() != other.location.String() {
		return false
	}
	return t.Hour() == other.Hour() && t.Minute() == other.Minute() && t.Second() == other.Second()
}

// Before возвращает true, если время t предшествует other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.location.String() != other.location.String() {
		return false, errors.New("not implemented")
	}
	return t.getCurrentSecondCounter() < other.getCurrentSecondCounter(), nil
}

// After возвращает true, если время t идет после other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.location.String() != other.location.String() {
		return false, errors.New("not implemented")
	}
	return t.getCurrentSecondCounter() > other.getCurrentSecondCounter(), nil
}

// MakeTimeOfDay создает время в пределах дня
func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	return TimeOfDay{hour: hour, minute: min, second: sec, location: loc}
}

// конец решения
