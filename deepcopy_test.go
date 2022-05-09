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

func TestCopy(t *testing.T) {
	b := Basics{}
	err := DeepCopy(&src, &b)
	assert.Equal(t, src, b)
	assert.Nil(t, err)
}

func TestCopyParallel(t *testing.T) {
	go func() {
		for i := 0; i < 100; i++ {
			b := Basics{}
			err := DeepCopy(&src, &b)
			assert.Equal(t, src, b)
			assert.Nil(t, err)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			b := BasicsJson{}
			err := DeepCopy(&srcJson, &b)
			assert.Equal(t, srcJson, b)
			assert.Nil(t, err)
		}
	}()
}

func Benchmark_Copy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst BasicsJson
		DeepCopy(&srcJson, &dst)
	}
}

func Benchmark_SonicCopy(b *testing.B) {

	for i := 0; i < b.N; i++ {
		bb, _ := sonic.Marshal(srcJson)
		y := BasicsJson{}
		_ = sonic.Unmarshal(bb, &y)

	}
}

func Benchmark_GOBCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var dst BasicsJson
		var buf bytes.Buffer
		if err := gob.NewEncoder(&buf).Encode(srcJson); err != nil {
			b.Error(err)
		}
		err := gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(&dst)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_ReflectCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst := deepcopy.Copy(srcJson).(BasicsJson)
		if !dst.Bool {
			b.Error("reflect deep copy failed")
		}
	}
}
func Benchmark_JsonCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst := &BasicsJson{}
		bb, err := json.Marshal(srcJson)
		if err != nil {
			b.Errorf("json deep copy failed: %v", err)
		}
		err = json.Unmarshal(bb, dst)
		if err != nil {
			b.Errorf("json deep copy failed: %v", err)
		}
		if !dst.Bool {
			b.Errorf("json deep copy failed: %v", err)
		}
	}
}

func Benchmark_CPYCopy(b *testing.B) {
	var copier = cpy.New(cpy.IgnoreAllUnexported())

	for i := 0; i < b.N; i++ {
		dst := copier.Copy(srcJson).(BasicsJson)
		if !dst.Bool {
			b.Error("reflect deep copy failed")
		}
	}
}
func Benchmark_RawCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dst := Custom()
		if !dst.Bool {
			b.Error("reflect deep copy failed")
		}
	}
}
func Custom() BasicsJson {
	dst := BasicsJson{String: "kimchi",
		Strings:    []string{"uni", "ika"},
		StringArr:  [4]string{"malort", "barenjager", "fernet", "salmiakki"},
		Bool:       true,
		Bools:      []bool{true, false, true},
		Byte:       'z',
		Bytes:      []byte("abc"),
		Int:        42,
		Ints:       []int{0, 1, 3, 4},
		Int8:       8,
		Int8s:      []int8{8, 9, 10},
		Int16:      16,
		Int16s:     []int16{16, 17, 18, 19},
		Int32:      32,
		Int32s:     []int32{32, 33},
		Int64:      64,
		Int64s:     []int64{64},
		Uint:       420,
		Uints:      []uint{11, 12, 13},
		Uint8:      81,
		Uint8s:     []uint8{81, 82},
		Uint16:     160,
		Uint16s:    []uint16{160, 161, 162, 163, 164},
		Uint32:     320,
		Uint32s:    []uint32{320, 321},
		Uint64:     640,
		Uint64s:    []uint64{6400, 6401, 6402, 6403},
		Float32:    32.32,
		Float32s:   []float32{32.32, 33},
		Float64:    64.1,
		Float64s:   []float64{64, 65, 66},
		Interfaces: []interface{}{42, true, "pan-galactic"},
	}
	return dst
}

type Basics struct {
	String      string
	Strings     []string
	StringArr   [4]string
	Bool        bool
	Bools       []bool
	Byte        byte
	Bytes       []byte
	Int         int
	Ints        []int
	Int8        int8
	Int8s       []int8
	Int16       int16
	Int16s      []int16
	Int32       int32
	Int32s      []int32
	Int64       int64
	Int64s      []int64
	Uint        uint
	Uints       []uint
	Uint8       uint8
	Uint8s      []uint8
	Uint16      uint16
	Uint16s     []uint16
	Uint32      uint32
	Uint32s     []uint32
	Uint64      uint64
	Uint64s     []uint64
	Float32     float32
	Float32s    []float32
	Float64     float64
	Float64s    []float64
	Complex64   complex64
	Complex64s  []complex64
	Complex128  complex128
	Complex128s []complex128
	Map         map[int]int
	Interface   interface{}
	Interfaces  []interface{}
}

type BasicsJson struct {
	String     string
	Strings    []string
	StringArr  [4]string
	Bool       bool
	Bools      []bool
	Byte       byte
	Bytes      []byte
	Int        int
	Ints       []int
	Int8       int8
	Int8s      []int8
	Int16      int16
	Int16s     []int16
	Int32      int32
	Int32s     []int32
	Int64      int64
	Int64s     []int64
	Uint       uint
	Uints      []uint
	Uint8      uint8
	Uint8s     []uint8
	Uint16     uint16
	Uint16s    []uint16
	Uint32     uint32
	Uint32s    []uint32
	Uint64     uint64
	Uint64s    []uint64
	Float32    float32
	Float32s   []float32
	Float64    float64
	Float64s   []float64
	Interface  interface{}
	Interfaces []interface{}
}

var src = Basics{
	String:      "kimchi",
	Strings:     []string{"uni", "ika", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager"},
	StringArr:   [4]string{"malort", "barenjager", "fernet", "salmiakki"},
	Bool:        true,
	Bools:       []bool{true, false, true},
	Byte:        'z',
	Bytes:       []byte("abcbarenjager"),
	Int:         42,
	Ints:        []int{0, 1, 3, 4},
	Int8:        8,
	Int8s:       []int8{8, 9, 10},
	Int16:       16,
	Int16s:      []int16{16, 17, 18, 19},
	Int32:       32,
	Int32s:      []int32{32, 33},
	Int64:       64,
	Int64s:      []int64{64},
	Uint:        420,
	Uints:       []uint{11, 12, 13},
	Uint8:       81,
	Uint8s:      []uint8{81, 82},
	Uint16:      160,
	Uint16s:     []uint16{160, 161, 162, 163, 164},
	Uint32:      320,
	Uint32s:     []uint32{320, 321},
	Uint64:      640,
	Uint64s:     []uint64{6400, 6401, 6402, 6403},
	Float32:     32.32,
	Float32s:    []float32{32.32, 33},
	Float64:     64.1,
	Float64s:    []float64{64, 65, 66},
	Complex64:   complex64(-64 + 12i),
	Complex64s:  []complex64{complex64(-65 + 11i), complex64(66 + 10i)},
	Complex128:  complex128(-128 + 12i),
	Complex128s: []complex128{complex128(-128 + 11i), complex128(129 + 10i)},
	Map:         map[int]int{1: 2, 3: 4, 5: 6, 7: 8, 9: 10},
	Interfaces:  []interface{}{42, true, "pan-galactic", '3', []float32{32.32, 33}},
}

var srcJson = BasicsJson{
	String:     "kimchi",
	Strings:    []string{"uni", "ika", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager", "barenjager"},
	StringArr:  [4]string{"malort", "barenjager", "fernet", "salmiakki"},
	Bool:       true,
	Bools:      []bool{true, false, true},
	Byte:       'z',
	Bytes:      []byte("abcbarenjager"),
	Int:        42,
	Ints:       []int{0, 1, 3, 4},
	Int8:       8,
	Int8s:      []int8{8, 9, 10},
	Int16:      16,
	Int16s:     []int16{16, 17, 18, 19},
	Int32:      32,
	Int32s:     []int32{32, 33},
	Int64:      64,
	Int64s:     []int64{64},
	Uint:       420,
	Uints:      []uint{11, 12, 13},
	Uint8:      81,
	Uint8s:     []uint8{81, 82},
	Uint16:     160,
	Uint16s:    []uint16{160, 161, 162, 163, 164},
	Uint32:     320,
	Uint32s:    []uint32{320, 321},
	Uint64:     640,
	Uint64s:    []uint64{6400, 6401, 6402, 6403},
	Float32:    32.32,
	Float32s:   []float32{32.32, 33},
	Float64:    64.1,
	Float64s:   []float64{64, 65, 66},
	Interfaces: []interface{}{42, true, "pan-galactic", '3', []float32{32.32, 33}},
}
