package controller

import (
  "net/http"
  "log"
  "encoding/json"
  "Ovni/API/model"
  "Ovni/API/shema"
)

func PostViaje(w http.ResponseWriter, r *http.Request) {
  var pasajero schema.Viaje
  _=json.NewDecoder(r.Body).Decode(&pasajero)
  log.Println("vamos a al pasajero en la nave : ", pasajero.Id_Pasajero, pasajero.Id_Aero)
  err := model.InsertViaje(pasajero)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }
}


func DeleteViaje(w http.ResponseWriter, r *http.Request) {
  var pasajero schema.Viaje
    _=json.NewDecoder(r.Body).Decode(&pasajero)
    log.Println("vamos a al pasajero en la nave : ", pasajero.Id_Pasajero, pasajero.Id_Aero)
    err := model.DeleteViaje(pasajero)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }
}
