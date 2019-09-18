package app

import (
	"fmt"
	"log"

	"github.com/NguyenHoaiPhuong/kanban/server/api"
	"github.com/NguyenHoaiPhuong/kanban/server/config"
)

// App struct
type App struct {
	cfg  *config.Config
	apis *api.APIs
}

// Init : initialize settings
func (a *App) Init() {
	a.initConfig()
}

func (a *App) initConfig() {
	log.Println("Init config:")
	a.cfg = config.SetupConfig("./resources/config-dev.json")

	fmt.Println("Host:", *a.cfg.MongoDBConfig.Host)
	fmt.Println("Port:", *a.cfg.MongoDBConfig.Port)
	fmt.Println("Database name:", *a.cfg.MongoDBConfig.DBName)
}

func (a *App) initAPIs() {
	log.Println("Initialize APIs")
	a.apis = new(api.APIs)
	a.apis.Init()
}

// Run server
func (a *App) Run() {
}
