package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() { //lire tout nos templates html
	tmpl := template.Must(template.ParseFiles("template/home.html", "template/level.html", "template/end.html", "template/menu.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home", nil)
	})

	http.HandleFunc("/level", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "level", nil)
	})

	http.HandleFunc("/end", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "end", nil)
	})

	http.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "menu", nil)

	})

	http.HandleFunc("/choice", func(w http.ResponseWriter, r *http.Request) {
		menu := r.FormValue("level")
		fmt.Println(menu)
	})

	http.HandleFunc("/level", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "level", nil)
	})

	http.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "menu", nil)

	})

	//lie le css
	rootDoc, _ := os.Getwd()
	fmt.Println("Serveur écoutant sur le port 8085")
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8085", nil)

}
