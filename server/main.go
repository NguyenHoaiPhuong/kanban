package main

import "github.com/NguyenHoaiPhuong/kanban/server/app"

func main() {
	a := new(app.App)
	a.Init()
	a.Run()
}
