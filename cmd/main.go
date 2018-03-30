package main

import (
  "net/http"
  "log"
  "Ovni/API/rutas"
)

func main() {
  m := rutas.RouteringMarcianos()
	log.Fatal(http.ListenAndServe(":4242",m))

}
