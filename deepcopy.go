package deepcopy

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type Copier interface {
	// src must be pointer type
	Copy(src interface{}, ptr unsafe.Pointer)
}

var copiers = map[reflect2.Type]Copier{}

func CopierOf(typ reflect2.Type) Copier {
	if c, ok := copiers[typ]; ok {
		return c
	}

	c := createCopierOf(typ)
	println("create copier of", typ.String())
	copiers[typ] = c
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
	return nil
}

func createCopierOfSlice(typ reflect2.Type) Copier {
	if typ.Kind() != reflect.Slice {
		return nil
	}
	sliceType := typ.(*reflect2.UnsafeSliceType)
	elemCopier := CopierOf(sliceType.Elem())
	if elemCopier == nil {
		return nil
	}
	return &sliceCopier{sliceType, elemCopier}
}

func Deepcopy(src interface{}, dst interface{}) error {
	typ := reflect2.TypeOf(dst)
	ptrType := typ.(*reflect2.UnsafePtrType)

	copier := CopierOf(ptrType.Elem())
	ptr := reflect2.PtrOf(dst)
	copier.Copy(src, ptr)
	return nil
}
