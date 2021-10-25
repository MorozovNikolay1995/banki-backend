package main

import (
	"banki/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/sample/", routes.SampleHandler)
	router.GET("/api/export/", routes.ExportHandler)
	router.GET("/api/stats/", routes.StatsHandler)

	router.Run("localhost:8000")
}
