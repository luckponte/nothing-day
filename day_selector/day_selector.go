package day_selector

import "math/rand"

type IRandomDaySelector interface {
	randomDay() int
}

type randomDaySelector struct{}

func (r *randomDaySelector) randomDay() int {
	return rand.Intn(30) + 1
}

func selectDays(numDays int, rd IRandomDaySelector) []int {
	selectedDays := make([]int, numDays)

	for i := 0; i < numDays; i++ {
		selectedDays[i] = rd.randomDay()
	}
	return selectedDays
}

func SelectDays(numDays int) []int {
	return selectDays(numDays, &randomDaySelector{})
}