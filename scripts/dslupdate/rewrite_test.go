package main

import (
	"strings"
	"testing"
)

func TestRewrite(t *testing.T) {
	actualLines := rewrite(strings.Split(old, "\n"))
	actual := strings.Join(actualLines, "\n")
	if actual != updated {
		t.Errorf("Rewrite failed. \nExpected:\n%s\nActual:\n%s\n", updated, actual)
	}
}

var old = `
package something

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSomething(t *testing.T) {
	var a int

	Convey("Top-level", t, func() {
		a = 42

		Convey("nested 1", func() {
			So(a, ShouldEqual, 42)
		})

		SkipConvey("skipped", func() {
			Convey("nested and skipped", func() {
				SkipSo(true, ShouldEqual, false)
			})
		})

		Reset(func() {
			// comments should not be modified
		})
	})
}
`

var updated = `
package something

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSomething(t *testing.T) {
	var a int

	Convey("Top-level", t, func(c Context, so Assert) {
		a = 42

		Convey("nested 1", c, func() {
			so(a, ShouldEqual, 42)
		})

		SkipConvey("skipped", c, func() {
			Convey("nested and skipped", c, func() {
				c.Skipso(true, ShouldEqual, false)
			})
		})

		Reset(c, func() {
			// comments should not be modified
		})
	})
}
`
