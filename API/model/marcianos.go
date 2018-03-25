package model

import(
 "database/sql"
  "log"
  "Ovni/API/shema"
  "Ovni/API/connection"
  "strconv"
)

var (
  err error
  db *sql.DB
)

func GetMarcianos(){
  log.Println("LLEGO AL MODELO, GET MARCIANOOOSSSSS")
}

//inserta un marciano, mayuscula pq es public
func InsertMarciano(alien schema.Marciano)(err error){
  connection.Connect()
  log.Println("LLEGO AL MODELO,  INSERT MARCIANO")
  _,err = db.Exec("INSERT INTO marciano VALUES (?,?)",alien.Id,alien.Nombre)
  connection.Disconnect()
  return err
}


func GetMarciano(id string)(row schema.Marciano,err error){
  connection.Connect()
  idNumber, _ := strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, GET MARCIANO")
  var ScanValue schema.Marciano
  err = db.QueryRow("SELECT id, nombre FROM marciano WHERE id = ?", idNumber).Scan(&ScanValue)
  connection.Disconnect()
  return ScanValue, err
}

func DeleteMarciano(id string)(row schema.Marciano,err error){
  connection.Connect()
  idNumber,_ := strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, DELETE MARCIANO")
  var ScanValue schema.Marciano
  err = db.QueryRow("DELETE FROM marciano WHERE id =?",idNumber).Scan(&ScanValue)
  connection.Disconnect()
  return ScanValue,err
}
