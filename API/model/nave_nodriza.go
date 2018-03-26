package model

import(
  "log"
  "Ovni/API/shema"
  //"database/sql"
  "Ovni/API/connection"
  "strconv"
)

func GetNaves(){
  log.Println("LLEGO AL MODELO, GET NAVES")
}

//inserta una nave
func InsertNave(nodriza schema.Nave)(err error){
  db := connection.Connect()
  log.Println("LLEGO AL MODELO, INSERT NAVE")
  _,err = db.Exec("INSERT INTO NAVE_NODRIZA VALUES (?,?)",nodriza.Id_nodriza,nodriza.Nombre)
  if(err != nil){
    log.Println("error en el modelors")
  }else{ log.Println("Insertado!!")}
  connection.Disconnect(db)
  return err
}

func GetNave(id string)(row schema.Nave,err error){
  db := connection.Connect()
  idNumber, _ := strconv.Atoi(id)
  log.Println("LLEGO AL MODELO, GET NAVE")
  var ScanValue schema.Nave
  err = db.QueryRow("SELECT ID_NODRIZA, NOMBRE FROM NAVE_NODRIZA WHERE ID_NODRIZA = ?", idNumber).Scan(&ScanValue.Id_nodriza,&ScanValue.Nombre)
  if (err != nil){
    log.Println("error en el modelo!")
  }else {
    log.Println("Se GETEO correctamente")
  }
  connection.Disconnect(db)
  return ScanValue, err
}

func DeleteNave(id string)(row schema.Nave,err error){
  db := connection.Connect()
  log.Println(id)
  idNumber,_ := strconv.Atoi(id)
  log.Println(idNumber)
  var ScanValue schema.Nave
  err = db.QueryRow("DELETE FROM NAVE_NODRIZA WHERE ID_NODRIZA = ?",idNumber).Scan(&ScanValue.Id_nodriza)
  if (err != nil){
    log.Println("error en el modelo!")
  }else {
    log.Println("Se deleteo correctamente")
  }
  connection.Disconnect(db)
  return ScanValue,err
}
