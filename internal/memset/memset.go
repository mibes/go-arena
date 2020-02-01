package memset

import "unsafe"

func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)

func Clear(ptr unsafe.Pointer, n uintptr) {
	memclrNoHeapPointers(ptr, n)
}
