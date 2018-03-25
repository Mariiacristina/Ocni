package main

import (
  "net/http"
  //"github.com/gorilla/mux"
  "log"
  "Ovni/API/rutas"
)

func main() {
  m := rutas.RouteringMarcianos()
	log.Fatal(http.ListenAndServe(":8000",m))

  n := rutas.RouteringNaves()
  log.Fatal(http.ListenAndServe(":8000",n))
}
