package rutas

import (
  "github.com/gorilla/mux"
  "github.com/Mariiacristina/Ocni/API/controller"
  "net/http"
)

func RouteringMarcianos() (http.Handler) {
	router := mux.NewRouter()

  //obtener varios marcianos
	router.HandleFunc("/Marciano", controller.GetMarcianos).Methods("GET")
  //postear un marciano
  router.HandleFunc("/Marciano", controller.PostMarciano).Methods("POST")
  //obtener un marciano
	router.HandleFunc("/Marciano/{id}", controller.GetMarciano).Methods("GET")
  //eliminar un marciano
  router.HandleFunc("/Marciano/{id}", controller.DeleteMarciano).Methods("DELETE")



	return router
}
