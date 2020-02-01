# go-arena
Arena memory allocator for Go

Experimental pooled memory allocator for Go.

It can be used like this:

```go 

type Tree struct {
	Left  *Tree
	Right *Tree
}

func main() {
    arena := NewArena(Tree{})
	tree := (*Tree)(arena.Alloc())
	arena.Release()
}
```
