package snuffler

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFile(t *testing.T) {
	type Conf struct {
		Test string
	}

	Convey("When adding files", t, func() {
		var conf Conf
		snoot := New(&conf)

		Convey("When adding a file", func() {

			Convey("Adding a non-existant errors", func() {
				err := snoot.AddFile("bad wolf")
				So(err, ShouldNotBeNil)
				So(len(snoot.files), ShouldEqual, 0)
			})

			Convey("One can add an existing file", func() {
				err := snoot.AddFile("_testdata/simple.yaml")
				So(err, ShouldBeNil)
				So(len(snoot.files), ShouldEqual, 1)
				So(snoot.files[0].fileName, ShouldEqual, "_testdata/simple.yaml")
			})
		})

		Convey("When maybe-adding a file", func() {

			Convey("Adding a non-existant file does not error", func() {
				snoot.MaybeAddFile("bad wolf")
				So(len(snoot.files), ShouldEqual, 0)
			})

			Convey("One can add an existing file", func() {
				snoot.MaybeAddFile("_testdata/simple.yaml")
				So(len(snoot.files), ShouldEqual, 1)
				So(snoot.files[0].fileName, ShouldEqual, "_testdata/simple.yaml")
			})
		})

		Convey("One can add multiple files", func() {
			snoot.MaybeAddFile("_testdata/simple.yaml")
			snoot.MaybeAddFile("_testdata/override.yaml")
			So(len(snoot.files), ShouldEqual, 2)
			So(snoot.files[0].fileName, ShouldEqual, "_testdata/simple.yaml")
			So(snoot.files[1].fileName, ShouldEqual, "_testdata/override.yaml")
		})

		Convey("Adding files is idempotent", func() {
			snoot.MaybeAddFile("_testdata/simple.yaml")
			snoot.MaybeAddFile("_testdata/simple.yaml")
			So(len(snoot.files), ShouldEqual, 1)
			So(snoot.files[0].fileName, ShouldEqual, "_testdata/simple.yaml")
		})
	})

	Convey("When adding globs", t, func() {
		var conf Conf
		snoot := New(&conf)

		Convey("All files matching the glob are added (also tests GetFileMatchList)", func() {
			snoot.AddGlob("_testdata/*.yaml")
			So(len(snoot.files), ShouldEqual, 3)

			matchList := snoot.GetFileMatchList()
			So(matchList, ShouldResemble, []string{
				"_testdata/overridden.yaml",
				"_testdata/override.yaml",
				"_testdata/simple.yaml",
			})

			Convey("But files that don't are not", func() {
				So(matchList, ShouldNotContain, "_testdata/bad-wolf")
			})
		})
	})

	Convey("It can guess at file type", t, func() {
		var conf Conf
		snoot := New(&conf)

		snoot.AddFile("_testdata/simple.yaml")
		snoot.AddFile("_testdata/simple.toml")
		snoot.AddFile("_testdata/simple.json")
		snoot.AddFile("_testdata/bad-wolf")

		var matches []_fileType
		for _, match := range snoot.files {
			matches = append(matches, match.fileType)
		}

		So(matches, ShouldResemble, []_fileType{
			yamlType,
			tomlType,
			jsonType,
			unknownType,
		})
	})
}
