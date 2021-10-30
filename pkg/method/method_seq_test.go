package method

import (
	"errors"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

func TestSeqCompute(t *testing.T) {
	com := &Computer{}
	convey.Convey(	"TestSeqCompute", t, func() {
		convey.Convey("default times is 1", func() {
			merr := errors.New("mock error")
			outputs := []gomonkey.OutputCell{
				{Values: gomonkey.Params{1, nil}},
				{Values: gomonkey.Params{2, nil}},
				{Values: gomonkey.Params{3, nil}},
				{Values: gomonkey.Params{4, merr}},
			}
			patches := gomonkey.ApplyMethodSeq(reflect.TypeOf(com), "Compute", outputs)
			defer patches.Reset()

			sum, err := com.Compute(1, 1)
			convey.So(sum, convey.ShouldEqual, 1)
			convey.So(err, convey.ShouldEqual, nil)

			sum, err = com.Compute(1, 1)
			convey.So(sum, convey.ShouldEqual, 2)
			convey.So(err, convey.ShouldEqual, nil)

			sum, err = com.Compute(1, 1)
			convey.So(sum, convey.ShouldEqual, 3)
			convey.So(err, convey.ShouldEqual, nil)

			sum, err = com.Compute(1, 1)
			convey.So(sum, convey.ShouldEqual, 4)
			convey.So(err, convey.ShouldEqual, merr)
		})
	})

}
