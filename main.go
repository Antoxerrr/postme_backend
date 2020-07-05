package main

import (
	"net/http"
	_config "postme/config"
	"postme/controllers"
	"postme/database"
	"postme/model"
)

func main() {
	config := _config.GetConfig()

	db := database.InitDB(config)
	defer db.Close()
	model.MigrateModels(db)

	r := controllers.InitRouter()
	http.ListenAndServe(":3000", r)
}
