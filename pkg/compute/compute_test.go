package compute

import (
	"github.com/agiledragon/gomonkey/v2"
	"testing"
)

func TestCompute(t *testing.T) {
	patches := gomonkey.ApplyFunc(networkCompute, func(a, b int) (int, error) {
		return 2, nil
	})

	defer patches.Reset()

	sum, err := Compute(1, 1)
	if sum != 2 || err != nil {
		t.Errorf("expected %v, got %v", 2, sum)
	}

}

func TestFunc(t *testing.T){
	info1 := "2"
	info2 := "3"
	info3 := "4"
	outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{info1, nil}},// 模拟函数的第1次输出
		{Values: gomonkey.Params{info2, nil}},// 模拟函数的第2次输出
		{Values: gomonkey.Params{info3, nil}},// 模拟函数的第3次输出
	}
	patches := gomonkey.ApplyFuncSeq(Compute, outputs)
	defer patches.Reset()

	output, err := Compute(1,1)
	if output != 2 || err != nil {
		t.Errorf("expected %v, got %v", 2, output)
	}

	output, err = Compute(1,2)
	if output != 3 || err != nil {
		t.Errorf("expected %v, got %v", 2, output)
	}

	output, err = Compute(1,3)
	if output != 4 || err != nil {
		t.Errorf("expected %v, got %v", 2, output)
	}

}
