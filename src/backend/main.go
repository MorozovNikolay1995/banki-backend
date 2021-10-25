package main

import (
	"banki/dbconfig"
	"fmt"
)

func main() {
	dbconf, err := dbconfig.GetDBConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbconf.DB_HOST)
}
