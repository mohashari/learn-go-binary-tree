package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummy(t *testing.T) {
	nd := NewNode(0, nil, nil)
	assert.NotNil(t, nd)
}

func TestIsBinarySearchTree_Nil(t *testing.T) {
	var node *Node
	assert.True(t, IsBinarySearchTree(node))
}

func TestIsBinarySearchTree_OneNode(t *testing.T) {
	node := NewNode(5, nil, nil)
	assert.True(t, IsBinarySearchTree(node))
}
