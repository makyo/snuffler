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
	unknownType _fileType = iota
	yamlType
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

// addFile checks if the file with the given path exists. If so, it adds it
// to the list of config files regardless of type. If not, it returns a
// FileNotFoundError.
func (s *Snuffler) addFile(p string) error {
	if _, err := os.Stat(p); err != nil {
		return err
	}
	cf := &confFile{
		fileName: p,
	}
	if isYAML(p) {
		cf.fileType = yamlType
	} else if isTOML(p) {
		cf.fileType = tomlType
	} else {
		cf.fileType = unknownType
	}
	s.files = append(s.files, cf)
	return nil
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

// isYAML attempts to guess in a very stupid fashion whether or not a file is
// a YAML file.
func isYAML(p string) bool {
	return p[len(p)-4:] == "yaml" || p[len(p)-3:] == "yml"
}

// isTOML attempts to guess in a very stupid fashion whether or not a file is
// a TOML file.
func isTOML(p string) bool {
	return p[len(p)-4:] == "toml"
}
