package main

import (
	"errors"
)

// snuffler is the primary object holding all of the bits required to snuffle
// through config files.
type snuffler struct {
	conf         *interface{}
	filePatterns []filePattern
	files        []*confFile
}

// AddFile accepts a string containing a filename and adds it to the list of
// files that the snuffler should load. This file must exist. If you want to
// add a file that may or may not exist, use MaybeAddFile.
func (s *snuffler) AddFile(p string) error {
	return errors.New("not implemented")
}

// AddFile accepts a string containing a filename and adds it to the list of
// files that the snuffler should load. This file need not exist. If you want
// to add a file that must exist, use AddFile.
func (s *snuffler) MaybeAddFile(p string) error {
	return errors.New("not implemented")
}

// AddGlob accepts a string containing a Glob[0] to search for config files.
// The fully specified portion of the path must exist (that is, for the glob
// `/path/to/files/*/*.yaml`, `/path/to/files` must exist). If you want to add
// a glob where that needn't be true, use MaybeAddGlob.
//
// [0]: https://golang.org/pkg/path/filepath/#Glob
func (s *snuffler) AddGlob(g string) error {
	return errors.New("not implemented")
}

// MaybeAddGlob accepts a string containing a Glob[0] to search for config
// files. The fully specified portion of the path need not exist (that is, for
// the glob `/path/to/files/*/*.yaml`, `/path/to/files` must exist). If you
// want to add a glob where the path must exist be true, use AddGlob.
//
// [0]: https://golang.org/pkg/path/filepath/#Glob
func (s *snuffler) MaybeAddGlob(g string) error {
	return errors.New("not implemented")
}

// Snuffle performs the noble task of snuffling through all of the specified
// config files and paths to populate the provided config object. Files are
// loaded in the order they were received, and values are overwritten if
// subsequent files specify them.
func (s *snuffler) Snuffle() error {
	return errors.New("not implemented")
}

// New creates a new snuffler object with the given interface. You can then
// add files to the resulting snuffler and, when run, it will load the config
// from each of them into the interface.
func New(c *interface{}) *snuffler {
	return &snuffler{
		conf: c,
	}
}
