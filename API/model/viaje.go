package model

import(
  "log"
  "Ovni/API/shema"
  //"database/sql"
  "Ovni/API/connection"
  "strconv"
  "time"
)

//inserta un marciano, mayuscula pq es public
func InsertViaje(pasajero schema.Viaje)(err error){
  db := connection.Connect()
  t := time.Now()
  log.Println("LLEGO AL MODELO,  INSERT VIAJERO")
  _,err = db.Exec("INSERT INTO VIAJE VALUES (?,?,?)",pasajero.Id_Pasajero,pasajero.Id_Aero, t)
  connection.Disconnect(db)
  return err
}

func DeleteViaje(id string)(row schema.Viaje,err error){
  db := connection.Connect()
  idNumber,_ := strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, DELETE VIAJE")
  var ScanValue schema.Viaje
  err = db.QueryRow("DELETE FROM VIAJE WHERE id_pasajero =?",idNumber).Scan(&ScanValue)
  connection.Disconnect(db)
  return ScanValue,err
}
