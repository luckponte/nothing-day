package activityselector

import "math/rand"

type IRandomActivitySelector interface {
	randomActivity() string
}

type randomActivitySelector struct {
	pool []string
}

func (r *randomActivitySelector) randomActivity() string {
	limit := len(r.pool)
	index := rand.Intn(limit)
	return r.pool[index]
}

func selectActivities(activityNum int, ra IRandomActivitySelector) []string {
	activityList := make([]string, activityNum)
	for i := 0; i < len(activityList); i++ {
		activityList[i] = ra.randomActivity()
	}
	return activityList
}

func SelectActivities(activityNum int, activityPool []string) []string {
	return selectActivities(activityNum, &randomActivitySelector{activityPool})
}
