package main

import (
	"reflect"
	"testing"
)

// ... [Your traversal functions remain unchanged]

func createTestTree() *TreeNode {
	return &TreeNode{
		Val: 15,
		Left: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 12,
			},
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 17,
			},
			Right: &TreeNode{
				Val: 25,
				Left: &TreeNode{
					Val: 23,
				},
			},
		},
	}
}

func TestTraversals(t *testing.T) {
	root := createTestTree()
	expectedResult := []int{2, 5, 7, 10, 12, 15, 17, 20, 23, 25}

	t.Run("Test IterativeInOrderTraversal", func(t *testing.T) {
		result := IterativeInOrderTraversal(root)
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("expected %v, got %v", expectedResult, result)
		}
	})

	t.Run("Test InOrderTraversalRecursive", func(t *testing.T) {
		result := InOrderTraversalRecursive(root)
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("expected %v, got %v", expectedResult, result)
		}
	})

	t.Run("Test InOrderTraversalUsingClosure", func(t *testing.T) {
		result := InOrderTraversalUsingClosure(root)
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("expected %v, got %v", expectedResult, result)
		}
	})

	t.Run("Test InOrderTraversalUsingCallback", func(t *testing.T) {
		var result []int
		InOrderTraversalUsingCallback(root, func(node *TreeNode) {
			result = append(result, node.Val)
		})
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("expected %v, got %v", expectedResult, result)
		}
	})
}
