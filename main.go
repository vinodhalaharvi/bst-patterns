package main

import (
	"math"
)

func inOrder(root *TreeNode) []int {
	result := []int{}
	var internal func(root *TreeNode)
	internal = func(root *TreeNode) {
		if root == nil {
			return
		}
		internal(root.Left)
		result = append(result, root.Val)
		internal(root.Right)
	}
	return result
}

func InOrderTraversal(root *TreeNode, callback func(*TreeNode)) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left, callback)
	callback(root)
	InOrderTraversal(root.Right, callback)
}

func PreOrderTraversal(root *TreeNode, callback func(*TreeNode)) {
	if root == nil {
		return
	}
	callback(root)
	InOrderTraversal(root.Left, callback)
	InOrderTraversal(root.Right, callback)
}

func PostOrderTraversal(root *TreeNode, callback func(*TreeNode)) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left, callback)
	InOrderTraversal(root.Right, callback)
	callback(root)
}

func SearchBST(root *TreeNode, value int, callback func(*TreeNode)) *TreeNode {
	if root == nil {
		return root // don't call back if the root is nil so the caller doesn't have to check for nil
	}
	if root.Val == value {
		callback(root)
	}
	if value < root.Val {
		return SearchBST(root.Left, value, callback)
	}
	return SearchBST(root.Right, value, callback)
}

func InsertBST(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: value}
	}
	if value < root.Val {
		root.Left = InsertBST(root.Left, value)
	}
	root.Right = InsertBST(root.Right, value)
	return root
}

func DeleteBST(root *TreeNode, value int, callback func(*TreeNode)) *TreeNode {
	if root == nil || root.Left == nil && root.Right == nil {
		return root
	}
	if value < root.Val {
		root.Left = DeleteBST(root.Left, value, callback)
	}
	if value > root.Val {
		root.Right = DeleteBST(root.Right, value, callback)
	}
	if value == root.Val {
		callback(root)
	}
	return root
}

func closestValue(root *TreeNode, target int) int {
	closestDiff := math.MaxInt
	closestNode := &TreeNode{Val: -1}
	internalClosestValue(root, target, &closestDiff, closestNode)
	return closestNode.Val
}

func internalClosestValue(
	root *TreeNode,
	target int,
	closestDiff *int,
	closestNode *TreeNode,
) {
	if root == nil {
		return
	}
	if *closestDiff > min(*closestDiff, abs(root.Val-target)) {
		*closestDiff = min(*closestDiff, abs(root.Val-target))
		closestNode.Val = root.Val
		return
	}
	internalClosestValue(root.Left, target, closestDiff, closestNode)
	internalClosestValue(root.Right, target, closestDiff, closestNode)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i

}

func FindMin(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	return FindMin(root.Left)
}

func FindMax(root *TreeNode) *TreeNode {
	if root == nil || root.Right == nil {
		return root
	}
	return FindMax(root.Right)
}

func countNodesEqualToSumOfDescendants(root *TreeNode) int {
	sumValue, count := 0, 0
	internalCountNodesEqualToSumOfDescendants(root, &sumValue, &count)
	return count
}

func internalCountNodesEqualToSumOfDescendants(root *TreeNode, sumValue *int, count *int) {
	if root == nil {
		return
	}
	internalCountNodesEqualToSumOfDescendants(root.Left, sumValue, count)
	internalCountNodesEqualToSumOfDescendants(root.Right, sumValue, count)
	if root.Val == *sumValue {
		*count++
	}
	*sumValue += root.Val
}

func nodeEqualToSumOfDescendants(root *TreeNode) *TreeNode {
	sumValue := 0
	return internalNodeEqualToSumOfDescendants(root, &sumValue)
}

func internalNodeEqualToSumOfDescendants(root *TreeNode, sumValue *int) *TreeNode {
	if root == nil {
		return nil
	}
	internalNodeEqualToSumOfDescendants(root.Left, sumValue)
	internalNodeEqualToSumOfDescendants(root.Right, sumValue)
	if *sumValue == root.Val {
		return root
	}
	*sumValue += root.Val
	return nil
}

//func main() {
//	root := &TreeNode{
//		Val: 17,
//		Left: &TreeNode{
//			Val: 7,
//			Left: &TreeNode{
//				Val: 5,
//			},
//			Right: &TreeNode{
//				Val: 2,
//			},
//		},
//
//		Right: &TreeNode{
//			Val: 3,
//		},
//	}
//	//descendants := nodeEqualToSumOfDescendants(root)
//	//if descendants != nil {
//	//	println(descendants.Val)
//	//} else {
//	//	println("No node found")
//	//}
//	value := closestValue(root, 7)
//	println(value)
//	//SearchBST(root, 7, func(node *TreeNode) {
//	//	fmt.Printf("Found node: %v\n", node.Val)
//	//})
//}
