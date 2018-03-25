package model

import(
  "log"
  "Ovni/API/shema"
  "Ovni/API/connection"
  "time"
)

//inserta un marciano, mayuscula pq es public
func InsertRevision(paco schema.Revision)(err error){
  connection.Connect()
  t := time.Now()
  log.Println("LLEGO AL MODELO,  INSERT REVISION")
  _,err = db.Exec("INSERT INTO REVISION VALUES (?,?,?,?)",paco.Id,paco.Nombre_r,paco.Id_Aero,t)
  connection.Disconnect()
  return err
}
