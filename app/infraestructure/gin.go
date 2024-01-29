package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pluvet/go-bank/app/primary_adapters"
)

func main() {
	r := gin.Default()
	DBConnect()
	primary_adapters.Routes(r)
	r.Run()
}
