package deepcopy

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type mapCopier struct {
	mapType               *reflect2.UnsafeMapType
	keyType, elemType     reflect2.Type
	keyCopier, elemCopier Copier
}

func (m *mapCopier) Copy(src, dst unsafe.Pointer) {

	l := maplen(*(*unsafe.Pointer)(src))
	m.mapType.UnsafeSet(dst, m.mapType.UnsafeMakeMap(l))
	iter := m.mapType.UnsafeIterate(src).(*reflect2.UnsafeMapIterator)
	for iter.HasNext() {
		sk, sv := iter.UnsafeNext()
		dk := m.keyType.UnsafeNew()
		dv := m.elemType.UnsafeNew()
		m.keyCopier.Copy(sk, dk)
		m.elemCopier.Copy(sv, dv)
		m.mapType.UnsafeSetIndex(dst, dk, dv)
	}
}

func createCopierOfMap(typ reflect2.Type) Copier {
	if typ.Kind() != reflect.Map {
		return nil
	}
	mapType := typ.(*reflect2.UnsafeMapType)
	keyDecoder := CopierOf(mapType.Key())
	elemDecoder := CopierOf(mapType.Elem())
	return &mapCopier{
		mapType:    mapType,
		keyType:    mapType.Key(),
		elemType:   mapType.Elem(),
		keyCopier:  keyDecoder,
		elemCopier: elemDecoder,
	}
}

//go:noescape
//go:linkname maplen reflect.maplen
func maplen(m unsafe.Pointer) int
