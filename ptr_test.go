package deepcopy

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/mohae/deepcopy"
	"github.com/stretchr/testify/assert"
)

type S struct {
	Y *int
	X *[]string
	M *map[string]int
	N *string
}

type S1 struct {
	a int
}

func TestPtrPtrCopy(t *testing.T) {

	x := S1{7}
	y := &x
	var a **S1 = &y
	var b **S1
	DeepCopy(&a, &b)
	assert.Equal(t, **a, **b)
	assert.Equal(t, *a, *b)
	assert.Equal(t, a, b)
}

func TestPtrCopy(t *testing.T) {
	i := 3
	x := "7777777777777777"
	a := &S{&i, &[]string{"1", "2"}, &map[string]int{"0": 9999}, &x}
	b := (*S)(nil)
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)
	assert.Equal(t, *a, *b)
}

func BenchmarkPtrCopy(b *testing.B) {
	i := 3
	x := "7777777777777777"
	a := &S{&i, &[]string{"1", "2"}, &map[string]int{"0": 9999}, &x}
	for i := 0; i < b.N; i++ {
		b := &S{}
		DeepCopy(&a, &b)
	}
}

func BenchmarkMohaePtrCopy(b *testing.B) {
	i := 3
	x := "7777777777777777"
	a := &S{&i, &[]string{"1", "2"}, &map[string]int{"0": 9999}, &x}

	for i := 0; i < b.N; i++ {
		_, _ = deepcopy.Copy(a).(*S)
	}
}

func BenchmarkSonicPtrCopy(b *testing.B) {
	i := 3
	x := "7777777777777777"
	a := &S{&i, &[]string{"1", "2"}, &map[string]int{"0": 9999}, &x}
	for i := 0; i < b.N; i++ {
		bb, _ := sonic.Marshal(a)
		var y *S
		_ = sonic.Unmarshal(bb, &y)

	}
}
