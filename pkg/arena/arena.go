package arena

import (
	"errors"
	"unsafe"
)

const (
	initialCapacity = 1
)

type Arena struct {
	dataSize int
	buffers  []*Buffer
	buffer   *Buffer
}

type Buffer struct {
	p        unsafe.Pointer
	pos      int
	length   int
	dataSize int
	buffer   []byte
}

func newBuffer(dataSize, capacity int) *Buffer {
	bufSize := dataSize * capacity
	buffer := make([]byte, bufSize, bufSize)
	p := unsafe.Pointer(&buffer[0])

	return &Buffer{
		buffer:   buffer,
		pos:      0,
		length:   capacity,
		dataSize: dataSize,
		p:        p,
	}
}

func NewArena(dataType interface{}) *Arena {
	dataSize := int(unsafe.Sizeof(dataType))
	buffer := newBuffer(dataSize, initialCapacity)

	return &Arena{
		dataSize: dataSize,
		buffers:  []*Buffer{buffer},
		buffer:   buffer,
	}
}

func (b *Buffer) move() error {
	b.pos++
	if b.pos >= b.length {
		return errors.New("out of memory")
	}

	b.p = unsafe.Pointer(&b.buffer[b.pos*b.dataSize])
	return nil
}

func (a *Arena) reAlloc(size int) {
	buffer := newBuffer(a.dataSize, size)
	a.buffers = append(a.buffers, buffer)
	a.buffer = buffer
}

func (a *Arena) Alloc() unsafe.Pointer {
	if err := a.buffer.move(); err != nil {
		a.reAlloc(a.buffer.length * 2)
	}

	return a.buffer.p
}

func (a *Arena) Release() {
	for idx := range a.buffers {
		a.buffers[idx] = nil
	}
}