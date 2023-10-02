package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

const SECRET string = "1235"
const APIKEY string = "5555"

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized: token is nil"))
		}
		token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthorized: cannot parse token"))
			}
			return []byte(SECRET), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized" + err.Error()))
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized: token un valid"))
		}
		next(w, r)
	})
}

func GetJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] == nil {
		return
	}
	if r.Header["Access"][0] != APIKEY {
		return
	}
	token, err := CreateJWT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot create jwt"))
	}
	fmt.Fprint(w, token)
}
