package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("it works"))
	})
	log.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
