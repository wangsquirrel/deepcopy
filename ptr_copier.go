package deepcopy

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type ptrCopier struct {
	typ        *reflect2.UnsafePtrType
	elemType   reflect2.Type
	elemCopier Copier
}

func (p *ptrCopier) Copy(src, dst unsafe.Pointer) {
	if p.typ.UnsafeIsNil(dst) {
		*((*unsafe.Pointer)(dst)) = p.elemType.UnsafeNew()
	}
	p.elemCopier.Copy(*((*unsafe.Pointer)(src)), *((*unsafe.Pointer)(dst)))
}

func createCopierOfPtr(typ reflect2.Type) Copier {
	if typ.Kind() != reflect.Ptr {
		return nil
	}
	ptrType := typ.(*reflect2.UnsafePtrType)
	return &ptrCopier{
		ptrType,
		ptrType.Elem(),
		CopierOf(ptrType.Elem()),
	}
}
