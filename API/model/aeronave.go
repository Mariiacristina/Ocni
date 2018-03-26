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
  _,err = db.Exec("INSERT INTO AERONAVE VALUES (?,?,?,?,?)",aeronave.Id,aeronave.Nombre,aeronave.Nave_o,aeronave.Nave_d,aeronave.Cant_max)
  if(err != nil) {
    log.Println("error en el modelo!")
  }else{ log.Println("Insertando!!")}
  connection.Disconnect(db)
  return err
}


func GetAero(id string)(row schema.Aero,err error){
  db := connection.Connect()
  idNumber, _ := strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, GET AERONAVE")
  var ScanValue schema.Aero
  err = db.QueryRow("SELECT ID_AERONAVE, NOMBRE, NAVE_ORIGEN, NAVE_DESTINO, CANT_MAX FROM AERONAVE WHERE ID_AERONAVE = ?", idNumber).Scan(&ScanValue.Id,&ScanValue.Nombre,&ScanValue.Nave_o,&ScanValue.Nave_d,&ScanValue.Cant_max)
  if (err != nil){
    log.Println("error en el modelo!")
  }else {
    log.Println("Se GETEO correctamente")
  }
  connection.Disconnect(db)
  return ScanValue, err
}

func DeleteAero(id string)(row schema.Aero,err error){
  db := connection.Connect()
  log.Println(id)
  idNumber,_ := strconv.Atoi(id)
  log.Println(idNumber)
  var ScanValue schema.Aero
  err = db.QueryRow("DELETE FROM AERONAVE WHERE ID_AERONAVE = ?",idNumber).Scan(&ScanValue.Id)
  if (err != nil){
    log.Println("error en el modelo!")
  }else {
    log.Println("Se deleteo correctamente")
  }
  connection.Disconnect(db)
  return ScanValue,err
}
