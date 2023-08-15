package day_selector

import "math/rand"

type IRandomDaySelector interface {
	randomDay() int
}

type randomDaySelector struct {
	dayLimit int
}

var monthLimits = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func (r *randomDaySelector) randomDay() int {
	return rand.Intn(r.dayLimit) + 1
}

func selectDays(numDays int, rd IRandomDaySelector) []int {
	selectedDays := make([]int, numDays)

	for i := 0; i < numDays; i++ {
		selectedDays[i] = rd.randomDay()
	}
	return selectedDays
}

func SelectDays(numDays int, month int) []int {
	daySelector := randomDaySelector{monthLimits[month]}
	return selectDays(numDays, &daySelector)
}
