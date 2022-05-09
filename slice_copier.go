package deepcopy

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type sliceCopier struct {
	typ        *reflect2.UnsafeSliceType
	elemCopier Copier
}

func (s *sliceCopier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	length := s.typ.UnsafeLengthOf(src)
	s.typ.UnsafeGrow(dst, length)
	//s.typ.UnsafeSet(ptr, s.typ.UnsafeMakeSlice(s.typ.LengthOf(src), s.typ.Cap(src)))

	for i := 0; i < length; i++ {
		elemPtr := s.typ.UnsafeGetIndex(dst, i)
		s.elemCopier.Copy(s.typ.UnsafeGetIndex(src, i), elemPtr)
	}
}

func createCopierOfSlice(typ reflect2.Type) Copier {
	if typ.Kind() != reflect.Slice {
		return nil
	}
	sliceType := typ.(*reflect2.UnsafeSliceType)
	elemCopier := CopierOf(sliceType.Elem())
	return &sliceCopier{sliceType, elemCopier}
}

type arrayCopier struct {
	typ        *reflect2.UnsafeArrayType
	elemCopier Copier
	length     int
}

func (s *arrayCopier) Copy(src, dst unsafe.Pointer) {
	for i := 0; i < s.length; i++ {
		elemPtr := s.typ.UnsafeGetIndex(dst, i)
		s.elemCopier.Copy(s.typ.UnsafeGetIndex(src, i), elemPtr)
	}
}

func createCopierOfArray(typ reflect2.Type) Copier {
	if typ.Kind() != reflect.Array {
		return nil
	}
	arrayType := typ.(*reflect2.UnsafeArrayType)

	return &arrayCopier{
		arrayType,
		CopierOf(arrayType.Elem()),
		arrayType.Len(),
	}
}
