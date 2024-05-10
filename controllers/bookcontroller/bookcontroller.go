package bookcontroller

import (
	"encoding/json"
	"errors"
	"go-api-native/config"
	"go-api-native/helper"
	"go-api-native/models"
	"net/http"

	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var book []models.Book
	if err := config.DB.Find(&book).Error; err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Book List", book)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var author models.Author

	if err := config.DB.First(&author,book.AuthorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success to create book", nil)
}