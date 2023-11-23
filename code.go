package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var Level string
var file string

func main() { //lire tout nos templates html

	tmpl, err := template.ParseGlob("./template/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home", nil)
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "level", nil)
	})

	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "result", nil)
	})

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "menu", nil)

	})

	http.HandleFunc("/choice", func(w http.ResponseWriter, r *http.Request) {

		Level = r.FormValue("level")
		filename := Selection(Level)
		readfile := getWordFromFile(filename)
		fmt.Println(Level)
	})

	//lie le css
	rootDoc, _ := os.Getwd()
	fmt.Println("Serveur écoutant sur le port 8085")
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8085", nil)

}

func Selection(Level string) string {
	switch Level {
	case "1":
		return "level 1"
	case "2":
		return "level 2"
	case "3":
		return "level 3"
	case "4":
		return "level 4"
	case "5":
		return "level 5"
	case "6":
		return "level 6"
	default:
		return "level 1"

	}
}

func getWordFromFile(filename string) string {
	// filePath, _ := filepath.Abs("../KANTIN-FAGN-FAGNIART-MOUSNIER-HANGMAN/Listes/" + filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

type Hangman struct {
	currentWordState, guesses []string
	maxGuesses, numOfTries    int
	word                      string
	verbose                   bool
}
