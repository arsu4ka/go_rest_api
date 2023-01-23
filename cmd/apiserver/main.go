package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/arsu4ka/go_rest_api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf", "./configs/apiserver.toml", "path to config file")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	check(err)

	s := apiserver.New(config)
	err = s.Start()
	check(err)
}
