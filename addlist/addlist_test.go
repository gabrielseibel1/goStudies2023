package addlist

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		// Input: l1 = [2,4,3], l2 = [5,6,4]
		// Output: [7,0,8]
		// Explanation: 342 + 465 = 807.
		// Example 2:

		// Input: l1 = [0], l2 = [0]
		// Output: [0]
		// Example 3:

		// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
		// Output: [8,9,9,9,0,0,0,1]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
