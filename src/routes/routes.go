package routes

import (
	"fmt"
	"net/http"
)

func SampleHandler(http.ResponseWriter, *http.Request) {
	// "select id, link, title, city, bank_name, score, status, username, create_dt, comments from home.dt_banki_responses order by id desc limit 10"
	// {"status": "success", "data": result, "time": datetime.now()}
	fmt.Println("sample")
}

func ExportHandler(http.ResponseWriter, *http.Request) {
	// "select *from home.dt_banki_responses where create_dt <= date(now())-2 order by id desc limit 100"
	// response["Content-Disposition"] = f"attachment; filename=export_{sysdate()}.csv"
	fmt.Println("Export")
}

func StatsHandler(http.ResponseWriter, *http.Request) {
	// "select * from home.v_stats"
	// {"status": "success", "data": result, "time": datetime.now()}
	fmt.Println("stats")
}
