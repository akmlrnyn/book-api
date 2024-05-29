package authcontroller

import (
	"encoding/json"
	"go-api-native/config"
	"go-api-native/helper"
	"go-api-native/models"
	"net/http"
)


func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil{
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	hashedPassword, err := helper.HashPassword(register.Password)

	if err != nil{
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	author := models.Author{
		Name: register.Name,
		Email: register.Email,
		Gender: register.Gender,
		Age: register.Age,
		Password: hashedPassword,
	}

	if err := config.DB.Create(&author).Error; err != nil{
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Registered Successfully", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil{
		helper.Response(w, 500, err.Error(), nil)
	}

	var author models.Author

	if err := config.DB.First(&author, "email = ?", login.Email).Error; err != nil{
		helper.Response(w, 500, err.Error(), nil)
	}

	if err := helper.VerifyPassword(author.Password, login.Password); err != nil{
		helper.Response(w, 404, err.Error(), nil)
		return
	}

	token, err := helper.GenerateToken(&author)

	if err != nil{
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "sucessfully login", token)

	
}