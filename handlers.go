package main

import "github.com/gin-gonic/gin"

func indexHandler(c *gin.Context) {
	Templates.ExecuteTemplate(c.Writer, "index.html", gin.H{})
}
