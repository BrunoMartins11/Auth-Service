package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

func signInHandler(w http.ResponseWriter, req *http.Request) {
	var user User

	error := json.NewDecoder(req.Body).Decode(&user)
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	//TODO check if user exists in DB

	expiresAt := time.Now().Add(time.Minute * 10000).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	//TODO save (token, expiration) in a database

	sendToken(w, tokenString, expiresAt)

}

func sendToken(w http.ResponseWriter, token string, expiresAt int64) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UserToken{
		Token:     token,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	})
}
