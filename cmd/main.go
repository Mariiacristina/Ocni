package main

import (
  "net/http"
  //"github.com/gorilla/mux"
  "log"
  "github.com/Mariiacristina/Ocni/API/rutas"
)

func main() {
  r := rutas.RouteringMarcianos()
	log.Fatal(http.ListenAndServe(":8000",r))
}
