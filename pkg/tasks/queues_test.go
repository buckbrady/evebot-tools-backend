package tasks

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetQueueOpts(t *testing.T) {
	Convey("TestGetQueueOpts", t, func() {
		Convey("Should return valid settings", func() {
			q := GetQueueOpts("name", 1)
			So(q.GetName(), ShouldEqual, "name")
			So(q.GetPriority(), ShouldEqual, 1)
			So(q.GetQueue(), ShouldNotBeNil)
		})
	})
}
