package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"sync"
)

var (
	configPath string
	config     *Config
	//app              *Application
	errorConfigInit  = "Configuration initialization failed."
	errAppLaunch     = "Application launching failed"
	errAppCreating   = "Application creating failed"
	onceCfg, onceApp sync.Once
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server,flow"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database,flow"`
}

type Server struct {
	httpServer *http.Server
}

type Database struct {
	db *sql.DB
}

type Application struct {
	server *Server
	db     *Database
}

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
	initApp()
}
