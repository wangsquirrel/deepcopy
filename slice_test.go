package deepcopy

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/mohae/deepcopy"
	"github.com/stretchr/testify/assert"
)

func TestSliceyCopy(t *testing.T) {

	a := []string{"a", "b", "c", "a"}
	var b []string = nil
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	b[0] = "0"
	assert.NotEqual(t, a, b)

	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	s := "a"
	x := [2]*string{&s, &s}

	var y [2]*string
	DeepCopy(&x, &y)
	assert.Equal(t, x, y)

	*y[0] = "999"
	assert.NotEqual(t, x, y)
	assert.NotEqual(t, *y[0], *y[1])

	*x[0] = "1"
	assert.Equal(t, *x[0], *x[1])
}

func TestArrayCopy(t *testing.T) {
	a := [999]string{"1", "2", "3"}
	var b [999]string
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	b = [999]string{"1"}
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	c := [998]string{"1"}
	err := DeepCopy(&a, &c)
	assert.Error(t, err)
}

func BenchmarkSliceCopy(b *testing.B) {
	var a []string
	for i := 0; i < 200; i++ {
		a = append(a, "a")
	}
	for i := 0; i < b.N; i++ {
		var y []string
		DeepCopy(&a, &y)
		if len(a) != len(y) {
			b.Error("aa")
		}
	}
}

func BenchmarkMohaeSliceCopy(b *testing.B) {
	var a []string
	for i := 0; i < 200; i++ {
		a = append(a, "a")
	}
	for i := 0; i < b.N; i++ {
		_, _ = deepcopy.Copy(a).([]string)
	}
}

func BenchmarkSonicSliceCopy(b *testing.B) {
	var a []string
	for i := 0; i < 200; i++ {
		a = append(a, "a")
	}
	for i := 0; i < b.N; i++ {
		var y []string
		bb, _ := sonic.Marshal(a)
		err := sonic.Unmarshal(bb, &y)
		if err != nil {
			b.Error()
		}
	}
}
