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
		return _boolCopier
	case reflect.Int:
		return _intCopier
	case reflect.Int8:
		return _int8Copier
	case reflect.Int16:
		return _int16Copier
	case reflect.Int32:
		return _int32Copier
	case reflect.Int64:
		return _int64Copier
	case reflect.Uint:
		return _uintCopier
	case reflect.Uint8:
		return _uint8Copier
	case reflect.Uint16:
		return _uint16Copier
	case reflect.Uint32:
		return _uint32Copier
	case reflect.Uint64:
		return _uint64Copier
	case reflect.Uintptr:
		return _uintptrCopier
	case reflect.Float32:
		return _float32Copier
	case reflect.Float64:
		return _float64Copier
	case reflect.Complex64:
		return _complex64Copier
	case reflect.Complex128:
		return _complex128Copier
	case reflect.String:
		return _stringCopier
	case reflect.Func:
		return _funcCopier
	case reflect.UnsafePointer, reflect.Chan: // Chan and unsafe.Pointer are not deep copied
		return _trivalCopier
	default:
		return nil
	}
}

var _boolCopier = &boolCopier{}
var _intCopier = &intCopier{}
var _int8Copier = &int8Copier{}
var _int16Copier = &int16Copier{}
var _int32Copier = &int32Copier{}
var _int64Copier = &int64Copier{}
var _uintCopier = &uintCopier{}
var _uint8Copier = &uint8Copier{}
var _uint16Copier = &uint16Copier{}
var _uint32Copier = &uint32Copier{}
var _uint64Copier = &uint64Copier{}
var _uintptrCopier = &uintptrCopier{}
var _float32Copier = &float32Copier{}
var _float64Copier = &float64Copier{}
var _complex64Copier = &complex64Copier{}
var _complex128Copier = &complex128Copier{}
var _stringCopier = &stringCopier{}
var _funcCopier = &funcCopier{}
var _trivalCopier = &trivalCopier{}
