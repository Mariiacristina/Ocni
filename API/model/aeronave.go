package model

import(
 //"database/sql"
  "log"
  "Ovni/API/shema"
  "Ovni/API/connection"
  "strconv"
)

var (
  err error
)

func GetAeros(){
  log.Println("LLEGO AL MODELO, GET MARCIANOOOSSSSS")
}

//inserta un marciano, mayuscula pq es public
func InsertAero(aeronave schema.Aero)(err error){
  db := connection.Connect()
  log.Println("LLEGO AL MODELO,  INSERT AERONAVE")
  _,err = db.Exec("INSERT INTO AERONAVE VALUES (?,?,?,?,?)",aeronave.Id,aeronave.Nombre,aeronave.Nave_o,aeronave.Nave_d,aeronave.Cant_max)
  connection.Disconnect(db)
  return err
}


func GetAero(id string)(row schema.Aero,err error){
  db := connection.Connect()
  idNumber, _ := strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, GET AERONAVE")
  var ScanValue schema.Aero
  err = db.QueryRow("SELECT id, nombre, nave_origen, nave_destino, cant_max FROM aeronave WHERE id = ?", idNumber).Scan(&ScanValue)
  connection.Disconnect(db)
  return ScanValue, err
}

func DeleteAero(id string)(row schema.Aero,err error){
  db := connection.Connect()
  idNumber,_ := strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, DELETE AERONAVE")
  var ScanValue schema.Aero
  err = db.QueryRow("DELETE FROM aeronave WHERE id =?",idNumber).Scan(&ScanValue)
  connection.Disconnect(db)
  return ScanValue,err
}
