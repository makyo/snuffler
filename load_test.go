package snuffler

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoad(t *testing.T) {
	type D struct {
		A int
		B string
		C []string
	}
	type E struct {
		A int
		B string
		C []string
		F string
	}
	type Conf struct {
		A int
		B string
		C []string
		D D
		E []E
	}

	simple := Conf{
		A: 1,
		B: "b",
		C: []string{"a", "b", "c"},
		D: D{
			A: 1,
			B: "b",
			C: []string{"a", "b", "c"},
		},
		E: []E{
			E{
				A: 1,
				B: "b",
				C: []string{"a", "b", "c"},
			},
			E{
				F: "ggg",
			},
		},
	}

	Convey("When loading a single file", t, func() {

		Convey("It should load based on specified filetypes, such as", func() {

			Convey("YAML", func() {
				cf := &configFile{
					fileName: "_testdata/simple.yaml",
					fileType: yamlType,
				}

				var conf Conf

				err := unmarshalFile(cf, &conf)
				So(err, ShouldBeNil)
				So(conf, ShouldResemble, simple)
			})

			Convey("TOML", func() {
				cf := &configFile{
					fileName: "_testdata/simple.toml",
					fileType: tomlType,
				}

				var conf Conf

				err := unmarshalFile(cf, &conf)
				So(err, ShouldBeNil)
				So(conf, ShouldResemble, simple)
			})

			Convey("JSON", func() {
				cf := &configFile{
					fileName: "_testdata/simple.json",
					fileType: jsonType,
				}

				var conf Conf

				err := unmarshalFile(cf, &conf)
				So(err, ShouldBeNil)
				So(conf, ShouldResemble, simple)
			})
		})

		Convey("It should try all types if it doesn't know the filetype", func() {

			Convey("And succeed with", func() {

				Convey("YAML", func() {
					cf := &configFile{
						fileName: "_testdata/simple.yaml",
						fileType: unknownType,
					}

					var conf Conf

					err := unmarshalFile(cf, &conf)
					So(err, ShouldBeNil)
					So(conf, ShouldResemble, simple)
				})

				Convey("TOML", func() {
					cf := &configFile{
						fileName: "_testdata/simple.toml",
						fileType: unknownType,
					}

					var conf Conf

					err := unmarshalFile(cf, &conf)
					So(err, ShouldBeNil)
					So(conf, ShouldResemble, simple)
				})

				Convey("JSON", func() {
					cf := &configFile{
						fileName: "_testdata/simple.json",
						fileType: unknownType,
					}

					var conf Conf

					err := unmarshalFile(cf, &conf)
					So(err, ShouldBeNil)
					So(conf, ShouldResemble, simple)
				})
			})

			Convey("But fail with bad data", func() {
				cf := &configFile{
					fileName: "_testdata/bad-wolf",
					fileType: unknownType,
				}

				var conf Conf

				err := unmarshalFile(cf, &conf)
				So(err.Error(), ShouldStartWith, "couldn't read _testdata/bad-wolf as YAML/JSON")
			})
		})
	})

	Convey("Unmarshallers should error on bad file names", t, func() {
		var conf Conf

		Convey("YAML", func() {
			err := unmarshalYAML(&configFile{fileName: "bad!wolf"}, &conf)
			So(err, ShouldNotBeNil)
		})

		Convey("JSON", func() {
			err := unmarshalJSON(&configFile{fileName: "bad!wolf"}, &conf)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("When loading multiple files", t, func() {
		snoot := New(&Conf{})

		Convey("It should work with no files", func() {
			var conf Conf
			err := snoot.unmarshalFiles(&conf)
			So(err, ShouldBeNil)
			So(conf, ShouldResemble, Conf{})
		})

		Convey("It should work with one file", func() {
			var conf Conf
			snoot.MaybeAddFile("_testdata/simple.yaml")
			err := snoot.unmarshalFiles(&conf)
			So(err, ShouldBeNil)
			So(conf, ShouldResemble, simple)
		})

		Convey("It should override with multiple files", func() {
			type Override struct {
				A int
				B string
				E int
				D bool
			}
			var override Override
			snoot.MaybeAddFile("_testdata/overridden.yaml")
			snoot.MaybeAddFile("_testdata/override.yaml")
			err := snoot.unmarshalFiles(&override)
			So(err, ShouldBeNil)
			So(override, ShouldResemble, Override{
				// from both, equal
				A: 1,
				// from both, overridden
				B: "b",
				// from override, added
				E: 42,
				// from overridden, kept
				D: true,
			})
		})

		Convey("It should surface errors", func() {
			var conf Conf
			snoot.MaybeAddFile("_testdata/bad-wolf")
			err := snoot.unmarshalFiles(&conf)
			So(err.Error(), ShouldStartWith, "couldn't read _testdata/bad-wolf as YAML/JSON")
		})
	})
}
