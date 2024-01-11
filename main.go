package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Arbitary number ; no assignment takes place based on this
const MAX_MEETINGS = 999

// Type to hold groups of cardinal meetings concatenated in a string
type CompliantMeetings struct {
	idx               int
	group_of_meetings string
	count             int
}

// Returns the meeting count in  the string
func get_meeting_count(meeting_group string) int {
	glist := split(meeting_group)
	return len(glist)
}

// joins a list of meetings in numeric format to make a string, sorting if needed
func join(groups_list []int, sort_flag bool) string {
	var str_list []string
	for _, v := range groups_list {
		str_list = append(str_list, fmt.Sprintf("%d", v))
	}
	if true == sort_flag {
		sort.Strings(str_list)
	}
	return strings.Join(str_list, " ")

}

// Splits the string into underlying numeric values
func split(groups string) []int {
	subgroupss := strings.Fields(groups)
	var res []int
	for _, v := range subgroupss {
		vi, _ := strconv.Atoi((v))
		res = append(res, vi)
	}
	return res
}

// Find the overlap between two groups of meetings
func common_meetings(groups1 string, groups2 string) CompliantMeetings {
	groups1_list := split(groups1)
	groups2_list := split(groups2)
	pool := make(map[int]bool)
	var common []int
	for _, v1 := range groups1_list {
		for _, v2 := range groups2_list {
			if v1 == v2 {
				if _, exists := pool[v1]; !exists {
					pool[v1] = true
					common = append(common, v1)
				}
			}
		}
	}
	v := CompliantMeetings{}
	v.count = len(common)
	v.group_of_meetings = join(common, true)
	return v
}

// Used to find the intersection of two groups of meetings to determine unique meetings
func intersect_participant(group1 []string, group2 []string) bool {
	var intersect []string
	for _, v1 := range group1 {
		for _, v2 := range group2 {
			if v1 == v2 {
				intersect = append(intersect, v1)
			}
		}
	}
	var res bool
	if len(intersect) > 0 {
		res = true
	} else {
		res = false
	}
	return res
}

// Reduces the meeting group to be examined when meetings are already slotted
func remove_group_from_uniques(uniques *[]CompliantMeetings, group string) {
	for idx, val := range *uniques {
		tlist := split(val.group_of_meetings)
		glist := split(group)
		rlist := make([]int, 0)
		var found bool
		for _, v1 := range tlist {
			found = false
			for _, v2 := range glist {
				if v1 == v2 {
					found = true
					break
				}
			}
			if false == found {
				rlist = append(rlist, v1)
			}
		}
		(*uniques)[idx].count = len(rlist)
		(*uniques)[idx].group_of_meetings = join(rlist, true)
	}
	return
}

// Logic.
// 1. Groups of unique combinations of meetings that do not share participants. The cumulative meeting pool
// is updated for every addition. Thus the meetings added are unique with respect to the group
//
// 2. A 2D array is built for every group in dim1 and the highest common factor of meetings in dim2
// 2.1 Group with the least HCF is selected for a slot as,
// this has the largest number of attendees, and its meetings are then
// excluded from further consideration.
// 4. Repeated from 2 until all meetings are scheduled into different slots.

func main() {
	meetings := [][]string{{"A", "E"}, {"B", "F"}, {"C", "G"}, {"D", "H"}, {"B", "C", "D"}, {"A", "C", "D"}, {"A", "B", "D"}, {"A", "B", "C"}}
	//meetings := [][]string{{"A"}, {"B"}, {"C"}, {"B", "D"}, {"Z"}}
	//meetings := [][]string{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"}}
	//meetings := [][]string{{"A", "B", "C", "D"}, {"B", "C"}, {"C", "D"}, {"B", "D"}}
	//meetings := [][]string{{"A"}, {"B"}, {"C"}, {"D"}, {"Z"}}
	var uniques []CompliantMeetings
	fmt.Println("Num of meetings=", len(meetings))
	// GM_i=Mi intersect M_j =0
	// Unique meetings with cumulative participants with respect to the current
	// meeting and stored in a group called uniques
	for idx1, _ := range meetings {
		var groups []int
		var res bool
		var pool_meetings []string

		groups_count := 1
		groups = append(groups, idx1)
		pool_meetings = append(pool_meetings, meetings[idx1]...)

		for idx2, val2 := range meetings {
			if idx1 == idx2 {
				continue
			}
			res = intersect_participant(pool_meetings, meetings[idx2])
			if true == res {
				continue
			}
			res = intersect_participant(meetings[idx1], val2)
			if true == res {
				continue
			}
			groups = append(groups, idx2)
			groups_count += 1
			pool_meetings = append(pool_meetings, meetings[idx2]...)
		}

		uniques = append(uniques, CompliantMeetings{idx1, join(groups, true), groups_count})
	}

	// Make square matrix for Common meetings with respect to the current meeting.
	dependency_matrix := make([][]CompliantMeetings, len(uniques))
	for i := 0; i < len(meetings); i++ {
		dependency_matrix[i] = make([]CompliantMeetings, len(uniques))
	}

	slots := make([][]string, len(meetings))
	slot_idx := 0
	meeting_score := make([]int, len(dependency_matrix))
	tmeeting_score := make([]int, len(dependency_matrix))
	total_slotted_meetings := 0
	record_flag := true
	// Loop till all slots are found.
	for {
		// For every meeting group find the common meetings.
		// Every iterations will select a slot and reduce the meeting group.
		// Only respect the selection scores on initial run, as susequently
		// the scores are affected by the removal of meetings.
		// CM_i=GM_i \union GM_j
		for i := 0; i < len(dependency_matrix); i++ {
			count := 0
			for j := 0; j < len(dependency_matrix); j++ {
				common_groups := common_meetings(uniques[i].group_of_meetings, uniques[j].group_of_meetings)
				dependency_matrix[i][j] = common_groups
				count += common_groups.count
			}
			if count == 0 {
				count = MAX_MEETINGS * 10
			}
			tmeeting_score[i] = count
		}
		if true == record_flag {
			copy(meeting_score, tmeeting_score)
			record_flag = false
		}
		// Select the meeting with the smallest overlap
		// and then remove it from selection.
		min_score := MAX_MEETINGS + 1
		min_meeting := -1
		for i := 0; i < len(meeting_score); i++ {
			if meeting_score[i] < min_score {
				min_score = meeting_score[i]
				min_meeting = i
				meeting_score[i] = MAX_MEETINGS * 10
			}
		}
		// Fill in the slot.
		slots[slot_idx] = append(slots[slot_idx], uniques[min_meeting].group_of_meetings)
		slot_idx += 1
		total_slotted_meetings += get_meeting_count(uniques[min_meeting].group_of_meetings)
		// Termination condition.
		if total_slotted_meetings == len(uniques) {
			break
		}
		// Remove the meetings in the slot from the next group to be examined
		remove_group_from_uniques(&uniques, uniques[min_meeting].group_of_meetings)
	}
	// Print allocated slots.
	for i := 0; i < slot_idx; i++ {
		fmt.Printf("Slot[%d]=%s\n", i, slots[i])
	}
}
