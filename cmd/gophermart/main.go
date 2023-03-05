package main

import (
	"log"

	"GopherMart2/internal/router"
)

func main() {
	rout := router.InitServer()
	err := rout.Router()

	if err != nil {
		log.Fatal("Router:", err)
	}
}
