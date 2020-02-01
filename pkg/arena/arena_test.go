package arena

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Tree struct {
	Left  *Tree
	Right *Tree
}

func TestEmptyAllocation(t *testing.T) {
	a := assert.New(t)
	arena := NewArena(Tree{})
	tree := (*Tree)(arena.Alloc())
	a.Nil(tree.Right)
	a.Nil(tree.Left)
	arena.Release()
}
