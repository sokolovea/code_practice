package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

// начало решения

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
	nanoseconds := t.UnixNano() % 1_000_000_000
	// for nanoseconds != 0 && nanoseconds%10 == 0 {
	// 	nanoseconds /= 10
	// }
	result := fmt.Sprintf("%v.%09v", t.Unix(), nanoseconds)
	for result[len(result)-1] == '0' && result[len(result)-2] != '.' {
		result = result[0 : len(result)-1]
	}
	return result
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
	regexpParse := regexp.MustCompile(`(\d+)\.(\d+)`)
	strDate := regexpParse.FindStringSubmatch(d)
	if len(strDate) != 3 {
		return time.Time{}, errors.New("Can't parse legacy date: wrong format!")
	}
	secondsDate, err := strconv.Atoi(strDate[1])
	if err != nil {
		return time.Time{}, err
	}
	// if secondsDate < 0 {
	// 	return time.Time{}, errors.New("Can't parse legacy date: can't be < 0!")
	// }
	nanosecondsDate, err := strconv.Atoi(strDate[2])
	if err != nil {
		return time.Time{}, err
	}
	dateFloat, err := strconv.ParseFloat(strDate[0], 64)
	if err != nil {
		return time.Time{}, err
	}
	_, fractionalPart := math.Modf(dateFloat)
	multiply := int64(1_000_000_000)
	for range 9 {
		fractionalPart *= 10
		multiply /= 10
		if nanosecondsDate-int(fractionalPart) < 1 {
			break
		}
	}
	resTime := time.Unix(int64(secondsDate), int64(nanosecondsDate)*multiply)
	return resTime, nil
}
