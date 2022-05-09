package deepcopy

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type boolCopier struct{}

func (*boolCopier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*bool)(dst)) = *(*bool)(src)
}

type intCopier struct{}

func (*intCopier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*int)(dst)) = *(*int)(src)
}

type int8Copier struct{}

func (*int8Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*int8)(dst)) = *(*int8)(src)
}

type int16Copier struct{}

func (*int16Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*int16)(dst)) = *(*int16)(src)
}

type int32Copier struct{}

func (*int32Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*int32)(dst)) = *(*int32)(src)
}

type int64Copier struct{}

func (*int64Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*int64)(dst)) = *(*int64)(src)
}

type uintCopier struct{}

func (*uintCopier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*uint)(dst)) = *(*uint)(src)
}

type uint8Copier struct{}

func (*uint8Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*uint8)(dst)) = *(*uint8)(src)
}

type uint16Copier struct{}

func (*uint16Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*uint16)(dst)) = *(*uint16)(src)

}

type uint32Copier struct{}

func (*uint32Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*uint32)(dst)) = *(*uint32)(src)
}

type uint64Copier struct{}

func (*uint64Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*uint64)(dst)) = *(*uint64)(src)
}

type uintptrCopier struct{}

func (*uintptrCopier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*uintptr)(dst)) = *(*uintptr)(src)
}

type float32Copier struct{}

func (*float32Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*float32)(dst)) = *(*float32)(src)
}

type float64Copier struct{}

func (*float64Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*float64)(dst)) = *(*float64)(src)
}

type complex64Copier struct{}

func (complex64Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*complex64)(dst)) = *(*complex64)(src)
}

type complex128Copier struct{}

func (complex128Copier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*complex128)(dst)) = *(*complex128)(src)
}

type stringCopier struct{}

func (*stringCopier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*string)(dst)) = *(*string)(src)
}

type funcCopier struct{}

func (*funcCopier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*unsafe.Pointer)(dst)) = *(*unsafe.Pointer)(src)
}

type trivalCopier struct{}

func (*trivalCopier) Copy(src unsafe.Pointer, dst unsafe.Pointer) {
	*((*unsafe.Pointer)(dst)) = *(*unsafe.Pointer)(src)
}

func createCopierOfNative(typ reflect2.Type) Copier {
	switch typ.Kind() {
	case reflect.Bool:
		return &boolCopier{}
	case reflect.Int:
		return &intCopier{}
	case reflect.Int8:
		return &int8Copier{}
	case reflect.Int16:
		return &int16Copier{}
	case reflect.Int32:
		return &int32Copier{}
	case reflect.Int64:
		return &int64Copier{}
	case reflect.Uint:
		return &uintCopier{}
	case reflect.Uint8:
		return &uint8Copier{}
	case reflect.Uint16:
		return &uint16Copier{}
	case reflect.Uint32:
		return &uint32Copier{}
	case reflect.Uint64:
		return &uint64Copier{}
	case reflect.Uintptr:
		return &uintptrCopier{}
	case reflect.Float32:
		return &float32Copier{}
	case reflect.Float64:
		return &float64Copier{}
	case reflect.Complex64:
		return &complex64Copier{}
	case reflect.Complex128:
		return &complex128Copier{}
	case reflect.String:
		return &stringCopier{}
	case reflect.Func:
		return &funcCopier{}
	case reflect.UnsafePointer, reflect.Chan: // Chan and unsafe.Pointer are not deep copied
		return &trivalCopier{}
	default:
		// impossible
		return nil
	}
}
