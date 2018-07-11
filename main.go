package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nicholasflee/nick-b/routers"
)

func main() {
	testFeature()
	r := gin.Default()
	routers.Routes(r)
	r.Run(":1234")
}

func testFeature() {

}
