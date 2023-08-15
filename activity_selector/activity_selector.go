package activityselector

import "math/rand"

var ActivityPool = []string{"foo", "bar"}

type IRandomActivitySelector interface {
	randomActivity() string
}

type randomActivitySelector struct{}

func (r *randomActivitySelector) randomActivity() string {
	limit := len(ActivityPool)
	index := rand.Intn(limit)
	return ActivityPool[index]
}

func selectActivities(activityNum int, ra IRandomActivitySelector) []string {
	activityList := make([]string, activityNum)
	for i := 0; i < len(activityList); i++ {
		activityList[i] = ra.randomActivity()
	}
	return activityList
}

func SelectActivities(activityNum int) []string {
	return selectActivities(activityNum, &randomActivitySelector{})
}
