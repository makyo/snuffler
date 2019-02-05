package snuffler

import (
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

// _fileType represents what type a given file is. So far, it knows about YAML,
// TOML, and JSON types.
type _fileType int

const (
	unknownType _fileType = iota
	yamlType
	tomlType
	jsonType
)

// configFile represents a file, along with its name and file type.
type configFile struct {
	fileName string
	fileType _fileType
}

// addFile checks if the file with the given path exists. If so, it adds it
// to the list of config files regardless of type. If not, it returns a
// FileNotFoundError.
func (s *Snuffler) addFile(p string) error {
	if _, err := os.Stat(p); err != nil {
		return err
	}
	for _, cf := range s.files {
		if p == cf.fileName {
			return nil
		}
	}
	cf := &configFile{
		fileName: p,
	}
	if isYAML(p) {
		cf.fileType = yamlType
	} else if isTOML(p) {
		cf.fileType = tomlType
	} else if isJSON(p) {
		cf.fileType = jsonType
	} else {
		cf.fileType = unknownType
	}
	s.files = append(s.files, cf)
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

// isJSON attempts to guess in a very stupid fashion whether or not a file is
// a JSON file.
func isJSON(p string) bool {
	return p[len(p)-4:] == "json"
}
