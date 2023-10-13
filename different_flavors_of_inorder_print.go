package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IterativeInOrderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	current := root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)
		current = current.Right
	}
	return result
}

func InOrderTraversalRecursive(root *TreeNode) []int {
	var result []int
	internalInOrderTraversalRecursive(root, &result)
	return result
}

func internalInOrderTraversalRecursive(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	internalInOrderTraversalRecursive(root.Left, result)
	*result = append(*result, root.Val)
	internalInOrderTraversalRecursive(root.Right, result)
}

func InOrderTraversalUsingClosure(root *TreeNode) []int {
	var result []int
	var internal func(root *TreeNode)
	internal = func(root *TreeNode) {
		if root == nil {
			return
		}
		internal(root.Left)
		result = append(result, root.Val)
		internal(root.Right)
	}
	internal(root)
	return result
}

func InOrderTraversalUsingCallback(root *TreeNode, callback func(root *TreeNode)) {
	if root == nil {
		return
	}
	InOrderTraversalUsingCallback(root.Left, callback)
	callback(root)
	InOrderTraversalUsingCallback(root.Right, callback)
}

func main() {
	node := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var result []int
	InOrderTraversalUsingCallback(&node, func(root *TreeNode) {
		result = append(result, root.Val)
	})
	fmt.Printf("InOrderTraversalUsingCallback: %v\n", result)

	result = InOrderTraversalUsingClosure(&node)
	fmt.Printf("InOrderTraversalUsingClosure: %v\n", result)

	result = IterativeInOrderTraversal(&node)
	fmt.Printf("IterativeInOrderTraversal: %v\n", result)

	result = InOrderTraversalRecursive(&node)
	fmt.Printf("InOrderTraversalRecursive: %v\n", result)
}
