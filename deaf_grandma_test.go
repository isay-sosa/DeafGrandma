package DeafGrandma

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func TestDeafGrandma(t *testing.T) {
	Convey("When I speak to my grandma", t, func() {
		speak := "hello grandma!"

		Convey("She answers me with 'Huh?! Speak up, sonny!'", func() {
			So(speakWithGrandma(speak), ShouldEqual, "Huh?! Speak up, sonny!")
		})
	})

	Convey("When I shout to my grandma", t, func() {
		speak := "HELLO GRANDMA!"

		Convey("Her anwser should contain 'No, not since'", func() {
			So(speakWithGrandma(speak), ShouldContainSubstring, "No, not since")
		})
	})
}

func TestRandomYear(t *testing.T) {
	Convey("When running randomYear function", t, func() {
		year, _ := strconv.Atoi(randomYear())

		Convey("We get a year between 1930-1950", func() {
			So(year, ShouldBeBetween, 1930, 1950)
		})
	})
}
