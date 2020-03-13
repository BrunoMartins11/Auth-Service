package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

type UserClaim struct {
	User
	*jwt.StandardClaims
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

	claim := &UserClaim{
		user,
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

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

func validateToken(w http.ResponseWriter, req *http.Request) {
	bearerToken := req.Header.Get("Authorization")
	strtoks := strings.Split(bearerToken, " ")

	if len(strtoks) != 2 {
		w.WriteHeader(http.StatusBadRequest)
	}

	tokenString := strtoks[1]
	claim := &UserClaim{}

	token, error := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})

	if error != nil {
		if error == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
