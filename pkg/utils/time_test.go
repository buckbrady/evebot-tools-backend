package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCheckTTL(t *testing.T) {
	Convey("Should return true if time is not past", t, func() {
		So(CheckTTL(time.Now().UTC(), 60), ShouldBeTrue)
	})
	Convey("Should return true if time is not past - long", t, func() {
		So(CheckTTL(time.Now().UTC(), 86400), ShouldBeTrue)
	})
	Convey("Should return false if time is past", t, func() {
		So(CheckTTL(time.Now().UTC().Add(-1*time.Hour), 30), ShouldBeFalse)
	})
}
