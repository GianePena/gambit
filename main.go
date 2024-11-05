package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/gambit/awsgo"
	"github.com/gambit/bd"
	"github.com/gambit/handlers"
	"os"
	"strings"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InicializoAWS()

	if !ValidoParametros() {
		panic("Error en los parametros debe enviar: secretname y UrlPrefix")
	}
	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1) //eliminar el prefijo de la ruta (prefix) en request.RawPath. Esto permite obtener la ruta relativa sin el prefijo base
	method := request.RequestContext.HTTP.Method
	body := request.Body //SE PASA COMO PARAMETRO A LAS DIFERENTES RUTAS
	header := request.Headers
	bd.ReadSecret()
	//LLAMADO AL HANDLER(MANEJADOR)--> creo 2 variables porque la funcion devulve 2 respuestas
	status, message := handlers.Manejadores(path, method, body, header, request)
	//RESPUESTA DE LA API
	headersResp := map[string]string{
		"Content-Type": "application/json",
	}
	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}
	return res, nil

}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}
	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro
}
