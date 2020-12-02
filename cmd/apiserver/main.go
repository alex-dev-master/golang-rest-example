package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/alex-dev-master/golang-rest-example/internal/app/server"
	"log"
)


var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main()  {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s", "world")
}