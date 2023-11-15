package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	temp, err := template.ParseGlob("./template/*.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur => %s", err.Error()))
		return
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "home", nil)
	})

	http.ListenAndServe("localhost:8085", nil)
}
