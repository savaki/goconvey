// Package convey contains all of the public-facing entry points to this project.
// This means that it should never be required of the user to import any other
// packages from this project as they serve internal purposes.
package convey

// Convey is the method intended for use when declaring the scopes
// of a specification. Each scope has a description and a func()
// which may contain other calls to Convey(), Reset() or Should-style
// assertions. Convey calls can be nested as far as you see fit.
//
// IMPORTANT NOTE: The top-level Convey() within a Test method
// must conform to the following signature:
//
//     Convey(description string, t *testing.T, action func(c Context, so Assert))
//
// All other calls should like this:
//
//     Convey(description string, c Context, action func())
//
// Don't worry, the goconvey will panic if you get it wrong so you can fix it.
//
// See the examples package for, well, examples.
func Convey(description string, context interface{}, action interface{}) {
}

// SkipConvey is analagous to Convey except that the scope is not executed
// (which means that child scopes defined within this scope are not run either).
// The reporter will be notified that this step was skipped.
func SkipConvey(items ...interface{}) {
}

// Reset registers a cleanup function to be run after each Convey()
// in the same scope. See the examples package for a simple use case.
func Reset(action func()) {
}
