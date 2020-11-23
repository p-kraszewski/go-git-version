package main

import (
	"flag"
	"log"
	"strings"

	"github.com/p-kraszewski/go-git-version/lib"
)

var (
	confName   = flag.String("conf", ".go-git-version.toml", "Config file path")
	dir        = flag.String("dir", ".", "Directory to search data in")
	genDefConf = flag.Bool("gendefconf", false, "Generate default config in current directory")
	lang       = flag.String("lang", "go", "Generate for a specific language, use ? to list supported languages")
	verbose    = flag.Bool("verbose", false, "Print additional information")
	mod        = flag.String("module", "", "Set module name if not default for language")
)

func main() {
	flag.Parse()

	if *lang == "?" {
		langList := lib.SupportedLanguages()
		langStr := strings.Join(langList, ", ")
		log.Println("Supported languages:", langStr)
		return
	}

	if *genDefConf {
		c := lib.DefConfig()
		err := c.WriteConfig(*confName)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	c, from := lib.SearchConfig(*confName)
	if *verbose {
		log.Printf("Config: %+v", c)
		log.Printf("Loaded from: %s", from)
	}

	i, err := lib.ParseGit(*dir)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = lib.Generate(*lang, c, i)
	if err != nil {
		log.Fatalln(err)
	}
}
