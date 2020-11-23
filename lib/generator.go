package lib

import (
	"errors"
	"sort"
)

type Gen interface {
	Generate(module string, c *Config, i *Info) error
}

var (
	backends = make(map[string]Gen)
)

func Generate(lang, mod string, c *Config, i *Info) error {
	if gen, found := backends[lang]; found {
		return gen.Generate(mod, c, i)
	} else {
		return errors.New("Unknown backend " + lang)
	}
}

func SupportedLanguages() []string {
	ll := []string{}
	for l := range backends {
		ll = append(ll, l)
	}
	sort.Strings(ll)
	return ll
}
