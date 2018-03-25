package rutas

import (
  "github.com/gorilla/mux"
  "Ovni/API/controller"
  "net/http"
)

func RouteringMarcianos() (http.Handler) {
	router := mux.NewRouter()
////////////MARCIANOS////////////////////////////////////////////////////////
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

func RouteringNaves() (http.Handler) {
  router_naves := mux.NewRouter()
////////////NAVE NODRIZA //////////////////////////////////////////////////////
  //obtener varias naves
  router_naves.HandlFunc("/Nave", controller.GetNaves).Methods("GET")
  //postear una nave
  router_naves.HandlFunc("/Nave",controller.PostNave).Methods("POST")
  //obtener una nave
  router_naves.HandlFunc("/Nave/{id}", controller.GetNave).Methods("GET")
  //eliminar una nave
  router_naves.HandleFunc("/Nave/{id}", controller.DeleteNave).Methods("DELETE")
  return router_naves
}
