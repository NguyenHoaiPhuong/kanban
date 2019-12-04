package app

import (
	"context"
	"net/http"
	"time"

	"github.com/NguyenHoaiPhuong/kanban/server/api"
	"github.com/NguyenHoaiPhuong/kanban/server/config"
	"github.com/NguyenHoaiPhuong/kanban/server/log"
	"github.com/NguyenHoaiPhuong/kanban/server/repo"
)

// App struct
type App struct {
	cfg  *config.Configurations
	apis *api.APIs
	mdb  *repo.MongoDB
}

// Init : initialize settings
func (a *App) Init() {
	a.initConfig()
	a.initRepo()
	a.initAPIs()
}

func (a *App) initConfig() {
	log.Info("Init config:")
	a.cfg = config.SetupConfig("./resources", "config-dev.json")

	log.Info("Host:", a.cfg.MGOConfig.ServerHost)
	log.Info("Port:", a.cfg.MGOConfig.ServerPort)
	log.Info("Username:", a.cfg.MGOConfig.ServerUsername)
	log.Info("Password:", a.cfg.MGOConfig.ServerPassword)
	log.Info("Database name:", a.cfg.MGOConfig.DatabaseName)
}

func (a *App) initAPIs() {
	log.Info("Initialize APIs")
	a.apis = new(api.APIs)
	a.apis.Init()

	a.apis.User.RegisterHandleFunction("POST", "/login", a.authenticate)
	a.apis.User.RegisterHandleFunction("OPTIONS", "/login", a.enableCORS)
	// a.apis.User.RegisterHandleFunction("OPTIONS", "/", a.enableCORS)
}

func (a *App) initRepo() {
	log.Info("Initialize MongoDB")
	a.mdb = new(repo.MongoDB)
	a.mdb.Init(a.cfg.MGOConfig.ServerHost, a.cfg.MGOConfig.ServerPort,
		a.cfg.MGOConfig.ServerUsername, a.cfg.MGOConfig.ServerPassword,
		a.cfg.MGOConfig.DatabaseName)
}

// Run server
func (a *App) Run() {
	log.Info("Run the app on port 5000")

	srv := &http.Server{
		Handler:      a.apis.Root,
		Addr:         ":5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

	defer a.mdb.Client.Disconnect(context.Background())
}
