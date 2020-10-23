package main

import (
	"github.com/labstack/echo"
	"language_tournament/route"
)

func main() {
	router := route.Init()	
	router.Run(fasthttp.New(":8888"))
}
