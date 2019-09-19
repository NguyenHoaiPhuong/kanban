package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/NguyenHoaiPhuong/kanban/server/api"
	"github.com/NguyenHoaiPhuong/kanban/server/config"
	"github.com/NguyenHoaiPhuong/kanban/server/repo"
)

// App struct
type App struct {
	cfg  *config.Config
	apis *api.APIs
	mdb  *repo.MongoDB
}

// Init : initialize settings
func (a *App) Init() {
	a.initConfig()
	a.initAPIs()
	a.initRepo()
}

func (a *App) initConfig() {
	log.Println("Init config:")
	a.cfg = config.SetupConfig("./resources/config-dev.json")

	log.Println("Host:", *a.cfg.Host)
	log.Println("Port:", *a.cfg.Port)
	log.Println("Database name:", *a.cfg.DBName)
}

func (a *App) initAPIs() {
	log.Println("Initialize APIs")
	a.apis = new(api.APIs)
	a.apis.Init()
}

func (a *App) initRepo() {
	log.Println("Initialize MongoDB")
	a.mdb = new(repo.MongoDB)
	a.mdb.Init(*a.cfg.Host, *a.cfg.Port, *a.cfg.DBName)
}

// Run server
func (a *App) Run() {
	log.Println("Run the app on port 9001")

	srv := &http.Server{
		Handler:      a.apis.Root,
		Addr:         ":9001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

	defer a.mdb.Client.Disconnect(context.Background())
}
