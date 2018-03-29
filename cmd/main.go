package main

import (
  "net/http"
  //"github.com/gorilla/mux"
  "log"
  "html/template"
  "Ovni/API/rutas"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

  m := rutas.RouteringMarcianos()
	log.Fatal(http.ListenAndServe(":8000",m))

  http.HandleFunc("/", index)
	http.HandleFunc("/asignar", asignar)
	http.HandleFunc("/bajar", bajar)
	http.HandleFunc("/aeronave", aeronave)
	http.HandleFunc("/revision", revision)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func asignar(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "asignar.gohtml", nil)
}

func bajar(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "bajar.gohtml", nil)
}

func aeronave(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "aeronave.gohtml", nil)
}

func revision(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "revision.gohtml", nil)
}
