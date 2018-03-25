
package connection

import(
  "log"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)


//variables para la conección
var (
	db  *sql.DB
  err error
)

func Connect(){
  log.Println("conectando a la base de dators...")
  _, err := sql.Open("mysql", "server=sql10.freemysqlhosting.net;user id=sql10228844;password=HzESsF45YT;")
  if (err != nil){
    log.Println("ERROR BDD")
  } else {
  log.Println("CONECCION UN EXITO")
  }
}

func Disconnect(){
  log.Println("desconectando de la base de dators...")
  err = db.Close()
  if (err != nil){
    log.Println("ERROR BDD")
  }
}
