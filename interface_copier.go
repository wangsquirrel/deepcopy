package deepcopy

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type efaceCopier struct {
}

var efaceCopierPtr = &efaceCopier{}

func (i *efaceCopier) Copy(src, dst unsafe.Pointer) {

	srcEface := *((*interface{})(src))
	dstEfacePtr := (*eface)(dst)
	if srcEface == nil {
		*dstEfacePtr = eface{nil, nil}
		return
	}

	if reflect2.IsNil(srcEface) {
		*dstEfacePtr = eface{(*eface)(src).rtype, nil}
		return
	}

	srcType := reflect2.TypeOf(srcEface)
	if srcType.Kind() != reflect.Ptr {
		*dstEfacePtr = eface{(*eface)(src).rtype, srcType.UnsafeNew()}
		CopierOf(srcType).Copy((*eface)(src).data, (*eface)(dst).data)
		return
	}

	ptrElemType := srcType.(*reflect2.UnsafePtrType).Elem()
	*dstEfacePtr = eface{(*eface)(src).rtype, ptrElemType.UnsafeNew()}

	CopierOf(ptrElemType).Copy((*eface)(src).data, (*eface)(dst).data)
}

func createCopierOfEface(typ reflect2.Type) Copier {
	_, ok := typ.(*reflect2.UnsafeEFaceType)
	if !ok {
		return nil
	}
	return efaceCopierPtr
}

type eface struct {
	rtype unsafe.Pointer
	data  unsafe.Pointer
}

type ifaceCopier struct {
	valType *reflect2.UnsafeIFaceType
}

func (i *ifaceCopier) Copy(src, dst unsafe.Pointer) {
	if i.valType.UnsafeIsNil(src) {
		i.valType.UnsafeSet(dst, i.valType.UnsafeNew())
		return
	}

	srcEface := i.valType.UnsafeIndirect(src)
	srcType := reflect2.TypeOf(srcEface)
	if reflect2.IsNil(srcEface) {
		i.valType.UnsafeSet(dst, unsafe.Pointer(&iface{(*iface)(src).itab, nil}))
		return
	}

	if srcType.Kind() == reflect.Ptr {
		ptrElemType := srcType.(*reflect2.UnsafePtrType).Elem()
		i.valType.UnsafeSet(dst, unsafe.Pointer(&iface{(*iface)(src).itab, ptrElemType.UnsafeNew()}))
		CopierOf(ptrElemType).Copy((*iface)(src).data, (*iface)(dst).data)
		return
	}

	i.valType.UnsafeSet(dst, unsafe.Pointer(&iface{(*iface)(src).itab, srcType.UnsafeNew()}))
	CopierOf(srcType).Copy((*iface)(src).data, (*iface)(dst).data)
}

func createCopierOfiface(typ reflect2.Type) Copier {
	ifaceTyp, ok := typ.(*reflect2.UnsafeIFaceType)
	if !ok {
		return nil
	}
	return &ifaceCopier{
		ifaceTyp,
	}
}

type iface struct {
	itab *itab
	data unsafe.Pointer
}

type itab struct {
	ignore unsafe.Pointer
	rtype  unsafe.Pointer
}
