package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.StaticFS("/static", StaticFS)

	r.GET("/", indexHandler)

	r.Run(":7777")
}
