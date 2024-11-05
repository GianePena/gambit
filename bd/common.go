package bd

import (
	"database/sql"
	"fmt"
	"github.com/gambit/models"
	"github.com/gambit/secretm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// funciones que se relaiona con la base de datos
var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB //todo lo que tenga que ver con conexion a base de datos se maneja con punteros

func ReadSecret() error {
	var err error
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

// conexion a base de datos
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping() //es un segundo control de conexion --> s devuelve error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("conexxion exitosa de base de datos")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "database-1"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	return dsn
}

//preguntar si el usario es admin

func UserIsAdmin(userUUID string) (bool, string) {
	fmt.Println("COMIENZA USER IS ADMNIN")
	err := DbConnect()
	if err != nil {
		return false, "ERROR EN CONEXION A BD EN USER IS ADMIN" + err.Error()
	}
	defer Db.Close()
	sentencia := "'SELECT 1 from users qhere User_UUID'" + userUUID + "'and User_Status = 0'"
	fmt.Println(sentencia)

	rows, err := Db.Query(sentencia)
	if err != nil {
		return false, err.Error()
	}
	var valor string
	rows.Next()
	rows.Scan(&valor)
	fmt.Println("UserIsAdmin -- Ejecucion exitosa" + valor)
	if valor == "1" {
		return true, ""
	}
	return false, "User in not admin"
}
