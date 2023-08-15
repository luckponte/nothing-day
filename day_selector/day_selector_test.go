package day_selector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSelectDay(t *testing.T) {
	testCases := []struct {
		want int
	}{
		{1},
		{2},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("should select %d days", rune(testCase.want)), func(t *testing.T) {
			t.Helper()

			got := SelectDays(testCase.want, 1)

			if len(got) != testCase.want {
				t.Errorf("Wanted %d length, got %d", testCase.want, len(got))
			}

		})
	}

	t.Run("numbers should be month days", func(t *testing.T) {
		t.Helper()

		testCases := []struct {
			month    int
			dayLimit int
		}{
			{1, 31},
			{2, 28},
			{3, 31},
			{4, 30},
			{5, 31},
			{6, 30},
			{7, 31},
			{8, 31},
			{9, 30},
			{10, 31},
			{11, 30},
			{12, 31},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("Month %d", testCase.month), func(t *testing.T) {
				t.Helper()
				got := SelectDays(1000, testCase.month)
				reachedLimit := false

				for i := 0; i < 1000; i++ {
					if got[i] > testCase.dayLimit || got[i] < 1 {
						t.Errorf("Wanted number between 1-%d, got %d (array: %v)", testCase.dayLimit, got[i], got)
					}

					if got[i] == testCase.dayLimit {
						reachedLimit = true
					}
				}

				if !reachedLimit {
					t.Errorf("Expected to reach day %d at least once", testCase.dayLimit)
				}
			})
		}
	})

	t.Run("numbers should be random", func(t *testing.T) {
		t.Helper()

		spyRandom := SpyRand{}
		got := selectDays(2, &spyRandom)
		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}

		if spyRandom.count == 0 {
			t.Errorf("random not called")
		}
	})
}

type SpyRand struct {
	count int
}

func (s *SpyRand) randomDay() int {
	s.count++
	return s.count
}
