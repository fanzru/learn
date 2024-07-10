package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Define two example linked lists
	list1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	list2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	// Call the function to add the two numbers represented by the linked lists
	result := addTwoNumbers(list1, list2)

	// Print the result for verification
	printList(result)
}

// longest_substring_without_repeating_characters

// Function to add two numbers represented by linked lists
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int // Carry to handle values >= 10
	var dummy = &ListNode{}
	current := dummy // Pointer to the current node in the resulting linked list

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry // Initialize sum with carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10          // Update carry
		current.Next = &ListNode{ // Append new node with sum % 10 to the result list
			Val:  sum % 10,
			Next: nil,
		}
		current = current.Next // Move current pointer to the next node
	}

	return dummy.Next // Return the head of the resulting linked list
}

// Helper function to print values of a linked list (for verification)
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	var nextNodeIncrementValue int
// 	var head *ListNode
// 	listResult := head

// 	for l1 != nil && l2 != nil {
// 		nodeValue := l1.Val + l2.Val + nextNodeIncrementValue
// 		nextNodeIncrementValue = 0

// 		if nodeValue >= 10 {
// 			nodeValue = nodeValue - 10
// 			nextNodeIncrementValue = 1
// 		}

// 		if listResult == nil {
// 			listResult = &ListNode{
// 				Val:  nodeValue,
// 				Next: nil,
// 			}
// 			listResult = listResult.Next
// 		} else {

// 			listResult = &ListNode{
// 				Val:  nodeValue,
// 				Next: nil,
// 			}
// 			listResult = listResult.Next
// 		}
// 		l1 = l1.Next
// 		l2 = l2.Next
// 	}

// 	h2 := head.Next
// 	h3 := h2.Next
// 	log.Println(head.Val, h2.Val, h3.Val)
// 	return head
// }
