package helper

import (
	"go-api-native/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigningKey = []byte("myToken")

type MyCustomClaims struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Gender string `json:"gender"`
	Age int `json:"age"`
	jwt.RegisteredClaims
}

func GenerateToken(author *models.Author) (string, error){
claims := MyCustomClaims{
	int(author.ID),
	author.Name,
	author.Email,
	author.Password,
	author.Gender,
	author.Age,
	jwt.RegisteredClaims{
		
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	},
}

token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
ss, err := token.SignedString(mySigningKey)

return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error){
		return mySigningKey, nil
	})

	claims, ok := token.Claims.(*MyCustomClaims)

	if err != nil{
		return nil, err
	}

	if !ok || !token.Valid{
		return nil, err
	}

	return claims, nil
}
