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

			got := SelectDays(testCase.want)

			if len(got) != testCase.want {
				t.Errorf("Wanted %d length, got %d", testCase.want, len(got))
			}

		})
	}

	t.Run("numbers should be month days", func(t *testing.T) {
		t.Helper()

		got := SelectDays(100)

		for i := 0; i < 100; i++ {
			if got[i] > 31 || got[i] < 1 {
				t.Errorf("Wanted number between 1-31, got %d (array: %v)", got[i], got)
			}
		}
	})

	t.Run("Numbers should be random", func(t *testing.T) {
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
