package DeafGrandma

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDeafGrandma(t *testing.T) {
	Convey("When I speak to my grandma", t, func() {
		speak := "hello grandma!"

		Convey("She answers with 'HUH?! SPEAK UP, SONNY!'", func() {
			So(speakWithGrandma(speak), ShouldEqual, "HUH?! SPEAK UP, SONNY!")
		})
	})
}
