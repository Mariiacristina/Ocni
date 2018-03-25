package controller

import(
  "net/http"
  "log"
  "encoding/json"
  "Ovni/API/model"
  "Ovni/API/shema"
)

func GetAeros(w http.ResponseWriter, r *http.Request) {
	log.Println("hiciste un get a todas las aeronaves")
}

//Insertar aeronave
func PostAero(w http.ResponseWriter, r *http.Request) {
  var aeronave schema.Aero
  _=json.NewDecoder(r.Body).Decode(&aeronave)
  log.Println("vamos a insertar: ", aeronave.Nombre)
  err := model.InsertAero(aeronave)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }

}


//get aeronave
func GetAero(w http.ResponseWriter, r *http.Request){
  v := r.URL.Query()
  id := v.Get("id")
  log.Println("vamos a obtener la aronave de id: ", id)
  aeronave, err := model.GetAero(id)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  } else {
    log.Println(aeronave)
    res,errjson := json.Marshal(aeronave)
        if errjson != nil {
            log.Println(errjson)
              http.Error(w, "Internal server error", http.StatusInternalServerError)
              return
            }
            w.Header().Set("Content-Type", "application/json")
            w.Write(res)
    }
}

func DeleteAero(w http.ResponseWriter, r *http.Request) {
  v := r.URL.Query()
  id := v.Get("id")
  log.Println("Vamos a deletear la aeronave de id: ",id)
  aeronave,err := model.DeleteAero(id)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  } else {
    log.Println("se elimino a: ",aeronave)
          }
}
