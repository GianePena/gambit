package handlers

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/gambit/auth"
	"github.com/gambit/routers"
	"strconv"
)

func Manejadores(path, method, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) { //devuelo el status(int)y mensaje
	fmt.Println("voy a procesas:" + path + "--" + method)
	//SIMEPRE si voy a relizar un delete o put --> siempre hay que recibir en la url el id de lo que hay que actualizar
	id := request.PathParameters["id"]
	//el id puede llegar como alfa numerico pero si llega en formato string hay que convertilo
	idn, _ := strconv.Atoi(id)
	isOk, statusCode, user := ValidoAuthorization(path, method, headers)
	if !isOk {
		return statusCode, user
	}
	switch path[0:4] {
	case "user":
		return ProcesoUsers(body, path, method, user, id, request)
	case "prod":
		return ProcesoProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcesoStock(body, path, method, user, idn, request)
	case "addr":
		return ProcesoAddres(body, path, method, user, idn, request)
	case "cate":
		return ProcesoCategory(body, path, method, user, idn, request)
	case "orde":
		return ProcesoOrder(body, path, method, user, idn, request)
	}
	return 400, "Method Invalid"
}
func ValidoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if path == "product" && method == "GET" || path == "category" && method == "GET" {
		return true, 200, ""
	}
	//si tengo que autorizar
	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido"
	}
	todoOK, err, msg := auth.ValidoToken(token)
	if !todoOK { //pregunto si todo ok es falso
		if err != nil {
			fmt.Println("Error en el token" + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token" + msg)
			return false, 401, msg
		}
	}
	fmt.Println("TOKEN OK")
	return true, 200, msg
}

func ProcesoUsers(body, path, method, user, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}
func ProcesoProducts(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}

func ProcesoCategory(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	}
	return 400, "Method invalid"
}

func ProcesoStock(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}
func ProcesoAddres(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}
func ProcesoOrder(body, path, method, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method invalid"
}
