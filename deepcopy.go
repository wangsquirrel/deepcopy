package deepcopy

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// DeepCopy copies the data pointed by src to dst
// both src and dst myst be pointers
func DeepCopy(src interface{}, dst interface{}) error {
	dstType := reflect2.TypeOf(dst)
	srcType := reflect2.TypeOf(src)
	if srcType != dstType {
		return fmt.Errorf("can not copy objects of different types(%v, %v)", srcType, dstType)
	}
	ptrType := dstType.(*reflect2.UnsafePtrType)
	copier := CopierOf(ptrType.Elem())
	if copier == nil {
		return fmt.Errorf("object of type %s can not be copied", dstType)
	}
	copier.Copy(reflect2.PtrOf(src), reflect2.PtrOf(dst))
	return nil
}

type Copier interface {
	// Copy copies the data pointed by src to dst
	// src and dst must be pointer type and they points to the objects of the same type
	Copy(src unsafe.Pointer, dst unsafe.Pointer)
}

var copiers = sync.Map{}

func CopierOf(typ reflect2.Type) Copier {
	if c, ok := copiers.Load(typ); ok {
		return c.(Copier)
	}
	c := createCopierOf(typ)
	//println("create copier of", typ.String())
	copiers.Store(typ, c)
	return c

}

func createCopierOf(typ reflect2.Type) Copier {

	c := createCopierOfNative(typ)
	if c != nil {
		return c
	}
	c = createCopierOfSlice(typ)
	if c != nil {
		return c
	}
	c = createCopierOfStruct(typ)
	if c != nil {
		return c
	}
	c = createCopierOfMap(typ)
	if c != nil {
		return c
	}
	c = createCopierOfPtr(typ)
	if c != nil {
		return c
	}
	c = createCopierOfEface(typ)
	if c != nil {
		return c
	}
	c = createCopierOfiface(typ)
	if c != nil {
		return c
	}
	c = createCopierOfArray(typ)
	if c != nil {
		return c
	}
	return nil
}
