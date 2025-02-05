package routes

import (
	"net/http"
	"fora/internal/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/register", handlers.Register)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/dashboard", handlers.Dashboard)
	mux.HandleFunc("/logout", handlers.Logout)
	mux.HandleFunc("/post", handlers.Post)
	mux.HandleFunc("/comment", handlers.Comment)

	fileserver := http.FileServer(http.Dir("../static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileserver))
	return mux
}