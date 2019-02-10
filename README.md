[![Build Status](https://travis-ci.org/makyo/snuffler.svg?branch=master)](https://travis-ci.org/makyo/snuffler)
[![GoDoc reference](https://godoc.org/github.com/makyo/snuffler?status.svg)](https://godoc.org/github.com/makyo/snuffler)
[![Support me on Patreon](https://img.shields.io/badge/patreon-support-%23222222.svg)](https://patreon.com/makyo)
[![Buy me a Coffee](https://img.shields.io/badge/kofi-support-%23222222.svg)](https://ko-fi.com/drabmakyo)
[![Support me on LiberaPay](https://img.shields.io/badge/liberapay-support-%23222222.svg)](https://liberapay.com/makyo)

# Snuffler

Snuffler will snuffle through all of the paths and globs you provide to look for config files (YAML, TOML, and JSON, for the moment), and use those to populate the config object you provide. It will clobber existing keys, but that's often what you want when generating user specific config which has the ability to override global config. Simply provide the paths/glob where config might live and it will root through them in order for config files.

For example:

```go
package main
import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/makyo/snuffler"
)

func main() {
	// Get the user's home directory.
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	// Somewhere else, you've created a YAML/TOML/JSON ready config struct...
	var cfg myConfig

	// Build a new snuffler with a pointer to your config object.
	s := snuffler.New(&cfg)

	// You can add a file by its absolute path. If it does not exist, you'll
	// get an error. Probably a sign that your app wasn't installed
	// correctly or the user can't read the file.
	if err := s.AddFile("/etc/myApp/master.yaml"); err != nil {
		panic(err)
	}

	// You can add a glob for a directory or series of files. It will just
	// grab all files, so be specific with your glob! An error will occur
	// only if your glob is malformed.
	if err := s.AddGlob("/etc/myApp/conf.d/*.yaml"); err != nil {
		panic(err)
	}
	if err := s.AddGlob(filepath.Join(home, ".config", "myApp", "*.[ty][oa]ml"); err != nil {
		panic(err)
	}

	// If a file might not exist, you can easily add it without worrying
	// about an error (and make it obvious to the readers of the code that
	// you're okay with that) by using MaybeAddFile.
	s.MaybeAddFile(filepath.Join(home, ".myApprc")

	// Snuffle will read in all of the files in the order specified above
	// and unmarshal them into the cfg object provided above.
	s.Snuffle()

	// If you need to snuffle config into a new object other than the one
	// you provided during construction, you can use Snorfle.
	var secondCfg myConfig
	s.Snorfle(&secondConfig)
}
```
