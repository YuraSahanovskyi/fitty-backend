package main

import (
	"log"
	"net/http"
)

func main() {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("it works"))
	})
	http.Handle("/test", fn)

	log.Println("server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
