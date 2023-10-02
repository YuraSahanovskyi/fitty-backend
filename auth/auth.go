package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

// TEMP: need to move to env

const SECRET string = "1235"
const APIKEY string = "5555"

func CreateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			handleUnauthorized(w, "nil token")
			return
		}
		if len(r.Header["Token"]) == 0 {
			handleUnauthorized(w, "no tokens")
			return
		}

		token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("different sign in method")
			}
			return []byte(SECRET), nil
		})

		if err != nil {
			handleUnauthorized(w, err.Error())
			return
		}

		if !token.Valid {
			handleUnauthorized(w, "token not valid")
			return
		}

		next(w, r)
	})
}

func handleUnauthorized(w http.ResponseWriter, s string) {
	errorString := fmt.Sprintf("unauthorized: %v", s)
	w.WriteHeader(http.StatusUnauthorized)
	if _, err := w.Write([]byte(errorString)); err != nil {
		return
	}
}

func GetJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(r.Header["Access"]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Header["Access"][0] != APIKEY {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := CreateJWT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("cannot create jwt"))
	}
	_, err = fmt.Fprint(w, token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
