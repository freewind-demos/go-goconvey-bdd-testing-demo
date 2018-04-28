package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some string with a starting name", t, func() {
		name := "goconvey"

		Convey("When greeting word is generated", func() {
			greeting := Greeting(name)

			Convey("The greeting word should be a hello", func() {
				So(greeting, ShouldEqual, "Hello, goconvey!")
			})
		})
	})
}
