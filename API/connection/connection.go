
package connection

import(
  "log"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)


//variables para la conección
var (
	db  *sql.DB
  errbdd error
)

func Connect(){
  log.Println("conectando a la base de dators...")
  _, errbdd := sql.Open("mysql", "server=sql10.freemysqlhosting.net;user id=sql10228844;password=HzESsF45YT;")
  if (errbdd != nil){
    log.Println("ERROR BDD")
  } else {
  log.Println("CONECCION UN EXITO")
  }
}

func Disconnect(){
  log.Println("desconectando de la base de dators...")
  errbdd = db.Close()
  if (errbdd != nil){
    log.Println("ERROR BDD")
  }
}
