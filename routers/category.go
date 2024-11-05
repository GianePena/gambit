package routers

//cada entidad va atenr un archivo en la carpeta de base de datos y router: categori, products, users...
import (
	"encoding/json"
	//"github.com/aws/aws-lambda-go/events"
	"github.com/gambit/bd"
	"github.com/gambit/models"
	"strconv"
)

func InsertCategory(body, user string) (int, string) {
	var t models.Category
	// en el body viene un json con los datos de las categorias--> convertir el body en ese modelo
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	}
	if len(t.CategName) == 0 {
		return 400, "Debe especificar el nombre de la categoria"
	}
	if len(t.CategPath) == 0 {
		return 400, "Debe especificar el path de la categoria"
	}
	isAdmin, msg := bd.UserIsAdmin(user)
	if !isAdmin {
		return 400, msg
	}
	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrio un erro al instentar realizar el registro de una category" + t.CategName + "--" + err.Error()
	}
	return 200, "{cateID: " + strconv.Itoa(int(result)) + "}"
}
