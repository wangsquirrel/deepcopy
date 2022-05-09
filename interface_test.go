package deepcopy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SS string

func (ss SS) String() string {
	return (string)(ss)
}

func TestIfaceCopy(t *testing.T) {

	s1 := (SS)("555")
	s := (SS)("222")
	var a, b fmt.Stringer

	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	a = (*SS)(nil)
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	a = s
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)
	assert.Equal(t, a.String(), b.String())

	a = &s
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)
	assert.Equal(t, a.String(), b.String())

	a = &s
	b = &s1
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)
	assert.Equal(t, a.String(), b.String())

	a = nil
	b = s1
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

}

func TestEfaceCopy(t *testing.T) {

	var a, b interface{}
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	a = (*string)(nil)
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	s := "ssssss"
	a = s
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	a = &s
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	b = (*int)(nil)
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	a = &s
	b = "ppp"
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)

	a = []interface{}{"aaa"}
	DeepCopy(&a, &b)
	assert.Equal(t, a, b)
}
