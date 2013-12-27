package examples

import (
	"testing"
	"time"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAssertions(t *testing.T) {
	Convey("All assertions should be accessible", t, func(c *Context, so Assert) {
		Convey("Equality assertions should be accessible", c, func() {
			thing1a := thing{a: "asdf"}
			thing1b := thing{a: "asdf"}
			thing2 := thing{a: "qwer"}

			so(1, ShouldEqual, 1)
			so(1, ShouldNotEqual, 2)
			so(1, ShouldAlmostEqual, 1.000000000000001)
			so(1, ShouldNotAlmostEqual, 2, 0.5)
			so(thing1a, ShouldResemble, thing1b)
			so(thing1a, ShouldNotResemble, thing2)
			so(&thing1a, ShouldPointTo, &thing1a)
			so(&thing1a, ShouldNotPointTo, &thing1b)
			so(nil, ShouldBeNil)
			so(1, ShouldNotBeNil)
			so(true, ShouldBeTrue)
			so(false, ShouldBeFalse)
			so(0, ShouldBeZeroValue)
		})

		Convey("Numeric comparison assertions should be accessible", c, func() {
			so(1, ShouldBeGreaterThan, 0)
			so(1, ShouldBeGreaterThanOrEqualTo, 1)
			so(1, ShouldBeLessThan, 2)
			so(1, ShouldBeLessThanOrEqualTo, 1)
			so(1, ShouldBeBetween, 0, 2)
			so(1, ShouldNotBeBetween, 2, 4)
		})

		Convey("Container assertions should be accessible", c, func() {
			so([]int{1, 2, 3}, ShouldContain, 2)
			so([]int{1, 2, 3}, ShouldNotContain, 4)
			so(1, ShouldBeIn, []int{1, 2, 3})
			so(4, ShouldNotBeIn, []int{1, 2, 3})
		})

		Convey("String assertions should be accessible", c, func() {
			so("asdf", ShouldStartWith, "a")
			so("asdf", ShouldNotStartWith, "z")
			so("asdf", ShouldEndWith, "df")
			so("asdf", ShouldNotEndWith, "as")
			so("", ShouldBeBlank)
			so("asdf", ShouldNotBeBlank)
			so("asdf", ShouldContainSubstring, "sd")
			so("asdf", ShouldNotContainSubstring, "af")
		})

		Convey("Panic recovery assertions should be accessible", c, func() {
			so(panics, ShouldPanic)
			so(func() {}, ShouldNotPanic)
			so(panics, ShouldPanicWith, "Goofy Gophers!")
			so(panics, ShouldNotPanicWith, "Guileless Gophers!")
		})

		Convey("Type-checking assertions should be accessible", c, func() {
			so(1, ShouldHaveSameTypeAs, 0)
			so(1, ShouldNotHaveSameTypeAs, "1")
		})

		Convey("Time assertions should be accessible", c, func() {
			january1, _ := time.Parse(timeLayout, "2013-01-01 00:00")
			january2, _ := time.Parse(timeLayout, "2013-01-02 00:00")
			january3, _ := time.Parse(timeLayout, "2013-01-03 00:00")
			january4, _ := time.Parse(timeLayout, "2013-01-04 00:00")
			january5, _ := time.Parse(timeLayout, "2013-01-05 00:00")
			oneDay, _ := time.ParseDuration("24h0m0s")

			so(january1, ShouldHappenBefore, january4)
			so(january1, ShouldHappenOnOrBefore, january1)
			so(january2, ShouldHappenAfter, january1)
			so(january2, ShouldHappenOnOrAfter, january2)
			so(january3, ShouldHappenBetween, january2, january5)
			so(january3, ShouldHappenOnOrBetween, january3, january5)
			so(january1, ShouldNotHappenOnOrBetween, january2, january5)
			so(january2, ShouldHappenWithin, oneDay, january3)
			so(january5, ShouldNotHappenWithin, oneDay, january1)
			so([]time.Time{january1, january2}, ShouldBeChronological)
		})
	})
}

type thing struct {
	a string
}

func panics() {
	panic("Goofy Gophers!")
}

const timeLayout = "2006-01-02 15:04"
