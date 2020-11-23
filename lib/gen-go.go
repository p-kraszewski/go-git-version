package lib

import (
	"errors"
)

type GenGo struct {
}

func init() {
	backends["go"] = &GenGo{}
}

func (g GenGo) Generate(mod string, c *Config, i *Info) error {
	return errors.New("Unimplemented backend")
}
