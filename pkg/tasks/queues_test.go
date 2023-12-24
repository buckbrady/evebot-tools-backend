package tasks

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetQueueOpts(t *testing.T) {
	Convey("TestGetQueueOpts", t, func() {
		Convey("Should return valid settings", func() {
			q := GetQueueOpts("group", "name", 1)
			So(q.GetName(), ShouldEqual, "group_name")
			So(q.GetPriority(), ShouldEqual, 1)
			So(q.GetQueue(), ShouldNotBeNil)
		})
	})
}
