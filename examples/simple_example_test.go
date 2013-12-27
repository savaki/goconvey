package examples

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Subject: Integer incrementation and decrementation", t, func(c *Context, so Assert) {
		var x int

		Convey("Given a starting integer value", c, func() {
			x = 42

			Convey("When incremented", c, func() {
				x++

				Convey("The value should be greater by one", c, func() {
					so(x, ShouldEqual, 43)
				})
				Convey("The value should NOT be what it used to be", c, func() {
					so(x, ShouldNotEqual, 42)
				})
			})
			Convey("When decremented", c, func() {
				x--

				Convey("The value should be lesser by one", c, func() {
					so(x, ShouldEqual, 41)
				})
				Convey("The value should NOT be what it used to be", c, func() {
					so(x, ShouldNotEqual, 42)
				})
			})
			Reset(c, func() {
				x = 0
			})
		})
	})
}
