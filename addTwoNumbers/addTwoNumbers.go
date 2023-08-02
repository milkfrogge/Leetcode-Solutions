package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers TODO: can be optimized
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	left, right, ref := 0, 0, 0
	s := make([]ListNode, 0)
	for !(l1 == nil && l2 == nil) {

		switch l1 {
		case nil:
			left = 0
		default:
			left = l1.Val
			l1 = l1.Next
		}
		switch l2 {
		case nil:
			right = 0
		default:
			right = l2.Val
			l2 = l2.Next
		}

		var temp ListNode
		temp.Val = (left + right + ref) % 10
		s = append(s, temp)
		ref = (left + right + ref) / 10

	}

	if ref != 0 {
		var temp ListNode
		temp.Val = ref
		s = append(s, temp)
	}

	for i := len(s) - 2; i >= 0; i-- {
		s[i].Next = &s[i+1]
	}
	return &s[0]
}
