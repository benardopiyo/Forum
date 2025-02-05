package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Wrong home path", http.StatusNotFound)
		return
	}

	tpl, err := template.ParseFiles("../templates/base.html", "../templates/home.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		return
	}

	err = tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, "Error executing home template", http.StatusInternalServerError)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/register.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		return
	}

	errExec := tmpl.ExecuteTemplate(w, "register.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing register template", http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/login.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		return
	}
	errExec := tmpl.ExecuteTemplate(w, "login.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing login template", http.StatusInternalServerError)
	}
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/dashboard.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		return
	}
	errExec := tmpl.ExecuteTemplate(w, "dashboard.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing dashboard template", http.StatusInternalServerError)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/logout.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		return
	}
	errExec := tmpl.ExecuteTemplate(w, "logout.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing logout template", http.StatusInternalServerError)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/post.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		return
	}
	errExec := tmpl.ExecuteTemplate(w, "post.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing post template", http.StatusInternalServerError)
	}
}

func Comment(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../templates/base.html", "../templates/comment.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		return
	}
	errExec := tmpl.ExecuteTemplate(w, "comment.html", nil)
	if errExec != nil {
		http.Error(w, "Error executing comment template", http.StatusInternalServerError)
	}
}
