package model

import(
  "log"
  "Ovni/API/shema"
  "Ovni/API/connection"
  "strconv"
)

func GetNaves(){
  log.Println("LLEGO AL MODELO, GET NAVES")
}

//inserta una nave
func InsertNave(nodriza schema.Nave)(err error){
  connection.Connect()
  log.Println("LLEGO AL MODELO, INSERT NAVE")
  _,err = db.Exec("INSERT INTO NAVE_NODRIZA VALUES (?,?)",nodriza.Id,nodriza.Nombre)
  connection.Disconnect()
  return err
}

func GetNave(id string)(row schema.Nave,err error){
  connection.Connect()
  idNumber, _ := strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, GET NAVE")
  var ScanValue schema.Nave
  err = db.QueryRow("SELECT id, nombre FROM NAVE_NODRIZA WHERE id = ?", idNumber).Scan(&ScanValue)
  connection.Disconnect()
  return ScanValue, err
}

func DeleteNave(id string)(row schema.Nave,err error){
  connection.Connect()
  idNumber,_ := strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, DELETE NAVE")
  var ScanValue schema.Nave
  err = db.QueryRow("DELETE FROM marciano WHERE id =?",idNumber).Scan(&ScanValue)
  connection.Disconnect()
  return ScanValue,err
}
