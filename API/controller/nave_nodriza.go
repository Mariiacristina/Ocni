package controller

import(
  "net/http"
  "log"
  "encoding/json"
  "Ovni/API/model"
  "Ovni/API/shema"
)

func GetNaves(w http.ResponseWriter, r *http.Request){
  log.Println("hiciste un get a todas las naves")
}

//insertar naves
func PostNave(w http.ResponseWriter, r *http.Request){
  var nodriza schema.Nave
  _=json.NewDecoder(r.Body).Decode(&nodriza)
  log.Println("Vamos a insertar: ", nodriza)
  err := model.InsertNave(nodriza)
  if err != nil {
    log.Println(err)
    http.Error(w,"Internal server error", http.StatusInternalServerError)
    return
  }
}

func GetNave(w http.ResponseWriter, r *http.Request){
  v := r.URL.Query()
  id := v.Get("id")
  log.Println("vamos a obtener la nave de id: ", id)
  nodriza, err := model.GetNave(id)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  } else {
    log.Println(nodriza)
    res,errjson := json.Marshal(nodriza)
        if errjson != nil {
            log.Println(errjson)
              http.Error(w, "Internal server error", http.StatusInternalServerError)
              return
            }
            w.Header().Set("Content-Type", "application/json")
            w.Write(res)
    }
}

func DeleteNave(w http.ResponseWriter, r *http.Request){
  v := r.URL.Query()
  id := v.Get("id")
  log.Println("Vamos a deletear la nave de id: ",id)
  nodriza,err := model.DeleteNave(id)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  } else {
    log.Println("se elimino a: ",nodriza)
          }


}
