package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/arsu4ka/go_rest_api/internal/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf", "./configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
