package main

import (
	"fmt"
	"sort"
)

type Activity struct {
	start  int
	finish int
}

func activitySelector(activities []Activity) []Activity {
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].finish < activities[j].finish
	})

	selected := []Activity{activities[0]}
	j := 0

	for i := 1; i < len(activities); i++ {
		if activities[i].start >= activities[j].finish {
			selected = append(selected, activities[i])
			j = i
		}
	}

	return selected
}

func main() {
	activities := []Activity{
		{1, 3}, {2, 8}, {4, 9}, {7, 10}, {9, 11},
		{11, 15}, {13, 14}, {15, 18}, {8, 11}, {6, 8},
		{5, 12}, {10, 13}, {2, 4},
	}

	selectedActivities := activitySelector(activities)

	fmt.Println("Вибрані заявки:")
	for _, activity := range selectedActivities {
		fmt.Printf("(%d, %d) ", activity.start, activity.finish)
	}
	fmt.Printf("\nМаксимальна кількість заявок, що не перетинаються: %d", len(selectedActivities))
}
