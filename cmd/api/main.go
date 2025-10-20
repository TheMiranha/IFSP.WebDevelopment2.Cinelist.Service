package main

import (
	infrastructure_database "cinelist/infrastructure/database"
	infrastructure_http "cinelist/infrastructure/http"
)

func main() {
	database := infrastructure_database.ConnectDB()
	infrastructure_http.InitializeServer(database)
}
