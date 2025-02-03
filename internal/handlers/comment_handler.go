package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

var (
	tmpl *template.Template
)

func init() {
	var err error
	tmpl, err = template.ParseGlob("../templates/*.html")
	if err != nil {
		fmt.Println("Error parsing template", err)
		return
	}
}
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Wrong home path", http.StatusNotFound)
		return
	}

	errExec := tmpl.ExecuteTemplate(w, "home.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
