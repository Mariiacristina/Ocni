package connection

import(
  "log"
//  "database/sql"

)


//variables para la conección
var (
	//db  *sql.DB
)

func Connect(){
  log.Println("conectando a la base de dators...")
//  db, err := sql.Open("mysql","root:password@tcp(127.0.0.1:3306)/marcianos")
}

func Disconnect(){
  log.Println("desconectando de la base de dators...")
//  err = db.Close()
}
