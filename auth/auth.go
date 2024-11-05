package auth

import (
	"encoding/base64" //importo solo una parte del paquete porque el paque entero es muy pesado
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub        string
	Event_Id   string
	Token_use  string
	Scope      string
	Auth_time  int
	Iss        string
	Exp        int
	Iat        int
	Cliente_id string
	Username   string
}

func ValidoToken(token string) (bool, error, string) {
	// el token viene en 3 partes: encabezado, payload y firma--> separado po un .
	parts := strings.Split(token, ".") //convierte un array de 3 elementos

	if len(parts) != 3 {
		fmt.Println("El token no es valido")
		return false, nil, "El token no es valido"
	}

	//decodificar ese token --> base64
	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("No se puede decodificar la parte del token", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("No se puede decodificar la estructura json", err.Error())
		return false, err, err.Error()
	}
	//validar la expiracion del token porque puede ser valido èro estar expirado
	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0) //modifica la fecha del exp a un formato = que la fecha del time

	if tm.Before(ahora) { //SI TM ES BEFOR QUE AHORA --> INDICA QUE EL TOKEN ESTA EXPORADO
		fmt.Println("Fecha de expiracion token= " + tm.String())
		fmt.Println("Token expirado")
		return false, err, "Token expirado"
	}
	return true, nil, string(tkj.Username)
} //si falla el token devulve false y el error y el mensaje del error, si no falla devuelve true y el resto en nil
