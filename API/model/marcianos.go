package model

import(
 "database/sql"
  "log"
  "Ovni/API/shema"
  "Ovni/API/connection"
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
  err = db.QueryRow("INSERT INTO marcianos VALUES (?,?)").Scan(alien.Id,alien.Nombre)
  connection.Disconnect()
  return err
}


func GetMarciano(id string){
  log.Println("LLEGO AL MODELO, GET MARCIANO")
}

func DeleteMarciano(id string){
  log.Println("LLEGO AL MODELO, DELETE MARCIANO")
}
