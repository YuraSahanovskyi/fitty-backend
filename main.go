package main

import (
	"github.com/YuraSahanovskyi/fitty-backend/auth"
	"log"
	"net/http"
)

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	_ = r
	_, _ = w.Write([]byte("it works"))
}

func main() {
	http.Handle("/test", auth.ValidateJWT(ProtectedEndpoint))

	http.HandleFunc("/jwt", auth.GetJWT)

	log.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
