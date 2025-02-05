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
		http.Error(w, "Error executing home template", http.StatusInternalServerError)
	}
}
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	errExec := tmpl.ExecuteTemplate(w, "register.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing register template", http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	errExec := tmpl.ExecuteTemplate(w, "login.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing login template", http.StatusInternalServerError)
	}
}
func Dashboard(w http.ResponseWriter, r *http.Request) {

	errExec := tmpl.ExecuteTemplate(w, "dashboard.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing dashboard template", http.StatusInternalServerError)
	}
}
func Logout(w http.ResponseWriter, r *http.Request) {

	errExec := tmpl.ExecuteTemplate(w, "logout.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing logout template", http.StatusInternalServerError)
	}
}
func Post(w http.ResponseWriter, r *http.Request) {

	errExec := tmpl.ExecuteTemplate(w, "post.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing post template", http.StatusInternalServerError)
	}
}

func Comment(w http.ResponseWriter, r *http.Request) {

	errExec := tmpl.ExecuteTemplate(w, "comment.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing comment template", http.StatusInternalServerError)
	}
}
