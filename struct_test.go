package deepcopy

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/mohae/deepcopy"
	"github.com/stretchr/testify/assert"
)

type ST struct {
	A int
	B ST2
	X int
	Y string
	Z []string
}
type ST2 struct {
	C int
	D []int
	E string
}

func TestStructCopy(t *testing.T) {
	a := ST{5, ST2{2, []int{1, 2, 3, 4, 5, 6, 7}, "iiiiiiiiiiiiiiiiiiiii"}, 1000000, "RRRRRRRRRR", []string{"x", "z", "b", "c", "b", "c", "a", "b", "c"}}
	b := ST{}
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)
}

func BenchmarkRawStructCopy(b *testing.B) {
	_ = ST{5, ST2{2, []int{1, 2, 3, 4, 5, 6, 7}, "iiiiiiiiiii iiiiiiiiii"}, 1000000, "RRRRRRRRRR", []string{"x", "z", "b", "c", "b", "c", "a", "b", "c"}}
	for i := 0; i < b.N; i++ {
		a := ST{5, ST2{2, []int{1, 2, 3, 4, 5, 6, 7}, "iiiiiiiiiiiiiiiiiiiii"}, 1000000, "RRRRRRRRRR", []string{"x", "z", "b", "c", "b", "c", "a", "b", "c"}}
		_ = a

	}
}

func BenchmarkStructCopy(b *testing.B) {
	a := ST{5, ST2{2, []int{1, 2, 3, 4, 5, 6, 7}, "iiiiiiiiiiiiiiiiiiiii"}, 1000000, "RRRRRRRRRR", []string{"x", "z", "b", "c", "b", "c", "a", "b", "c"}}
	for i := 0; i < b.N; i++ {
		b := ST{}
		DeepCopy(&a, &b)
	}
}

func BenchmarkMohaeStructCopy(b *testing.B) {
	a := ST{5, ST2{2, []int{1, 2, 3, 4, 5, 6, 7}, "iiiiiiiiiiiiiiiiiiiii"}, 1000000, "RRRRRRRRRR", []string{"x", "z", "b", "c", "b", "c", "a", "b", "c"}}

	for i := 0; i < b.N; i++ {
		_, _ = deepcopy.Copy(a).([]string)
	}
}

func BenchmarkSonicStructCopy(b *testing.B) {
	a := ST{5, ST2{2, []int{1, 2, 3, 4, 5, 6, 7}, "iiiiiiiiiiiiiiiiiiiii"}, 1000000, "RRRRRRRRRR", []string{"x", "z", "b", "c", "b", "c", "a", "b", "c"}}
	for i := 0; i < b.N; i++ {
		bb, _ := sonic.Marshal(a)
		y := ST{}
		_ = sonic.Unmarshal(bb, &y)
	}
}
