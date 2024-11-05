package awsgo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InicializoAWS() {
	Ctx = context.TODO() //todo significa que no ponga ningun tipo de limiatante al contexto, ni variables de entrni--> queda como un contexto vacio
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		panic("Error al cargar la configuracion. aws/config" + err.Error())
	}

}
