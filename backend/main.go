package main

import (
	"github.com/m1a9s9a4/language_tournament/conf"
	"github.com/m1a9s9a4/language_tournament/route"
)

func main() {
	conf.SetEnv()
	router := route.Init()
	router.Start(":8080")
}
