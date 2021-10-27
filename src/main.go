package main

import (
	"banki/routes"
	"net/http"
	"os"
)

func getServerString() string {
	ip:=os.Getenv("SERVER_IP")
	port:=os.Getenv("SERVER_PORT")
	if(port=="" || ip==""){
		panic("Incorrect server settings. Please, set environment variables: SERVER_IP, SERVER_PORT")
	}
	return ip+":"+port
}

func main() {
	http.HandleFunc("/api/sample/", routes.SampleHandler)
	http.HandleFunc("/api/export/", routes.ExportHandler)
	http.HandleFunc("/api/stats/", routes.StatsHandler)

	http.ListenAndServe(getServerString(), nil)
}
