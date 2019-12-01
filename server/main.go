package main

import (
	"github.com/NguyenHoaiPhuong/warehouse/server/app"
	"github.com/NguyenHoaiPhuong/warehouse/server/log"
)

func init() {
	// IMPORTANT: to make a config for logging level
	log.SetSTDHook(5)
}

func main() {
	a := new(app.App)
	a.Init()
	a.Run()
}
