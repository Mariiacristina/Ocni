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
  _,err = db.Exec("INSERT INTO marcianos VALUES (?,?)",alien.Id,alien.Nombre)
  connection.Disconnect()
  return err
}


func GetMarciano(id string)(row Marciano,err error){
  connection.Connect()
  id,_ = strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, GET MARCIANO")
  row, err = db.QueryRow("SELECT id, nombre FROM marcianos WHERE id=?",id).Scan(&id,&row.nombre)
  connection.Disconnect()
  return row,err
}

func DeleteMarciano(id string){
  log.Println("LLEGO AL MODELO, DELETE MARCIANO")
}
