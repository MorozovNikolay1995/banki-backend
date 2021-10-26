package routes

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/sqltocsv"
	_ "github.com/lib/pq"
)

var dbinfo string

func init() {
	err := GetDBInfo()
	if err != nil {
		panic(err)
	}
}

func GetDBInfo() error {
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")

	if db_host == "" || db_name == "" || db_password == "" || db_port == "" || db_user == "" {
		return errors.New("One or more environment variables is not set: DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT")
	}
	dbinfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_password, db_name)

	return nil
}

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic(err)
	}

	return db
}

type SampleUnit struct {
	Id       *uint   `json:"id"`
	Link     *string `json:"link"`
	Title    *string `json:"title"`
	City     *string `json:"city"`
	BankName *string `json:"bank_name"`
	Score    *uint8  `json:"score"`
	Status   *string `json:"status"`
	Username *string `json:"username"`
	CreateDT *string `json:"create_dt"`
	Comments *uint   `json:"comments"`
}

type BankInfo struct {
	BankName *string  `json:"bank_name"`
	Scorex   *float32 `json:"scorex"`
	Median   *uint8   `json:"median"`
	Cnt      *uint    `json:"cnt"`
}

type JsonResponseSampleUnits struct {
	Status string       `json:"status"`
	Data   []SampleUnit `json:"data"`
	Time   string       `json:"time"`
}

type JsonResponseBankStatus struct {
	Status string     `json:"status"`
	Data   []BankInfo `json:"data"`
	Time   string     `json:"time"`
}

func SampleHandler(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	rows, err := db.Query("select id, link, title, city , bank_name, score, status, username, create_dt, comments from home.dt_banki_responses order by id desc limit 10")
	defer rows.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var sUnits []SampleUnit

	for rows.Next() {
		var unit SampleUnit

		err := rows.Scan(&unit.Id, &unit.Link, &unit.Title, &unit.City, &unit.BankName, &unit.Score, &unit.Status, &unit.Username, &unit.CreateDT, &unit.Comments)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sUnits = append(sUnits, unit)
	}

	var response = JsonResponseSampleUnits{Status: "success", Data: sUnits, Time: time.Now().String()}

	json.NewEncoder(w).Encode(response)
}

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	rows, err := db.Query("select *from home.dt_banki_responses where create_dt <= date(now())-2 order by id desc limit 100")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=\"report.csv\"")

	sqltocsv.Write(w, rows)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	rows, err := db.Query("select * from home.v_stats")
	defer rows.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var bankInfos []BankInfo

	for rows.Next() {
		var unit BankInfo

		err := rows.Scan(&unit.BankName, &unit.Scorex, &unit.Median, &unit.Cnt)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bankInfos = append(bankInfos, unit)
	}

	var response = JsonResponseBankStatus{Status: "success", Data: bankInfos, Time: time.Now().String()}

	json.NewEncoder(w).Encode(response)
}
