package main

import (
	"errors"
)

// unmarshalFile loads the contents of one file into the provided interface.
func (s *snuffler) unmarshalFile(f *confFile) error {
	return errors.New("not implemented")
}

// unmarshalFiles loads each file in turn into the interface, starting with
// the first files added to the snuffler.
func (s *snuffler) unmarshalFiles() error {
	return errors.New("not implemented")
}
