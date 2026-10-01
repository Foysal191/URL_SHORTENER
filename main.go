<<<<<<< HEAD
package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
)

var urls = make(map[string]string)

func shortKey() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz123456789")

	for {
		b := make([]rune, 4)

		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		key := string(b)

		if _, exists := urls[key]; !exists {
			return key
		}
	}
}

func main() {

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {

			key := strings.TrimPrefix(r.URL.Path, "/")

			if target, ok := urls[key]; ok {
				http.Redirect(w, r, target, http.StatusSeeOther)
				return
			}

			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")

		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		url := strings.TrimSpace(r.FormValue("url"))

		if url == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(url, "http://") &&
			!strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		key := shortKey()
		urls[key] = url

		tmpl, err := template.ParseFiles("templates/index.html")

		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]string{
			"ShortURL": fmt.Sprintf("http://%s/%s", r.Host, key),
		})
	})

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
=======
package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
)

var urls = make(map[string]string)

func shortKey() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz123456789")

	for {
		b := make([]rune, 4)

		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		key := string(b)

		if _, exists := urls[key]; !exists {
			return key
		}
	}
}

func main() {

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {

			key := strings.TrimPrefix(r.URL.Path, "/")

			if target, ok := urls[key]; ok {
				http.Redirect(w, r, target, http.StatusSeeOther)
				return
			}

			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles("templates/index.html")

		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		url := strings.TrimSpace(r.FormValue("url"))

		if url == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(url, "http://") &&
			!strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		key := shortKey()
		urls[key] = url

		tmpl, err := template.ParseFiles("templates/index.html")

		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]string{
			"ShortURL": fmt.Sprintf("http://%s/%s", r.Host, key),
		})
	})

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
>>>>>>> e14adc1597877f2a12868b8a89de88ee0786570e
}