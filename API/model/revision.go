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
  _,err = db.Exec("INSERT INTO REVISION VALUES (?,?,?,?)",paco.Id_revision,paco.Nombre_r,paco.Id_Aero,t)
  if(err != nil) {
    log.Println("error en el modelo!")
    return err
  }else{
  /*  idd_lista := paco.Id_revision
    rows,err := db.Exec("INSERT INTO LISTA_REVISION(ID_REVISION, ID_MARCIANO, ID_LISTA) SELECT ?,ID_PASAJERO,? FROM VIAJE WHERE ESTADO=1 AND ID_AERONAVE = ? ",paco.Id_revision,idd_lista,paco.Id_Aero)
    if(err != nil) {
    log.Println("error en el modelo!")}
    log.Println(rows)*/
    log.Println("en reconstruccion :( ")
  connection.Disconnect(db)
  return err}
}
