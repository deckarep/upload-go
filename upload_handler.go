package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/upload", fileHandler)
	http.ListenAndServe(":9292", nil)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {

	f, _, err := r.FormFile("image")

	if err != nil {
		log.Fatal("Can't Find Image ")
	}

	log.Println("Received image!")

	t, _ := ioutil.TempFile(".", "image-")
	defer t.Close()

	_, err = io.Copy(t, f) // copy the image

	if err != nil {
		log.Fatal("Somthing went wrong")
	}

	log.Println("Image saved to: ", t.Name())
}
