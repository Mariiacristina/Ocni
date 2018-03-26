package model

import(
  "log"
  "Ovni/API/shema"
  "Ovni/API/connection"
  "strconv"
)

func GetMarcianos(){
  log.Println("LLEGO AL MODELO, GET MARCIANOOOSSSSS")
}

//inserta un marciano, mayuscula pq es public
func InsertMarciano(alien schema.Marciano)(err error){
  db := connection.Connect()
  log.Println("LLEGO AL MODELO,  INSERT MARCIANO")
  _,err = db.Exec("INSERT INTO MARCIANO VALUES (?,?)",alien.Id_marciano,alien.Nombre)
    if(err != nil) {
      log.Println("error en el modelo!")
      }else{ log.Println("Insertando!!")}
  connection.Disconnect(db)
  return err
}


func GetMarciano(id string)(row schema.Marciano,err error){
  db := connection.Connect()
  idNumber, _ := strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, GET MARCIANO")
  var ScanValue schema.Marciano
  err = db.QueryRow("SELECT ID_MARCIANO, NOMBRE FROM MARCIANO WHERE ID_MARCIANO = ?", idNumber).Scan(&ScanValue.Id_marciano,&ScanValue.Nombre)
  if (err != nil){
    log.Println("error en el modelo!")
  }else {
    log.Println("Se GETEO correctamente")
  }
  connection.Disconnect(db)
  return ScanValue, err
}

func DeleteMarciano(id_marciano string)(row schema.Marciano,err error){
  db := connection.Connect()
  idNumber,_ := strconv.Atoi(id_marciano)
  log.Println("LLEGO AL MODELO, DELETE MARCIANO")
  var ScanValue schema.Marciano
  err = db.QueryRow("DELETE FROM MARCIANO WHERE ID_MARCIANO = ?",idNumber).Scan(&ScanValue.Id_marciano)
  if (err != nil){
    log.Println("error en el modelo!")
  }else {
    log.Println("Se deleteo correctamente")
  }
  connection.Disconnect(db)
  return ScanValue,err
}
