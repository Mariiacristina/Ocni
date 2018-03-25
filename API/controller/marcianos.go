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

//obtener marcianoSSS
func GetMarcianos(w http.ResponseWriter, r *http.Request) {
	log.Println("hiciste un get a todos los marcianos")
}

//Insertar marciano
func PostMarciano(w http.ResponseWriter, r *http.Request) {
  var alien schema.Marciano
  _=json.NewDecoder(r.Body).Decode(&alien)
  log.Println("vamos a insertar: ", alien.Nombre)
  err := model.InsertMarciano(alien)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }

}


//get marciano
func GetMarciano(w http.ResponseWriter, r *http.Request) {
  v := r.URL.Query()
  id := v.Get("id")
  log.Println("vamos a obtener un marciano de id: ", id)
model.GetMarciano(id)



}

func DeleteMarciano(w http.ResponseWriter, r *http.Request) {
  v := r.URL.Query()
  id := v.Get("id")

  log.Println("Vamos a deletear a un marciano de id: ",id)
  model.DeleteMarciano(id)
}
