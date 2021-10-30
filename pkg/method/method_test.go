package method

import (
	"github.com/agiledragon/gomonkey/v2"
	"reflect"
	"testing"
)

func TestCompute(t *testing.T) {
	var c *Computer
	patches := gomonkey.ApplyMethod(reflect.TypeOf(c), "NetworkCompute", func(_ *Computer, a, b int) (int, error) {
		return 2, nil
	})

	defer patches.Reset()

	cp := &Computer{}
	sum, err := cp.Compute(1, 1)
	if sum != 2 || err != nil {
		t.Errorf("expected %v, got %v", 2, sum)
	}

}
