package deepcopy

import (
	"testing"
	"time"

	"github.com/modern-go/reflect2"
)

func TestR(t *testing.T) {
	s := []int{1}
	a := reflect2.TypeOf(s)
	tt := a.(*reflect2.UnsafeSliceType)
	t.Log(tt.GetIndex(&s, 0))
	print("\n")
	time.Sleep(time.Second)

}
