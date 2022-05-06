package deepcopy

import (
	"unsafe"

	"github.com/modern-go/reflect2"
)

type sliceCopier struct {
	typ        *reflect2.UnsafeSliceType
	elemCopier Copier
}

func (s *sliceCopier) Copy(src interface{}, ptr unsafe.Pointer) {
	length := s.typ.LengthOf(src)
	s.typ.UnsafeGrow(ptr, length)
	//s.typ.UnsafeSet(ptr, s.typ.UnsafeMakeSlice(s.typ.LengthOf(src), s.typ.Cap(src)))

	for i := 0; i < length; i++ {
		elemPtr := s.typ.UnsafeGetIndex(ptr, i)
		s.elemCopier.Copy(s.typ.GetIndex(src, i), elemPtr)
	}

}
