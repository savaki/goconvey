package convey

import (
	"github.com/smartystreets/goconvey/execution"
	"github.com/smartystreets/goconvey/reporting"
)

// assertion is an alias for a function with a signature that the Context.assert()
// method can handle. Any future or custom assertions should conform to this
// method signature. The return value should be an empty string if the assertion
// passes and a well-formed failure message if not.
type assertion func(actual interface{}, expected ...interface{}) string

const assertionSuccess = ""

type Assert func(actual interface{}, assert assertion, expected ...interface{})

type Context struct {
	reporter reporting.Reporter
	runner   execution.Runner
}

func NewContext() *Context {
	self := &Context{}
	self.reporter = buildReporter()
	self.runner = execution.NewRunner()
	self.runner.UpgradeReporter(reporter)
	return self
}

func (self *Context) assertion(actual interface{}, assert assertion, expected ...interface{}) {
	if result := assert(actual, expected...); result == assertionSuccess {
		self.reporter.Report(reporting.NewSuccessReport())
	} else {
		self.reporter.Report(reporting.NewFailureReport(result))
	}
}
