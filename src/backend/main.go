package main

import (
	"banki/dbconfig"
	"banki/routes"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := dbconfig.SetDBConfig()
	if err != nil {
		fmt.Println(err)
	}
	router := gin.Default()
	router.GET("/api/sample/", routes.SampleHandler)
	router.GET("/api/export/", routes.ExportHandler)
	router.GET("/api/stats/", routes.StatsHandler)

	router.Run("localhost:" + os.Getenv("PORT"))
}
