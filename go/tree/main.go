package main

import (
	"fmt"
)

// Node represents a node in the tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BFS performs Breadth-First Search on a tree starting from the root node
func BFS(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	fmt.Println("len : ", len(queue))

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	fmt.Println()
}

// DFS performs Depth-First Search on a tree starting from the root node
func DFS(root *Node) {
	if root == nil {
		return
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", current.Value)

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}
	fmt.Println()
}

func main() {
	// Example tree:
	//       1
	//     /   \
	//    2     3
	//   / \   / \
	//  4   5 6   7

	root := &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
			Left: &Node{
				Value: 4,
			},
			Right: &Node{
				Value: 5,
			},
		},
		Right: &Node{
			Value: 3,
			Left: &Node{
				Value: 6,
			},
			Right: &Node{
				Value: 7,
			},
		},
	}

	fmt.Println("BFS traversal:")
	BFS(root)

	fmt.Println("DFS traversal:")
	DFS(root)
}
