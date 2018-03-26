package model

import(
  "log"
  "Ovni/API/shema"
  //"database/sql"
  "Ovni/API/connection"
  "time"
)
var(
  err1 error
  err2 error
  err3 error
  cant int
  count int
  count2 int
  countt int
)

//inserta un marciano, mayuscula pq es public
func InsertViaje(pasajero schema.Viaje)(err error){
  db := connection.Connect()
  t := time.Now()
  log.Println("id de Aero es: ",pasajero.Id_Aero)

//cantidad de gente que admite la AERONAVE
  err1 = db.QueryRow("SELECT CANT_MAX FROM AERONAVE WHERE ID_AERONAVE = ?",pasajero.Id_Aero).Scan(&cant)
  if (err1 != nil){
    log.Println("fallo en sacando cantidad de pasajeros aeronave")
  }
  log.Println("cantidad maxima de pasajeros es: ",cant)
  //cantidad de personas arriba de la aeronave (historica)
  err2 = db.QueryRow("SELECT COUNT(ESTADO) FROM VIAJE WHERE ID_AERONAVE = ? AND ESTADO = 1",pasajero.Id_Aero).Scan(&count)
  log.Println("cantidad de los que van arriba son: ", count)
  if (err2 != nil){
    log.Println("error en contar la cantidad de pasajeros")
  }
  err3 = db.QueryRow("SELECT COUNT(ESTADO) FROM VIAJE WHERE ID_AERONAVE = ? AND ESTADO = 0",pasajero.Id_Aero).Scan(&count2)
  log.Println("cantidad de los que se bajaron son: ", count2)
  if (err3 != nil){
    log.Println("error en contar la cantidad de pasajeros")
  }
  //total de los que estan arriba
  countt = count - count2

  if(countt+1>cant){
    log.Println("no se puede subir :( hay muchos")
    return
  }else{
  log.Println("INSERT VIAJERO")
  _,err = db.Exec("INSERT INTO VIAJE VALUES (?,?,?,?,?)",pasajero.Id_Pasajero,pasajero.Id_Aero, t,pasajero.Id_viaje,1)
  connection.Disconnect(db)
  return err
  }
}

func DeleteViaje(pasajero schema.Viaje)(err error){
  db := connection.Connect()
  t := time.Now()
  log.Println("BAJAR VIAJERO")
  _,err = db.Exec("INSERT INTO VIAJE VALUES (?,?,?,?,?)",pasajero.Id_Pasajero,pasajero.Id_Aero, t,pasajero.Id_viaje,0)
  connection.Disconnect(db)
  return err
  }
