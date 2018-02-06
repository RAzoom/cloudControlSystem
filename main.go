package main

import (
	"github.com/go-martini/martini"
	"cloudControlSystem/config"
	"cloudControlSystem/controllers"
)

func main() {
	m := martini.Classic()
	m.Post("/api/v1/install", controllers.ColdStartSystem)
	m.RunOnAddr(config.PORT)
}
