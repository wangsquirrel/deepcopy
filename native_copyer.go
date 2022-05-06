package deepcopy

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type boolCopier struct{}

func (*boolCopier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*bool)(ptr)) = *src.(*bool)
}

type intCopier struct{}

func (*intCopier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*int)(ptr)) = *src.(*int)
}

type int8Copier struct{}

func (*int8Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*int8)(ptr)) = *src.(*int8)
}

type int16Copier struct{}

func (*int16Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*int16)(ptr)) = *src.(*int16)
}

type int32Copier struct{}

func (*int32Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*int32)(ptr)) = *src.(*int32)
}

type int64Copier struct{}

func (*int64Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*int64)(ptr)) = *src.(*int64)
}

type uintCopier struct{}

func (*uintCopier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*uint)(ptr)) = *src.(*uint)
}

type uint8Copier struct{}

func (*uint8Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*uint8)(ptr)) = *src.(*uint8)
}

type uint16Copier struct{}

func (*uint16Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*uint16)(ptr)) = *src.(*uint16)

}

type uint32Copier struct{}

func (*uint32Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*uint32)(ptr)) = *src.(*uint32)
}

type uint64Copier struct{}

func (*uint64Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*uint64)(ptr)) = *src.(*uint64)
}

type uintptrCopier struct{}

func (*uintptrCopier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*uintptr)(ptr)) = *src.(*uintptr)
}

type float32Copier struct{}

func (*float32Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*float32)(ptr)) = *src.(*float32)
}

type float64Copier struct{}

func (*float64Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*float64)(ptr)) = *src.(*float64)
}

type complex64Copier struct{}

func (complex64Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*complex64)(ptr)) = *src.(*complex64)
}

type complex128Copier struct{}

func (complex128Copier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*complex128)(ptr)) = *src.(*complex128)
}

type stringCopier struct{}

func (*stringCopier) Copy(src interface{}, ptr unsafe.Pointer) {
	*((*string)(ptr)) = *src.(*string)
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
	default:
		/*
			Array
			Chan
			Func
			Interface
			Map
			Ptr
			Slice
			Struct
			UnsafePointer
		*/
		return nil
	}
}
