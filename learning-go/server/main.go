package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"Anonim", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("../frontend/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
