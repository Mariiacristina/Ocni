package rutas

import (
  "github.com/gorilla/mux"
  "Ovni/API/controller"
  "net/http"
)

func RouteringMarcianos() (http.Handler) {
	router := mux.NewRouter()
	router.HandleFunc("/Marciano", controller.GetMarcianos).Methods("GET")
  router.HandleFunc("/Marciano", controller.CreateMarciano).Methods("POST")
	router.HandleFunc("/Marciano/{id}", controller.GetMarciano).Methods("GET")
	router.HandleFunc("/Marciano/{id}", controller.DeleteMarciano).Methods("DELETE")
	return router
}
