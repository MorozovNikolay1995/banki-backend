package main

import (
	"banki/dbconfig"
	"banki/routes"

	"fmt"
	"net/http"
	"os"
)

func main() {
	err := dbconfig.SetDBConfig()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/api/sample/", routes.SampleHandler)
	http.HandleFunc("/api/export/", routes.ExportHandler)
	http.HandleFunc("/api/stats/", routes.StatsHandler)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
