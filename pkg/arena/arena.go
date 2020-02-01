package arena

import (
	"errors"
	"math"
	"unsafe"
)

const (
	initialCapacity      = 1
	scaleFactor          = 2
	maxCapacityIncrement = 1024 * 1024 * 128 // 128 Megabyte
)

type Arena struct {
	dataSize            int
	maxCapacityIncrease int
	totalBytes          int
	buffers             []*Buffer
	buffer              *Buffer
}

type Buffer struct {
	p        unsafe.Pointer
	pos      int
	capacity int
	dataSize int
	buffer   []byte
}

func newBuffer(dataSize, capacity, maxCapacity int) (*Buffer, int) {
	newCapacity := int(math.Min(float64(capacity), float64(maxCapacity)))
	bufSize := dataSize * newCapacity

	buffer := make([]byte, bufSize, bufSize)
	p := unsafe.Pointer(&buffer[0])

	return &Buffer{
		buffer:   buffer,
		pos:      0,
		capacity: newCapacity,
		dataSize: dataSize,
		p:        p,
	}, bufSize
}

func NewArena(dataType interface{}) *Arena {
	dataSize := int(unsafe.Sizeof(dataType))
	maxCapacity := int(math.Floor(maxCapacityIncrement / float64(dataSize)))
	buffer, totalBytes := newBuffer(dataSize, initialCapacity, maxCapacity)

	return &Arena{
		dataSize:            dataSize,
		buffers:             []*Buffer{buffer},
		buffer:              buffer,
		totalBytes:          totalBytes,
		maxCapacityIncrease: maxCapacity,
	}
}

func (b *Buffer) move() error {
	b.pos++
	if b.pos >= b.capacity {
		return errors.New("out of memory")
	}

	b.p = unsafe.Pointer(&b.buffer[b.pos*b.dataSize])
	return nil
}

func (b *Buffer) Release() {
	b.buffer = nil
	b.capacity = 0
	b.pos = 0
}

func (a *Arena) reAlloc(size int) {
	var incrBytes int
	a.buffer, incrBytes = newBuffer(a.dataSize, size, a.maxCapacityIncrease)
	a.totalBytes += incrBytes
	a.buffers = append(a.buffers, a.buffer)
}

func (a *Arena) Alloc() unsafe.Pointer {
	if err := a.buffer.move(); err != nil {
		a.reAlloc(a.buffer.capacity * scaleFactor)
	}

	return a.buffer.p
}

func (a *Arena) Release() {
	for idx := range a.buffers {
		a.buffers[idx].Release()
	}

	a.buffers = nil
	a.buffer = nil
}
