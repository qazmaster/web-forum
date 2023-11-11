package main

import (
	"database/sql"
	"fmt"
	"forum-authentication/internal/content"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

func initApp() {
	log.Println("Starting application..")
	app, err := newApp(config)
	if err != nil {
		log.Printf("%v\n", err)
		log.Fatalf("oops..\n%v\n", errAppCreating)
	}
	log.Printf("Application created. Server addr is \n%v\n", app.server.httpServer.Addr)
	err = app.Run()
	if err != nil {
		log.Printf("%v\n", err)
		log.Fatalf("oops..\n%q\n", errAppLaunch)
	}
	log.Println("Application launched")
	defer log.Println("Application closed")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("db close failed")
		}
	}(app.db.db)
}
func newApp(cfg *Config) (*Application, error) {
	var a *Application
	db, err := openDB(cfg.Database.DSN)
	if err != nil {
		return nil, err
	}
	repo := content.NewSqlRepo(db)
	cases := content.NewCases(*repo)
	mux := content.NewHTTPController(cases)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           mux,
		MaxHeaderBytes:    1 << 20,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	onceApp.Do(func() {
		a = &Application{
			server: &Server{
				httpServer: httpServer,
			}, db: &Database{
				db: db,
			},
		}
	})
	return a, nil
}

func (a *Application) Run() error {
	err := a.server.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func openDB(dsn string) (*sql.DB, error) {
	log.Println("Opening db...")
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	log.Println("db is opened")
	err = db.Ping()
	log.Println("Connecting to the db...")
	if err != nil {
		return nil, err
	}
	log.Println("db is connected")
	return db, nil
}
