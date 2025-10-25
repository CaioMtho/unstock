package main

import (
	"github.com/gin-gonic/gin"
	"github.com/CaioMtho/unstock/internal/config"
	"github.com/CaioMtho/unstock/api"
)

func main() {
	config.InitDB()
	router := api.SetupEndpoints()
	router.Run(":8080")
}