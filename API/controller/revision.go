package controller

import (
  "net/http"
  "log"
  "encoding/json"
  "Ovni/API/model"
  "Ovni/API/shema"
)

//Insertar revision
func PostRevision(w http.ResponseWriter, r *http.Request) {
  var paco schema.Revision
  _=json.NewDecoder(r.Body).Decode(&paco)
  log.Println("vamos a revisar a %s por %s : ", paco.Id_Aero, paco.Nombre_r)
  err := model.InsertRevision(paco)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }

}

//get todas las revisiones
/*func GetRevision(w http.ResponseWriter, r *http.Request){
  log.Printlm("vamos a imprimir todas las revisiones")
  err := model.GetRevisiones()
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
  }else{
    log.Println("buena")
  }

}
*/
