package bookcontroller

import (
	"encoding/json"
	"go-api-native/config"
	"go-api-native/helper"
	"go-api-native/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var book []models.Book
	if err := config.DB.Find(&book).Error; err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Book List", book)
}

