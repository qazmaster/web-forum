package main

import (
	"flag"
	"log"
)

var (
	configPath      string
	errorConfigInit = "Configuration initialization failed."
)

func initMain() {
	log.Println("Starting initilazation..")
	flag.StringVar(&configPath, "config-path", "configs/config.yaml", "path to config file")
	flag.Parse()
	err := initConfig(configPath)
	if err != nil {
		log.Printf("wtf..\n%v\n", err)
		log.Fatalf("oops..\n%v\n", errorConfigInit)
	}
	log.Printf("Got the configuration\n%v\n", config)
}

func main() {
	initMain()

}
