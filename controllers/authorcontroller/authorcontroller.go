package authorcontroller

import (
	"encoding/json"
	"go-api-native/config"
	"go-api-native/helper"
	"go-api-native/models"
	"net/http"
)


func Index(w http.ResponseWriter, r *http.Request) {
	var author []models.Author

	if err := config.DB.Find(&author).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List author", author)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&author).Error; err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Successs!!", nil)

}