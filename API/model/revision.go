package model

import(
  "log"
  "Ovni/API/shema"
  "Ovni/API/connection"
  "time"
  "fmt"
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
    rows, err := db.Query("SELECT ID_PASAJERO FROM VIAJE WHERE ESTADO = 1 AND ID_AERONAVE = ? ORDER BY FECHA DESC LIMIT 0,4",paco.Id_Aero)
    if(err != nil) {
    log.Println("error en mostrar pasajeros!")
  }
  defer rows.Close()
  for rows.Next() {
		var iid_marciano int
		if err := rows.Scan(&iid_marciano); err != nil {
			log.Fatal(err)
		}
		fmt.Println("esta arriba :", iid_marciano, "\n")
	}
  connection.Disconnect(db)
  return err}
}
