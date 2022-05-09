package deepcopy

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestFuncCopy(t *testing.T) {
	a := func(x int) int { return x + 1 }
	var b func(x int) int
	DeepCopy(&a, &b)
	assert.Equal(t, a(1), b(1))
}

func TestUnsafePointerCopy(t *testing.T) {
	i := 1
	a := unsafe.Pointer(&i)
	var b unsafe.Pointer

	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	a = unsafe.Pointer(nil)
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)
}

func TestChanCopy(t *testing.T) {
	var a chan int = make(chan int, 9)
	var b chan int = make(chan int, 98)
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)
}

func TestUintptrCopy(t *testing.T) {
	x := 1
	a := uintptr(unsafe.Pointer(&x))
	var b uintptr
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)
}
