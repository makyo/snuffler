package snuffler

import (
	"errors"
	"os"
)

// _patternType represents what type of pattern each filePattern is. It can be
// a path or a glob.
type _patternType int

const (
	filePath _patternType = iota
	fileGlob
)

// filePattern associates a provided pattern with what type it is.
type filePattern struct {
	pattern     string
	patternType _patternType
}

// _fileType represents what type a given file is. So far, it knows about YAML
// and TOML types.
type _fileType int

const (
	yamlType _fileType = iota
	tomlType
)

// confFile represents a file, along with its name and file type.
type confFile struct {
	fileName string
	fileType _fileType
	file     *os.File
}

// open opens the config file.
func (c *confFile) open() error {
	return errors.New("not implemented")
}

// read reads the config file in its entirety and returns a byte slice to be
// used in unmarshalling.
func (c *confFile) read() ([]byte, error) {
	return []byte{}, errors.New("not implemented")
}

// openFiles opens each file in the fileMatches attribute to populate the files
// attribute.
func (s *Snuffler) openFiles() error {
	for _, cf := range s.files {
		if err := cf.open(); err != nil {
			return err
		}
	}
	return nil
}

// expandGlob finds all matching files in the pattern and returns confFiles it
// finds.
func (s *Snuffler) expandGlob(g string) ([]*confFile, error) {
	return []*confFile{}, errors.New("not implemented")
}

// expandPath gets the absolute path of a filePath pattern and returns the
// associated confFile.
func (s *Snuffler) expandPath(p string) (*confFile, error) {
	return nil, errors.New("not implemented")
}

// GetFileMatchList returns the list of file names matched by the list of
// patterns the snuffler knows about.
func (s *Snuffler) GetFileMatchList() []string {
	matchList := make([]string, len(s.files))
	for i, cf := range s.files {
		matchList[i] = cf.fileName
	}
	return matchList
}

// GetFileList returns all of the files the snuffler knows about.
func (s *Snuffler) GetFileList() []*os.File {
	fileList := make([]*os.File, len(s.files))
	for i, cf := range s.files {
		fileList[i] = cf.file
	}
	return fileList
}
