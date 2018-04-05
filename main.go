package main

import (
	"github.com/go-martini/martini"
	"cloudControlSystem/config"
	"cloudControlSystem/controllers"
)

func main() {
	m := martini.Classic()
	m.Get("/api/v1/install", controllers.InstallScheme)
	m.RunOnAddr(config.PORT)
}
