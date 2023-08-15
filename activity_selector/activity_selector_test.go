package activityselector

import (
	"fmt"
	"reflect"
	"testing"
)

func TestActivitySelector(t *testing.T) {
	var ActivityPool = []string{"foo", "bar"}

	t.Run("activities should be selected",func(t *testing.T) {
		got := SelectActivities(1, ActivityPool)
	
		if len(got) != 1 {
			t.Errorf("Expected to select one string, got %s", got)
		}
	
		got = SelectActivities(2, ActivityPool)
	
		if len(got) != 2 {
			t.Errorf("Expected to select two strings, got %s", got)
		}
	})

	t.Run("activities should be from list", func(t *testing.T) {
		t.Helper()

		got := SelectActivities(2, ActivityPool)

		assertValueInRange(got, ActivityPool, t)

		ActivityPool = []string{"bar"}
		got = SelectActivities(2, ActivityPool)

		assertValueInRange(got, ActivityPool, t)
	})

	t.Run("activities should be picked at random", func(t *testing.T) {
		spyRandom := SpyRand{}

		got := selectActivities(2, &spyRandom)
		want := []string{"1", "2"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}

		if spyRandom.count == 0 {
			t.Errorf("random not called")
		}
	})
}

func assertValueInRange(got []string, list []string, t *testing.T) {
	for _, activityResult := range got {
		inRange := false
		for _, activity := range list {
			if activity == activityResult {
				inRange = true
				break
			}
		}

		if !inRange {
			t.Errorf("Expected %s to be in range %q", activityResult, list)
		}
	}
}

type SpyRand struct {
	count int
}

func (s *SpyRand) randomActivity() string {
	s.count++
	return fmt.Sprint(s.count)
}
