package convey

/*
func TestSingleScope(t *testing.T) {
	output := prepare()

	Convey("hi", t, func(c *Context, so Assert) {
		output += "done"
	})

	expectEqual(t, "done", output)
}

func TestSingleScopeWithMultipleConveys(t *testing.T) {
	output := prepare()

	Convey("1", t, func(c *Context, so Assert) {
		output += "1"
	})

	Convey("2", t, func(c *Context, so Assert) {
		output += "2"
	})

	expectEqual(t, "12", output)
}

func TestNestedScopes(t *testing.T) {
	output := prepare()

	Convey("a", t, func(c *Context, so Assert) {
		output += "a "

		Convey("bb", c, func() {
			output += "bb "

			Convey("ccc", c, func() {
				output += "ccc | "
			})
		})
	})

	expectEqual(t, "a bb ccc | ", output)
}

func TestNestedScopesWithIsolatedExecution(t *testing.T) {
	output := prepare()

	Convey("a", t, func(c *Context, so Assert) {
		output += "a "

		Convey("aa", c, func() {
			output += "aa "

			Convey("aaa", c, func() {
				output += "aaa | "
			})

			Convey("aaa1", c, func() {
				output += "aaa1 | "
			})
		})

		Convey("ab", c, func() {
			output += "ab "

			Convey("abb", c, func() {
				output += "abb | "
			})
		})
	})

	expectEqual(t, "a aa aaa | a aa aaa1 | a ab abb | ", output)
}

func TestSingleScopeWithConveyAndNestedReset(t *testing.T) {
	output := prepare()

	Convey("1", t, func(c *Context, so Assert) {
		output += "1"

		Reset(c, func() {
			output += "a"
		})
	})

	expectEqual(t, "1a", output)
}

func TestSingleScopeWithMultipleRegistrationsAndReset(t *testing.T) {
	output := prepare()

	Convey("reset after each nested convey", t, func(c *Context, so Assert) {
		Convey("first output", c, func() {
			output += "1"
		})

		Convey("second output", c, func() {
			output += "2"
		})

		Reset(c, func() {
			output += "a"
		})
	})

	expectEqual(t, "1a2a", output)
}

func TestSingleScopeWithMultipleRegistrationsAndMultipleResets(t *testing.T) {
	output := prepare()

	Convey("each reset is run at end of each nested convey", t, func(c *Context, so Assert) {
		Convey("1", c, func() {
			output += "1"
		})

		Convey("2", c, func() {
			output += "2"
		})

		Reset(c, func() {
			output += "a"
		})

		Reset(c, func() {
			output += "b"
		})
	})

	expectEqual(t, "1ab2ab", output)
}

func TestPanicAtHigherLevelScopePreventsChildScopesFromRunning(t *testing.T) {
	output := prepare()

	Convey("This step panics", t, func(c *Context, so Assert) {
		Convey("this should NOT be executed", c, func() {
			output += "1"
		})

		panic("Hi")
	})

	expectEqual(t, "", output)
}

func TestPanicInChildScopeDoes_NOT_PreventExecutionOfSiblingScopes(t *testing.T) {
	output := prepare()

	Convey("This is the parent", t, func(c *Context, so Assert) {
		Convey("This step panics", c, func() {
			panic("Hi")
			output += "1"
		})

		Convey("This sibling should execute", c, func() {
			output += "2"
		})
	})

	expectEqual(t, "2", output)
}

func TestResetsAreAlwaysExecutedAfterScopePanics(t *testing.T) {
	output := prepare()

	Convey("This is the parent", t, func(c *Context, so Assert) {
		Convey("This step panics", c, func() {
			panic("Hi")
			output += "1"
		})

		Convey("This sibling step does not panic", c, func() {
			output += "a"

			Reset(c, func() {
				output += "b"
			})
		})

		Reset(c, func() {
			output += "2"
		})
	})

	expectEqual(t, "2ab2", output)
}

func TestSkipTopLevel(t *testing.T) {
	output := prepare()

	SkipConvey("hi", t, c, func() {
		output += "This shouldn't be executed!"
	})

	expectEqual(t, "", output)
}

func TestSkipNestedLevel(t *testing.T) {
	output := prepare()

	Convey("hi", t, func(c *Context, so Assert) {
		output += "yes"

		SkipConvey("bye", c, func() {
			output += "no"
		})
	})

	expectEqual(t, "yes", output)
}

func TestSkipNestedLevelSkipsAllChildLevels(t *testing.T) {
	output := prepare()

	Convey("hi", t, func(c *Context, so Assert) {
		output += "yes"

		SkipConvey("bye", c, func() {
			output += "no"

			Convey("byebye", c, func() {
				output += "no-no"
			})
		})
	})

	expectEqual(t, "yes", output)
}

func TestIterativeConveys(t *testing.T) {
	output := prepare()

	Convey("Test", t, func(c *Context, so Assert) {
		for x := 0; x < 10; x++ {
			y := strconv.Itoa(x)

			Convey(y, func() {
				output += y
			})
		}
	})

	expectEqual(t, "0123456789", output)
}

func prepare() string {
	runner = execution.NewRunner()
	reporting.QuietMode()
	return ""
}
*/
