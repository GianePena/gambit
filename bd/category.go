package bd

//cada entidad va atenr un archivo en la carpeta de base de datos y router: categori, products, users...

import (
	"database/sql"
	"fmt"
	"github.com/gambit/models"
	// "github.com/gambit/tools"
	_ "github.com/go-sql-driver/mysql"
	// "strconv"
	// "strings"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("insert category")
	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sentencia := "INSERT INTO category (Categ_name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"
	var result sql.Result
	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err
	}
	fmt.Println("Insert category - ejecucion exitosa")
	return LastInsertId, nil
}
