package global_var

import (
	"github.com/agiledragon/gomonkey/v2"
	"testing"
)


func TestGlobalVar(t *testing.T){
	patches := gomonkey.ApplyGlobalVar(&num, 12)
	defer patches.Reset()

	if num != 12 {
		t.Errorf("expected %v, got %v", 12, num)
	}
}
