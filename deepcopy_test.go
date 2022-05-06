package deepcopy

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/google/go-cpy/cpy"
	"github.com/mohae/deepcopy"
	"github.com/stretchr/testify/assert"
)

type Basics struct {
	String    string
	Strings   []string
	StringArr [4]string
	Bool      bool
	Bools     []bool
	Byte      byte
	Bytes     []byte
	Int       int
	Ints      []int
	Int8      int8
	Int8s     []int8
	Int16     int16
	Int16s    []int16
	Int32     int32
	Int32s    []int32
	Int64     int64
	Int64s    []int64
	Uint      uint
	Uints     []uint
	Uint8     uint8
	Uint8s    []uint8
	Uint16    uint16
	Uint16s   []uint16
	Uint32    uint32
	Uint32s   []uint32
	Uint64    uint64
	Uint64s   []uint64
	Float32   float32
	Float32s  []float32
	Float64   float64
	Float64s  []float64
	//Complex64   complex64
	//Complex64s  []complex64
	//Complex128  complex128
	//Complex128s []complex128
	Interface  interface{}
	Interfaces []interface{}
}

var src = Basics{
	String:    "kimchi",
	Strings:   []string{"uni", "ika"},
	StringArr: [4]string{"malort", "barenjager", "fernet", "salmiakki"},
	Bool:      true,
	Bools:     []bool{true, false, true},
	Byte:      'z',
	Bytes:     []byte("abc"),
	Int:       42,
	Ints:      []int{0, 1, 3, 4},
	Int8:      8,
	Int8s:     []int8{8, 9, 10},
	Int16:     16,
	Int16s:    []int16{16, 17, 18, 19},
	Int32:     32,
	Int32s:    []int32{32, 33},
	Int64:     64,
	Int64s:    []int64{64},
	Uint:      420,
	Uints:     []uint{11, 12, 13},
	Uint8:     81,
	Uint8s:    []uint8{81, 82},
	Uint16:    160,
	Uint16s:   []uint16{160, 161, 162, 163, 164},
	Uint32:    320,
	Uint32s:   []uint32{320, 321},
	Uint64:    640,
	Uint64s:   []uint64{6400, 6401, 6402, 6403},
	Float32:   32.32,
	Float32s:  []float32{32.32, 33},
	Float64:   64.1,
	Float64s:  []float64{64, 65, 66},
	//Complex64:   complex64(-64 + 12i),
	//Complex64s:  []complex64{complex64(-65 + 11i), complex64(66 + 10i)},
	//Complex128:  complex128(-128 + 12i),
	//Complex128s: []complex128{complex128(-128 + 11i), complex128(129 + 10i)},
	Interfaces: []interface{}{42, true, "pan-galactic"},
}

func Benchmark_GOBDeepCopy(b *testing.B) {
	// use b.N for looping
	for i := 0; i < b.N; i++ {
		var dst Basics
		err := GOBDeepCopy(&dst, &src)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_ReflectDeepCopy(b *testing.B) {
	// use b.N for looping
	for i := 0; i < b.N; i++ {
		dst := deepcopy.Copy(src).(Basics)
		if !dst.Bool {
			b.Error("reflect deep copy failed")
		}
	}
}
func Benchmark_JsonDeepcopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst := &Basics{}
		err := JsonDeeoCopy(dst, src)
		if err != nil || !dst.Bool {
			b.Errorf("json deep copy failed: %v", err)
		}
	}
}
func Benchmark_CustomDeepcopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst := Custom()
		if !dst.Bool {
			b.Error("reflect deep copy failed")
		}
	}
}
func Benchmark_CPY(b *testing.B) {
	var copier = cpy.New(cpy.IgnoreAllUnexported())

	for i := 0; i < b.N; i++ {
		dst := copier.Copy(src).(Basics)
		if !dst.Bool {
			b.Error("reflect deep copy failed")
		}
	}
}

// GOBDeepCopy provides the method to creates a deep copy of whatever is passed to
// it and returns the copy in an interface. The returned value will need to be
// asserted to the correct type.
func GOBDeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func JsonDeeoCopy(dst, src interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, dst)
	if err != nil {
		return err
	}
	return nil
}

func Custom() Basics {
	dst := Basics{String: "kimchi",
		Strings:   []string{"uni", "ika"},
		StringArr: [4]string{"malort", "barenjager", "fernet", "salmiakki"},
		Bool:      true,
		Bools:     []bool{true, false, true},
		Byte:      'z',
		Bytes:     []byte("abc"),
		Int:       42,
		Ints:      []int{0, 1, 3, 4},
		Int8:      8,
		Int8s:     []int8{8, 9, 10},
		Int16:     16,
		Int16s:    []int16{16, 17, 18, 19},
		Int32:     32,
		Int32s:    []int32{32, 33},
		Int64:     64,
		Int64s:    []int64{64},
		Uint:      420,
		Uints:     []uint{11, 12, 13},
		Uint8:     81,
		Uint8s:    []uint8{81, 82},
		Uint16:    160,
		Uint16s:   []uint16{160, 161, 162, 163, 164},
		Uint32:    320,
		Uint32s:   []uint32{320, 321},
		Uint64:    640,
		Uint64s:   []uint64{6400, 6401, 6402, 6403},
		Float32:   32.32,
		Float32s:  []float32{32.32, 33},
		Float64:   64.1,
		Float64s:  []float64{64, 65, 66},
		//Complex64:   complex64(-64 + 12i),
		//Complex64s:  []complex64{complex64(-65 + 11i), complex64(66 + 10i)},
		//Complex128:  complex128(-128 + 12i),
		//Complex128s: []complex128{complex128(-128 + 11i), complex128(129 + 10i)},
		Interfaces: []interface{}{42, true, "pan-galactic"},
	}
	return dst
}

func TestMyCopy(t *testing.T) {
	a := []string{"a", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c"}
	b := []string{}
	Deepcopy(&a, &b)
	assert.Equal(t, a, b)
}
func TestMyCopy1(t *testing.T) {
	a := "xxxxxxxxxxx"
	b := "12"
	Deepcopy(&a, &b)
	assert.Equal(t, a, b)
}

type A struct {
	X int
	Y string
	Z []string
}

func TestMyCopy2(t *testing.T) {
	a := A{5, "t", []string{"x", "z"}}
	b := A{}
	Deepcopy(&a, &b)
	assert.Equal(t, a, b)
}

func BenchmarkMyCopy(b *testing.B) {
	a := []string{"a", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c"}

	for i := 0; i < b.N; i++ {
		b := []string{}
		Deepcopy(&a, &b)
	}
}

func BenchmarkStandardCopy(b *testing.B) {
	a := []string{"a", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c"}

	for i := 0; i < b.N; i++ {
		_, _ = deepcopy.Copy(a).([]string)
	}
}

func BenchmarkSonic(b *testing.B) {
	a := []string{"a", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c", "b", "c", "a", "b", "c"}
	for i := 0; i < b.N; i++ {
		bb, _ := sonic.Marshal(a)
		y := []string{}
		_ = sonic.Unmarshal(bb, &y)

	}
}
