package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	//http.ListenAndServe(":8083", nil)
	log.Fatal(http.ListenAndServe(":8084", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {

	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
		return
	case <-time.After(5 * time.Second):
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso"))
	}
}
