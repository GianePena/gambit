package secretm

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/gambit/awsgo"
	"github.com/gambit/models"
)

// cuando secret manager lee el secreto lo guarda en una estructura
// es toda la funcion  de secret manager
func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println("> Pido secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg) //INICIALIZO EL SERVICIO: PASO POR PARAMETRO LO QUE YA HICE EN EL LOGUE POR AMAZON --> LE PASO LAS VARIABLES DE ENTORNO DE AMAZON NO DE LA LAMBDA
	//svc ya tiene toda la importacion por lo que ya es un secretmanager
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}
	//secrttring es la colleccion de los 6 datos quese devolvieron--> json con 6 campos
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret) //todos los datos os va a guardar en datos ecret en la structura de models
	fmt.Println("> LECTURA SECRET OK " + nombreSecret)
	return datosSecret, nil
}

//este paquete los vamos a llamar desde el paquete de la base de datos
