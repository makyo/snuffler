package snuffler

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSnuffler(t *testing.T) {
	type Conf struct {
		Test string
	}

	Convey("When creating a new Snuffler", t, func() {
		var conf Conf

		Convey("The new Snuffler contains the config object we passed", func() {
			snoot := New(&conf)
			So(snoot.conf, ShouldEqual, &conf)
		})
	})

	Convey("Snufflers can both", t, func() {
		var conf Conf
		snoot := New(&conf)

		Convey("Snuffle...", func() {
			err := snoot.Snuffle()
			So(err, ShouldBeNil)
		})

		Convey("...and Snorfle", func() {
			var conf2 Conf
			err := snoot.Snorfle(&conf2)
			So(err, ShouldBeNil)
		})
	})
}
