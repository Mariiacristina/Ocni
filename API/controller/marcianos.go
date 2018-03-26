package controller

import(
  "net/http"
  "log"
  "encoding/json"
  "Ovni/API/model"
  "Ovni/API/shema"
  "github.com/gorilla/mux"
)

//obtener marcianoSSS
func GetMarcianos(w http.ResponseWriter, r *http.Request) {
	log.Println("hiciste un get a todos los marcianos")
}

//Insertar marciano
func PostMarciano(w http.ResponseWriter, r *http.Request) {
  var alien schema.Marciano
  _=json.NewDecoder(r.Body).Decode(&alien)
  log.Println("vamos a insertar: ",alien)
  err := model.InsertMarciano(alien)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error Lo puse yo jiji", http.StatusInternalServerError)
    return
  }

}


//get marciano
func GetMarciano(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  id := vars["id"]
  log.Println("vamos a obtener un marciano de id: ", id)
  marcianoDeVio, err := model.GetMarciano(id)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error lo puse yo jiji", http.StatusInternalServerError)
    return
  } else {
    log.Println(marcianoDeVio)
    res,errjson := json.Marshal(marcianoDeVio)
        if errjson != nil {
            log.Println(errjson)
              http.Error(w, "Internal server error lo puse yo jiji", http.StatusInternalServerError)
              return
            }
            w.Header().Set("Content-Type", "application/json")
            w.Write(res)
    }
}

func DeleteMarciano(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id := vars["id"]
  log.Println("Vamos a deletear a un marciano de id: ",id)
  marcianoDeVio,err := model.DeleteMarciano(id)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  } else {
    log.Println("se elimino a: ",marcianoDeVio)
          }
}
