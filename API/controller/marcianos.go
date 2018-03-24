package controller

import (
  "net/http"
  "log"
  "fmt"
  "encoding/json"
  "Ovni/API/model"
  "Ovni/API/shema"
)

//obtener marcianoSSS
func GetMarcianos(w http.ResponseWriter, r *http.Request) {
	log.Println("Todo ok :)")
  fmt.Println("holi")
}


func CreateMarciano(w http.ResponseWriter, r *http.Request) {
  var alien schema.Marciano
  _= json.NewDecoder(r.Body).Decode(&alien)
  log.Println("vamos a insertar: ", alien)
  model.InsertMarciano(alien)
}




func GetMarciano(w http.ResponseWriter, r *http.Request) {}
func DeleteMarciano(w http.ResponseWriter, r *http.Request) {}
