package views

import (
	"net/http"
	"postme/database"
	"postme/model"
	"postme/utils"
)

type Message struct {
	Text string `json:"text"`
}

func TestGetView(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	var tests []model.Test
	result := db.Find(&tests)
	utils.RespondJSON(w, 200, result)
}

func TestPostView(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := utils.ParseBody(r, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	db := database.GetDB()
	db.Create(&model.Test{Text: msg.Text})
	utils.RespondJSON(w, 200, map[string]string{"message": "ok"})
}
