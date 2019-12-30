package main

import (
	"fmt"

	"github.com/biuboombiuboom/leetcode/nornal/algorithm"
)

func main() {

	l1 := biuldListNode([]int{7, 2, 4, 3})
	l2 := biuldListNode([]int{5, 6, 4})

	sum := algorithm.AddTwoNumbers(l1, l2)
	printListNode(sum)

}

func printListNode(head *algorithm.ListNode) {
	for head != nil {
		fmt.Printf("%d,", head.Val)
		head = head.Next
	}
}

func biuldListNode(nums []int) *algorithm.ListNode {
	begin := &algorithm.ListNode{}
	pre := begin
	for i := 0; i < len(nums); i++ {
		pre.Next = &algorithm.ListNode{
			Val: nums[i],
		}
		pre = pre.Next
	}
	return begin.Next
}

func buildCycleListNode(nums []int, pos int) *algorithm.ListNode {
	begin := &algorithm.ListNode{}
	pre := begin
	var cycleBegin *algorithm.ListNode
	for i := 0; i < len(nums); i++ {
		pre.Next = &algorithm.ListNode{
			Val: nums[i],
		}
		pre = pre.Next
		if pos == i {
			cycleBegin = pre
		}
	}
	pre.Next = cycleBegin
	return begin.Next
}
