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
  router_naves.HandleFunc("/Nave", controller.GetNaves).Methods("GET")
  //postear una nave
  router_naves.HandleFunc("/Nave",controller.PostNave).Methods("POST")
  //obtener una nave
  router_naves.HandleFunc("/Nave/{id}", controller.GetNave).Methods("GET")
  //eliminar una nave
  router_naves.HandleFunc("/Nave/{id}", controller.DeleteNave).Methods("DELETE")
  return router_naves
}


func RouteringAero()(http.Handler){
router_aero := mux.NewRouter()
///////////AERONAVE//////////////////////////////////////////////////////////////
//obtener varias aeronave
router_aero.HandleFunc("/Aero", controller.GetAeros).Methods("GET")
//postear una aeronave
router_aero.HandleFunc("/Aero", controller.PostAero).Methods("POST")
//obtener una aeronave
router_aero.HandleFunc("/Aero/{id}", controller.GetAero).Methods("GET")
//eliminar una AERONAVE
router_aero.HandleFunc("/Aero/{id}", controller.DeleteAero).Methods("DELETE")
return router_aero
}

func RouteringViaje()(http.Handler){
  router_viaje := mux.NewRouter()
  /////////////VIAJEEES/////////////////////////////////////////////////////////
  //insert viaje
  router_viaje.HandleFunc("/Viaje", controller.PostViaje).Methods("POST")
  //delete Viaje
  router_viaje.HandleFunc("/Viaje/{id}", controller.PostViaje).Methods("DELETE")
  return router_viaje
}
