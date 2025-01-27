package binaryTree_test

import (
	"github.com/thaksananan-01/7-solution-test/binaryTree"
	"testing"
)

func TestbinaryTreeWithoutLoop(t *testing.T) {
	triangle1 := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	expected1 := 237
	result1 := binaryTree.TriangleSumValueWithoutLoop(triangle1)
	if result1 != expected1 {
		t.Errorf("Expected %d, but got %d for triangle1", expected1, result1)
	}

}
