package main

import (
	"html/template"
	"log"
	"net/http"
)

// サーバ

// type MyHnadler struct{}

// func (h *MyHnadler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// }

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "Hello world111!")
}

func main() {
	http.HandleFunc("/top", top)
	http.ListenAndServe(":8080", nil)
}
