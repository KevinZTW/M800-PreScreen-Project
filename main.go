package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"m800/internal/pkg/route"
)

func main() {
	r := gin.Default()
	log.Default().Println("server runs on default port: 8080")
	route.SetUp(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
