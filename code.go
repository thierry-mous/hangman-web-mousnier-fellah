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
	
	http.HandleFunc("/level1", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "level1", nil)
	})
	http.HandleFunc("/level2", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "level2", nil)
	})
	http.HandleFunc("/level3", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "level3", nil)
	})
	http.HandleFunc("/level4", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "level4", nil)
	})
	http.HandleFunc("/level5", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "level5", nil)
	})
	http.HandleFunc("/level6", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "level6", nil)
	})
	//lie le css
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8085", nil)
}
