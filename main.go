package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nicholasflee/nick-b/routers"
)

func main() {
	r := gin.Default()
	routers.Routes(r)
	r.Run(":80")
}
