package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/otiai10/gosseract"
)

func ocrImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ocrImage")
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("./test/" + handler.Filename)
	text, _ := client.Text()
	fmt.Println(text)
	err = os.Remove("./test/" + handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	w.Write([]byte(text))

}

func handleRequests() {
	http.HandleFunc("/ocrImage", ocrImage)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequests()
}
