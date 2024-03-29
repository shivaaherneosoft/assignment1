package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/shivaaherneosoft/assignment1/api/models"
	"github.com/shivaaherneosoft/assignment1/config"
)

// var users = map[string]string{
// 	"user1": "password1",
// 	"user2": "password2",
// }

func Signin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := config.CONFIG.Users[creds.Username]["password"]

	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	roleid, ok := config.CONFIG.Users[creds.Username]["roleid"]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Println("roleid:", roleid)

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &models.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Id:        roleid,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.CONFIG.JWTKey)
	if err != nil {
		fmt.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authtoken := models.AuthResponse{Token: tokenString}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authtoken)

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })
}

func Refresh(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	authtoken := r.Header.Get("Authorization")
	if authtoken == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(authtoken, claims, func(token *jwt.Token) (interface{}, error) {
		return config.CONFIG.JWTKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.CONFIG.JWTKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := models.AuthResponse{Token: tokenString}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
