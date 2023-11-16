package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() { //lire tout nos templates html
	temp, err := template.ParseGlob("./template/*.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur => %s", err.Error()))
		return
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "home", nil)
	})

	http.HandleFunc("/choice", func(w http.ResponseWriter, r *http.Request) {
		menu := r.FormValue("menu")
		fmt.Println(menu)
	})

	http.HandleFunc("/level", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "level", nil)
	})

	http.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "menu", nil)
	})
	//lie le css
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8085", nil)
}
