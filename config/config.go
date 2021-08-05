package config

import (
	yaml "gopkg.in/yaml.v3"
)

type YamlUnmarshaler struct{}

func (y *YamlUnmarshaler) Unmarshal(data []byte, val interface{}) error {
	return yaml.Unmarshal(data, val)
}
