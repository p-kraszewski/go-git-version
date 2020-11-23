package lib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type Config struct {
	VerFileNameBase string
	VerVarBase      string
	DateFormat      string
	CommitIdLength  uint8
}

func DefConfig() *Config {
	return &Config{
		VerFileNameBase: "VERSION",
		VerVarBase:      "VER",
		DateFormat:      "2006-01-02 15:04:05",
		CommitIdLength:  8,
	}
}

func SearchConfig(name string) (*Config, string) {

	log.Println(path.Dir(name))
	if path.Base(name) != name {
		c := DefConfig()
		_, err := toml.DecodeFile(name, &c)
		if err == nil {
			return c, name
		} else {
			return DefConfig(), ""
		}
	}

	pwd, err := os.Getwd()
	if err != nil {
		return DefConfig(), ""
	}

	for {
		src := path.Join(pwd, name)
		c := DefConfig()
		_, err := toml.DecodeFile(src, &c)
		if err == nil {
			return c, src
		}

		if pwd == "" || pwd == "/" {
			return DefConfig(), ""
		}

		pwd = path.Dir(pwd)
	}
}

func (c *Config) WriteConfig(name string) error {
	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(c); err != nil {
		return fmt.Errorf("TOML %w", err)
	}

	if err := ioutil.WriteFile(name, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("file %w", err)
	}

	return nil
}
