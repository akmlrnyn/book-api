package middleware

import (
	"context"
	"go-api-native/helper"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//get the token from postman's header
       accessToken := r.Header.Get("Authorization")

	   //if the token is empty string
	   if accessToken == "" {
		helper.Response(w, 404, "Unathorized", nil)
		return
	   }

	   //validate the token
	   user, err := helper.ValidateToken(accessToken)

	   if err != nil {
		helper.Response(w, 404, err.Error(), nil)
		return
	   }

	   ctx := context.WithValue(r.Context(), "userinfo", user)

    	next.ServeHTTP(w, r.WithContext(ctx))
    })
}