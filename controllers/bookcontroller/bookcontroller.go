package bookcontroller

import (
	"encoding/json"
	"errors"
	"go-api-native/config"
	"go-api-native/helper"
	"go-api-native/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var book []models.Book
	var bookResponse []models.BookResponse

	if err := config.DB.Joins("Author").Find(&book).Find(&bookResponse).Error; err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Book List", bookResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var author models.Author

	//check author
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

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book
	var bookResponse models.BookResponse

	if err := config.DB.Joins("Author").First(&book, id).First(&bookResponse).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.Response(w, 404, "Book not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
	
	}

	helper.Response(w, 200, "Book found", bookResponse)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author not found", nil)
		}

		helper.Response(w, 500, "Book not found", nil)
	}

	var bookPayload models.Book
	if err := json.NewDecoder(r.Body).Decode(&bookPayload); err != nil {
		helper.Response(w, 401, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var author models.Author
	if bookPayload.AuthorID == 0 {
		if err := config.DB.First(author, &bookPayload.AuthorID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Author not found", nil)
				return
			}

			helper.Response(w, 500, err.Error(), nil)
			return
		}
		return
	}

	if err := config.DB.Where("id = ?", id).Updates(&book).Error; err != nil {
		helper.Response(w, 404, "Book not found", nil)
		return
	}


	helper.Response(w, 201, "Success", nil)
}