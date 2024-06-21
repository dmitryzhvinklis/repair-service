package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/dmitryzhvinklis/carfix_dev/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}