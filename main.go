package main

import (
	"log"
	"net/http"
)

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works"))
}

func main() {
	http.Handle("/test", ValidateJWT(ProtectedEndpoint))

	http.HandleFunc("/jwt", GetJWT)

	log.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
