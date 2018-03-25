package connection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("psql","gkeppswcmlwpxe:be21b2bda8836c06764653af84916f2ef6b987d24bc6bc4491f8a0e404759c09@tcp(ec2-107-20-249-48.compute-1.amazonaws.com:5432)/delm769cqht311")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}