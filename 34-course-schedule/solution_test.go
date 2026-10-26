package courseschedule

import "testing"

type scheduleCase struct {
	name          string
	numCourses    int
	prerequisites [][]int
	possible      bool
}

var scheduleCases = []scheduleCase{
	{name: "single course no prerequisites", numCourses: 1, prerequisites: nil, possible: true},
	{name: "simple chain", numCourses: 2, prerequisites: [][]int{{1, 0}}, possible: true},
	{name: "two-course cycle", numCourses: 2, prerequisites: [][]int{{1, 0}, {0, 1}}, possible: false},
	{name: "diamond", numCourses: 4, prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, possible: true},
	{name: "self loop", numCourses: 3, prerequisites: [][]int{{1, 0}, {2, 2}}, possible: false},
	{name: "long cycle", numCourses: 4, prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}}, possible: false},
	{name: "independent courses", numCourses: 5, prerequisites: nil, possible: true},
	{name: "disconnected components one with cycle", numCourses: 5, prerequisites: [][]int{{1, 0}, {3, 2}, {4, 3}, {2, 4}}, possible: false},
	{name: "duplicate edges", numCourses: 3, prerequisites: [][]int{{1, 0}, {1, 0}, {2, 1}}, possible: true},
	{name: "many prerequisites for one course", numCourses: 6, prerequisites: [][]int{{5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {4, 0}}, possible: true},
}

func TestCanFinish(t *testing.T) {
	for _, tc := range scheduleCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CanFinish(tc.numCourses, tc.prerequisites)
			if got != tc.possible {
				t.Errorf("CanFinish(%d, %v) = %v, want %v", tc.numCourses, tc.prerequisites, got, tc.possible)
			}
		})
	}
}

func TestFindOrder(t *testing.T) {
	for _, tc := range scheduleCases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindOrder(tc.numCourses, tc.prerequisites)
			if !tc.possible {
				if len(got) != 0 {
					t.Errorf("FindOrder(%d, %v) = %v, want an empty slice because of a cycle", tc.numCourses, tc.prerequisites, got)
				}
				return
			}
			if msg := validateOrder(tc.numCourses, tc.prerequisites, got); msg != "" {
				t.Errorf("FindOrder(%d, %v) = %v is not a valid order: %s", tc.numCourses, tc.prerequisites, got, msg)
			}
		})
	}
}

func validateOrder(numCourses int, prerequisites [][]int, order []int) string {
	if len(order) != numCourses {
		return "expected every course exactly once"
	}
	position := make(map[int]int, numCourses)
	for i, course := range order {
		if course < 0 || course >= numCourses {
			return "contains an out-of-range course"
		}
		if _, seen := position[course]; seen {
			return "contains a duplicate course"
		}
		position[course] = i
	}
	for _, edge := range prerequisites {
		course, prereq := edge[0], edge[1]
		if position[prereq] > position[course] {
			return "a course appears before one of its prerequisites"
		}
	}
	return ""
}
