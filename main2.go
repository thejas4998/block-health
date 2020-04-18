package main

import (
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

var templates *template.Template

func main(){

	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/hello", indexHandlerFunc).Methods("POST")
	r.HandleFunc("/goodbye", goodbyeHandlerFunc).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)

}	

func indexHandlerFunc(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	templates.ExecuteTemplate(w, "index1.html", nil)

}

func goodbyeHandlerFunc(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "index.html", nil)
}
