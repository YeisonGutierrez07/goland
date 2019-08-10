package shared

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "conexion_goland"
)

var db *gorm.DB
var err error

// Init Inicializar coneccion con postgresql
func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	fmt.Println("Successfully connected!")
}

// GetDb funcion para retornar la base de datos para los diferentes modelos
func GetDb() *gorm.DB {
	fmt.Println("----------------------------------------------------------")
	fmt.Println("----------------------------------------------------------")
	fmt.Println(db)
	fmt.Println("----------------------------------------------------------")
	fmt.Println("----------------------------------------------------------")
	return db
}
