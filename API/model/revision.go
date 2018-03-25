package model

import(
  "log"
  "Ovni/API/shema"
  //"database/sql"
  "Ovni/API/connection"
  "time"
)

//inserta un marciano, mayuscula pq es public
func InsertRevision(paco schema.Revision)(err error){
  db := connection.Connect()
  t := time.Now()
  log.Println("LLEGO AL MODELO,  INSERT REVISION")
  _,err = db.Exec("INSERT INTO REVISION VALUES (?,?,?,?)",paco.Id,paco.Nombre_r,paco.Id_Aero,t)
  connection.Disconnect(db)
  return err
}
