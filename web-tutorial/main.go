package main

import (
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
	"web-tutorial/handler"
)

// package level scope, which means it can be access anywhere in this packager.
var tpl *template.Template

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/list", handler.PatientIndex)
	http.HandleFunc("/patient/show", handler.PatientShow)
	http.HandleFunc("/patient/create", handler.PatientCreate)
	http.HandleFunc("/patient/create/process", handler.PatientCreateProcess)
	http.HandleFunc("/patient/update", handler.PatientUpdate)
	http.HandleFunc("/patient/update/process", handler.PatientUpdateProcess)
	http.HandleFunc("/patient/delete/process", handler.PatientDeleteProcess)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}
