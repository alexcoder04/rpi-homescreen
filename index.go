package main

import (
	"github.com/alexcoder04/rpi-homescreen/data"
	"github.com/gin-gonic/gin"
)

func MapMerge(maps ...map[string]any) map[string]any {
	res := map[string]any{}
	for _, m := range maps {
		for key, val := range m {
			res[key] = val
		}
	}
	return res
}

func indexHandler(c *gin.Context) {
	Templates.ExecuteTemplate(c.Writer, "index.html", MapMerge(data.Sysstat(), data.Time()))
}
