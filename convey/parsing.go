package convey

import (
	"github.com/smartystreets/goconvey/execution"
	"github.com/smartystreets/goconvey/gotest"
)

func discover(items []interface{}) *execution.Registration {
	ensureEnough(items)

	name := parseName(items)
	test := parseGoTest(items)
	action := parseAction(items, test)

	return execution.NewRegistration(name, action, test)
}
func ensureEnough(items []interface{}) {
	if len(items) < 2 {
		panic(parseError)
	}
}
func parseName(items []interface{}) string {
	if name, isType := items[0].(string); isType {
		return name
	}
	panic(parseError)
}
func parseGoTest(items []interface{}) gotest.T {
	if test, isType := items[1].(gotest.T); isType {
		return test
	}
	return nil
}
func parseAction(items []interface{}, test gotest.T) *execution.Action {
	var index = 1
	if test != nil {
		index = 2
	}

	if likelyHasInlineAssertionFunc(items, index) {
		return execution.NewAction(func() {
			So(items[index+1],
				items[index+2].(func(actual interface{}, expected ...interface{}) string),
				items[index+3:]...)
		})
	}
	if action, isType := items[index].(func()); isType {
		return execution.NewAction(action)
	}
	if items[index] == nil {
		return execution.NewSkippedAction(skipReport)
	}
	panic(parseError)
}

func likelyHasInlineAssertionFunc(items []interface{}, index int) bool {
	return len(items) >= index+minimumAssertionArgumentsOffset
}

const minimumAssertionArgumentsOffset = 2
const parseError = "You must provide a name (string), then a *testing.T (if in outermost scope), and then an action (func())."
