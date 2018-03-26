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

////////////NAVE NODRIZA //////////////////////////////////////////////////////
  //obtener varias naves
  router.HandleFunc("/Nave", controller.GetNaves).Methods("GET")
  //postear una nave
  router.HandleFunc("/Nave",controller.PostNave).Methods("POST")
  //obtener una nave
  router.HandleFunc("/Nave/{id}", controller.GetNave).Methods("GET")
  //eliminar una nave
  router.HandleFunc("/Nave/{id}", controller.DeleteNave).Methods("DELETE")

///////////AERONAVE//////////////////////////////////////////////////////////////
  //obtener varias aeronave
  router.HandleFunc("/Aero", controller.GetAeros).Methods("GET")
  //postear una aeronave
  router.HandleFunc("/Aero", controller.PostAero).Methods("POST")
  //obtener una aeronave
  router.HandleFunc("/Aero/{id}", controller.GetAero).Methods("GET")
  //eliminar una AERONAVE
  router.HandleFunc("/Aero/{id}", controller.DeleteAero).Methods("DELETE")

  /////////////VIAJEEES/////////////////////////////////////////////////////////
  //insert viaje
  router.HandleFunc("/Viaje", controller.PostViaje).Methods("POST")
  //delete Viaje
  router.HandleFunc("/Adios", controller.DeleteViaje).Methods("POST")
  

  /////////////////REVISION//////////////////////////////////////////////////////
  //insert REVISION
  router.HandleFunc("/Revision", controller.PostRevision).Methods("POST")
  //router.HandleFunc("/Revision", controller.GetRevision).Methods("GET")
  return router
}
