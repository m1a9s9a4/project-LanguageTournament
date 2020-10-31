package main

import (
	"github.com/m1a9s9a4/language_tournament/route"
)

func main() {
	router := route.Init()
	router.Start(":8080")
}
