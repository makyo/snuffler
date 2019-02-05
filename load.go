package snuffler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	yaml "gopkg.in/yaml.v2"
)

// unmarshalFiles loads each file in turn into the interface, starting with
// the first files added to the snuffler.
func (s *Snuffler) unmarshalFiles(conf interface{}) error {
	for _, f := range s.files {
		if err := unmarshalFile(f, conf); err != nil {
			return err
		}
	}
	return nil
}

// unmarshalFile loads the contents of one file into the provided interface.
func unmarshalFile(f *configFile, conf interface{}) error {
	if f.fileType == yamlType {
		return unmarshalYAML(f, conf)
	} else if f.fileType == tomlType {
		return unmarshalTOML(f, conf)
	} else if f.fileType == jsonType {
		return unmarshalJSON(f, conf)
	} else {
		// Since JSON is valid YAML, we don't need to try that one.
		if errYAML := unmarshalYAML(f, conf); errYAML != nil {
			if errTOML := unmarshalTOML(f, conf); errTOML != nil {
				return fmt.Errorf("couldn't read %s as YAML/JSON (%v) or TOML (%v)", f.fileName, errYAML, errTOML)
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
}

func unmarshalYAML(f *configFile, conf interface{}) error {
	contents, err := ioutil.ReadFile(f.fileName)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(contents, conf)
}

func unmarshalTOML(f *configFile, conf interface{}) error {
	_, err := toml.DecodeFile(f.fileName, conf)
	return err
}

func unmarshalJSON(f *configFile, conf interface{}) error {
	contents, err := ioutil.ReadFile(f.fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(contents, conf)
}
