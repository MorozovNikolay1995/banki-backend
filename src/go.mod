module banki/backend

go 1.17

replace banki/routes => ./routes

require banki/routes v0.0.0-00010101000000-000000000000

require (
	github.com/joho/sqltocsv v0.0.0-20210428211105-a6d6801d59df // indirect
	github.com/lib/pq v1.10.3 // indirect
)
