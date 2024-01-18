package main

import (
	"fmt"
	"strings"
	"testing"
)

// TestGetMeetingCount tests the get_meeting_count function.
func TestGetMeetingCount(t *testing.T) {
	meeting_group := "1 2 3 4 5 6 7 8 9 10"
	expected_result := 10
	result := get_meeting_count(meeting_group)
	if expected_result != result {
		t.Errorf("***Expected %d Got %d", expected_result, result)
	}
}

// TestJoin tests the join function.
func TestJoin(t *testing.T) {
	test_group := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected_result := "1 10 2 3 4 5 6 7 8 9"
	result := join(test_group, true)
	if expected_result != result {
		t.Errorf("***Expected %s Got %s", expected_result, result)
	}
}
func CompareIntLists(short_l1 []int, long_l2 []int, size int) bool {
	if -1 == size {
		if len(short_l1) != len(long_l2) {
			return false
		}
	}
	count := 0
	for _, v1 := range short_l1 {
		for _, v2 := range long_l2 {
			if v1 == v2 {
				count += 1
			}
		}
	}
	if count == len(short_l1) {
		return true
	}
	return false
}
func CompareStringLists(short_l1 []string, long_l2 []string, size int) bool {
	if -1 == size {
		if len(short_l1) != len(long_l2) {
			return false
		}
	}
	count := 0
	for _, v1 := range short_l1 {
		for _, v2 := range long_l2 {
			if strings.TrimSpace(v1) == strings.TrimSpace(v2) {
				count += 1
			}
		}
	}
	if count == len(short_l1) {
		return true
	}
	return false
}

// TestSplit tests the split function.
func TestSplit(t *testing.T) {
	test_group := "1 10 2 3 4 5 6 7 8 9"
	expected_result := []int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9}
	result := split(test_group)
	if false == CompareIntLists(expected_result, result, -1) {
		t.Errorf("***Expected %v Got %v", expected_result, result)
	}
}

// TestCommonMeetings tests the common_meetings function.
func TestCommonMeetings(t *testing.T) {
	meetings1 := "1 2 3 4 5 6 7"
	meetings2 := "1 3 5 7 8 9 10"
	expected_result := "1 3 5 7"
	result := common_meetings(meetings1, meetings2)
	if result.count != 4 {
		t.Errorf("***Expected %v Got %v", expected_result, result.group_of_meetings)
	}
	if result.group_of_meetings != expected_result {
		t.Errorf("***Expected %v Got %v", expected_result, result.group_of_meetings)
	}
}

// TestIntersectParticipant tests the intersect_participant function.
func TestIntersectParticipant(t *testing.T) {
	meetings1 := []string{"1", "2", "3", "4", "5", "6", "7"}
	meetings2 := []string{"3", "8", "9", "10", "11", "12", "13"}
	expected_result := true
	result := intersect_participant(meetings1, meetings2)
	if false == result {
		t.Errorf("***Expected %v Got %v", expected_result, result)
	}
	meetings2 = []string{"8", "9", "10", "11", "12", "13"}
	expected_result = false
	result = intersect_participant(meetings1, meetings2)
	if true == result {
		t.Errorf("***Expected %v Got %v", expected_result, result)
	}
}

// TestRemoveGroupFromUniques tests the remove_group_from_uniques function.
func TestRemoveGroupFromUniques(t *testing.T) {
	uniques := []CompliantMeetings{{0, "1 2", 2}, {1, "2 3", 2}}
	group := "2"
	expected_result := []CompliantMeetings{{0, "1", 1}, {1, "3", 1}}
	remove_group_from_uniques(&uniques, group)
	for _, v1 := range uniques {
		for _, v2 := range expected_result {
			if v1.idx == v2.idx {
				if v1.group_of_meetings != v2.group_of_meetings {
					t.Errorf("***Un-expected common string %s", v1.group_of_meetings)
				}
			}
		}
	}
}

// TestBasicFunctionality tests the basic scheduling of non-overlapping meetings.
func TestBasicFunctionality(t *testing.T) {
	meetings := [][]string{{"A"}, {"B"}, {"C"}, {"D"}}
	expected_result := []string{"0 1 2 3"}
	result, result_count := schedule(meetings)
	if len(expected_result) != result_count {
		t.Errorf("***Expected %v Got %v", expected_result, result)
	}
	if false == CompareStringLists(expected_result, result, result_count) {
		t.Errorf("***Expected %v Got %v", expected_result, result)
	}
	fmt.Printf("meetings=%v\n", meetings)
	for i := 0; i < result_count; i++ {
		fmt.Printf("Slot[%d]=%v\n", i, result[i])
	}
}

// TestOverlappingParticipants tests scheduling of meetings with overlapping participants.
func TestOverlappingParticipants(t *testing.T) {
	meetings := [][]string{{"A", "E"}, {"B", "F"}, {"C", "G"}, {"D", "H"}, {"B", "C", "D"}, {"A", "C", "D"}, {"A", "B", "D"}, {"A", "B", "C"}}
	expected_result := []string{"0 4", "1 5", "2 6", "3 7"}
	result, result_count := schedule(meetings)
	if len(expected_result) != result_count {
		t.Errorf("***Expected %v Got %v", expected_result, result)
	}
	if false == CompareStringLists(expected_result, result, result_count) {
		t.Errorf("***Expected %v Got %v", expected_result, result)
	}
	fmt.Printf("meetings=%v\n", meetings)
	for i := 0; i < result_count; i++ {
		fmt.Printf("Slot[%d]=%v\n", i, result[i])
	}
}

// TestLargeNumberOfParticipants tests the algorithm with a large number of participants.
func TestLargeNumberOfParticipants(t *testing.T) {
	meetings := [][]string{
		{"A0", "B1"},                                                                               // Meeting 1: 2 participants
		{"A1", "B0", "C2"},                                                                         // Meeting 2: 3 participants
		{"B2", "C1", "E0", "F1"},                                                                   // Meeting 3: 4 participants
		{"C0", "D1", "F2", "G2", "H1"},                                                             // Meeting 4: 5 participants
		{"D0", "E1", "G0", "H2", "I1", "J0"},                                                       // Meeting 5: 6 participants
		{"E2", "F0", "I0", "J1", "A2", "B1", "C2"},                                                 // Meeting 6: 7 participants
		{"G1", "H0", "I2", "J2", "A0", "B2", "C0", "D2"},                                           // Meeting 7: 8 participants
		{"A1", "B0", "C1", "D0", "E2", "F1", "G2", "H1", "I0"},                                     // Meeting 8: 9 participants
		{"J0", "A2", "B1", "C2", "D1", "E0", "F2", "G0", "H2", "I1"},                               // Meeting 9: 10 participants
		{"J1", "A0", "B2", "C0", "D2", "E1", "F0", "G1", "H0", "I2", "J2"},                         // Meeting 10: 11 participants
		{"A1", "B0", "C1", "D0", "E2", "F1", "G2", "H1", "I0", "J0", "A2", "B1", "C2"},             // Meeting 11: 13 participants
		{"D1", "E0", "F2", "G0", "H2", "I1", "J1", "A0", "B2", "C0", "D2", "E1", "F0", "G1", "H0"}, // Meeting 12: 15 participants
		{"B1", "C2", "D1", "E0", "F0", "G1", "H2", "I2", "J0", "A1"},                               // Meeting 13: 10 participants
		{"C0", "D2", "E1", "F2", "G0", "H0", "I0", "J1", "A2", "B0", "B2"},                         // Meeting 14: 11 participants
		{"D0", "E2", "F1", "G2", "H1", "I1", "J2", "A0", "B1", "C1"},                               // Meeting 15: 10 participants
		{"E0", "F0", "G0", "H0", "I0", "J0", "A1", "B2", "C2", "D1", "E1"},                         // Meeting 16: 11 participants
		{"F1", "G1", "H1", "I1", "J1", "A2", "B0", "C0", "D2", "E2", "F2"},                         // Meeting 17: 11 participants
		{"G2", "H2", "I2", "J2", "A0", "B1", "C1", "D0", "E0", "F0", "G0"},                         // Meeting 18: 11 participants
		{"H1", "I1", "J1", "A1", "B2", "C2", "D1", "E1", "F1", "G1", "H2", "I2", "J2"},             // Meeting 19: 13 participants
		{"I0", "J0", "A0", "B0", "C0", "D0", "E2", "F2", "G2", "H0", "I1", "J1", "A1", "B1", "C1"}, // Meeting 20: 15 participants
	}
	result, result_count := schedule(meetings)
	fmt.Printf("meetings=%v\n", meetings)
	for i := 0; i < result_count; i++ {
		fmt.Printf("Slot[%d]=%v\n", i, result[i])
	}
}

// TestEmptyMeetingsList tests the behavior with an empty list of meetings.
func TestEmptyMeetingsList(t *testing.T) {
	meetings := [][]string{{}, {}}
	result, result_count := schedule(meetings)
	fmt.Printf("meetings=%v\n", meetings)
	for i := 0; i < result_count; i++ {
		fmt.Printf("Slot[%d]=%v\n", i, result[i])
	}
}

// TestAllParticipantsOverlapping tests when all meetings have overlapping participants.
func TestAllParticipantsOverlapping(t *testing.T) {
	meetings := [][]string{{"A", "B", "D"}, {"B", "C"}, {"C", "D"}, {"B", "D", "E"}}
	result, result_count := schedule(meetings)
	fmt.Printf("meetings=%v\n", meetings)
	for i := 0; i < result_count; i++ {
		fmt.Printf("Slot[%d]=%v\n", i, result[i])
	}
}

// TestSingleMeeting tests the scheduling of a single meeting.
func TestSingleMeeting(t *testing.T) {
	meetings := [][]string{{"A", "B"}}
	result, result_count := schedule(meetings)
	fmt.Printf("meetings=%v\n", meetings)
	for i := 0; i < result_count; i++ {
		fmt.Printf("Slot[%d]=%v\n", i, result[i])
	}
}

// TestDuplicateMeetings tests the scheduling of duplicate meetings.
func TestDuplicateMeetings(t *testing.T) {
	meetings := [][]string{{"A", "B"}, {"A", "B"}}
	result, result_count := schedule(meetings)
	fmt.Printf("meetings=%v\n", meetings)
	for i := 0; i < result_count; i++ {
		fmt.Printf("Slot[%d]=%v\n", i, result[i])
	}
}
