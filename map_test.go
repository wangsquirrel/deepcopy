package deepcopy

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/mohae/deepcopy"
	"github.com/stretchr/testify/assert"
)

type AA struct {
	M map[string]string
	N map[int][]bool
}

func TestMapCopy(t *testing.T) {
	a := AA{
		M: map[string]string{"1": "2", "3": "4"},
		N: map[int][]bool{1: {true, false, true}},
	}
	var b AA = AA{}
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

}

func BenchmarkRawMapCopy(b *testing.B) {
	a := map[int]string{}
	for i := 0; i < 200; i++ {
		a[i] = "111111"
	}
	for i := 0; i < b.N; i++ {
		for i := 0; i < 200; i++ {
			a[i] = "111111"
		}

	}
}

func BenchmarkMapCopy(b *testing.B) {
	a := map[int]string{}
	for i := 0; i < 200; i++ {
		a[i] = "1"
	}
	for i := 0; i < b.N; i++ {
		y := map[int]string{}
		DeepCopy(&a, &y)
		if len(a) != len(y) {
			b.Error("not equal")

		}
	}
}

func BenchmarkStandardMapCopy(b *testing.B) {
	a := map[int]string{}
	for i := 0; i < 200; i++ {
		a[i] = "1"
	}

	for i := 0; i < b.N; i++ {
		_, _ = deepcopy.Copy(a).(map[int]string)

	}
}

func BenchmarkSonicMapCopy(b *testing.B) {
	a := map[int]string{}
	for i := 0; i < 200; i++ {
		a[i] = "1"
	}
	for i := 0; i < b.N; i++ {
		bb, _ := sonic.Marshal(a)
		var y map[int]string
		_ = sonic.Unmarshal(bb, &y)
		if len(a) != len(y) {
			b.Error("")

		}

	}
}
