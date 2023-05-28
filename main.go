package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.StaticFS("/static", StaticFS)

	r.GET("/", indexHandler)

	r.GET("/api/day-counter", SchoolDaysHandler)
	r.GET("/api/sysstat", SysstatHandler)
	r.GET("/api/time", TimeHandler)

	r.Run(":7777")
}
