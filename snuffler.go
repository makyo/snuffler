package snuffler

import (
	"errors"
	"path/filepath"
)

// Snuffler is the primary object holding all of the bits required to snuffle
// through config files.
type Snuffler struct {
	conf         *interface{}
	filePatterns []filePattern
	files        []*confFile
}

// AddFile accepts a string containing a filename and adds it to the list of
// files that the snuffler should load. This file must exist. If you want to
// add a file that may or may not exist, use MaybeAddFile.
func (s *Snuffler) AddFile(p string) error {
	return s.addFile(p)
}

// MaybeAddFile accepts a string containing a filename and adds it to the list
// of files that the snuffler should load. This file need not exist. If you
// want to add a file that must exist, use AddFile.
func (s *Snuffler) MaybeAddFile(p string) {
	s.addFile(p)
}

// AddGlob accepts a string containing a Glob[0] to search for config files.
//
// [0]: https://golang.org/pkg/path/filepath/#Glob
func (s *Snuffler) AddGlob(g string) error {
	matches, err := filepath.Glob(g)
	if err != nil {
		return err
	}
	for _, match := range matches {
		s.addFile(match)
	}
	return nil
}

// Snuffle performs the noble task of snuffling through all of the specified
// config files and paths to populate the provided config object. Files are
// loaded in the order they were received, and values are overwritten if
// subsequent files specify them.
func (s *Snuffler) Snuffle() error {
	return errors.New("not implemented")
}

// Snorfle performs all the same tasks as Snuffle, but does not use the stored
// config object reference, instead populating the one that is provided as an
// argument.
func (s *Snuffler) Snorfle(cfg *interface{}) error {
	return errors.New("not implemented")
}

// New creates a new snuffler object with the given interface. You can then
// add files to the resulting snuffler and, when run, it will load the config
// from each of them into the interface.
func New(c *interface{}) *Snuffler {
	return &Snuffler{
		conf: c,
	}
}
