package main

import (
	"net/http"

	"github.com/alexcoder04/rpi-homescreen/data"
	"github.com/gin-gonic/gin"
)

func TimeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, data.Time())
}

func SchoolDaysHandler(c *gin.Context) {
	d, err := data.DayCounter()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, d)
}

func SysstatHandler(c *gin.Context) {
	c.JSON(http.StatusOK, data.Sysstat())
}
