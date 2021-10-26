package main

import (
	"banki/routes"
	"net/http"
)

func main() {
	http.HandleFunc("/api/sample/", routes.SampleHandler)
	http.HandleFunc("/api/export/", routes.ExportHandler)
	http.HandleFunc("/api/stats/", routes.StatsHandler)

	http.ListenAndServe(":8000", nil)
}
