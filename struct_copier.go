package deepcopy

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type structCopier struct {
	fieldTypes []reflect2.StructField
	copiers    []Copier
	numField   int
}

func (s *structCopier) Copy(src, ptr unsafe.Pointer) {
	for i := 0; i < s.numField; i++ {
		//根据struct ptr获取成员的ptr
		fieldPtr := s.fieldTypes[i].UnsafeGet(ptr)
		//根据*struct interface{}获取src的成员的指针的interface{}
		s.copiers[i].Copy(s.fieldTypes[i].UnsafeGet(src), fieldPtr)
	}
}

func createCopierOfStruct(typ reflect2.Type) Copier {
	if typ.Kind() != reflect.Struct {
		return nil
	}
	structType := typ.(*reflect2.UnsafeStructType)
	numField := structType.NumField()
	FieldTypes := make([]reflect2.StructField, numField)
	copiers := make([]Copier, numField)
	for i := 0; i < numField; i++ {
		FieldTypes[i] = structType.Field(i)
		copiers[i] = CopierOf(structType.Field(i).Type())
	}
	return &structCopier{FieldTypes, copiers, numField}
}
