# go-arena
Arena memory allocator for Go

Experimental pooled memory allocator for Go.

It can be used like this:

```go 
package main

import (
    "github.com/mibes/go-arena/pkg/arena"
)

type Tree struct {
    Left  *Tree
    Right *Tree
}

func main() {
    a := arena.NewArena(Tree{})
    tree := (*Tree)(a.Alloc())
    left := (*Tree)(a.Alloc())
    right := (*Tree)(a.Alloc())

    tree.Right, tree.Left = right, left
    a.Release()
}

```
