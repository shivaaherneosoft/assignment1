package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/shivaaherneosoft/assignment1/api/models"
	"github.com/shivaaherneosoft/assignment1/config"
)

// func ProtectRequest(handleFunc httprouter.Handle) httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 		authtoken := r.Header.Get("Authorization")
// 		if authtoken != "" && authtoken == "shivaaher" {
// 			handleFunc(w, r, params)
// 		} else {
// 			fmt.Println("[ERROR]:Invalid authtoken")
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("invalid authorization token"))
// 		}
// 	}
// }

func ProtectRequest(handleFunc httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tknStr := c.Value

		claims := &models.Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
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

		//valid user request
		handleFunc(w, r, params)
	}
}
