package controller

import (
  "net/http"
  "log"
  "encoding/json"
  "Ovni/API/model"
  "Ovni/API/shema"
)
//var (
//  err error
//)

func PostViaje(w http.ResponseWriter, r *http.Request) {
  var pasajero schema.Viaje
  _=json.NewDecoder(r.Body).Decode(&pasajero)
  log.Println("vamos a al pasajero %s en la nave %s : ", pasajero.Id_Pasajero, pasajero.Id_Aero)
  err := model.InsertViaje(pasajero)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }
}


func DeleteViaje(w http.ResponseWriter, r *http.Request) {
  v := r.URL.Query()
  id := v.Get("id")
  log.Println("Vamos a deletear el viaje del id_pasajero: ",id)
  pasajero,err := model.DeleteViaje(id)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  } else {
    log.Println("se elimino el viaje de: ",pasajero.Id_Pasajero)
          }
}
