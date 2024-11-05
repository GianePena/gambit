package models

// aca solo se defiene estructutras
// ESTRUCTUR 1 SECRET MANAGER
type SecretRDSJson struct { //devuelve 6 valores(username, pasword, engine, port,host y dscluster..) em formato json
	Username            string `json:"username"` //hay que aclarar quen va aestar en minuscula: indica que n el formato json que nos va a devolveer es username
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

//ESTRUCTURA 2 REGISTRO DE USUARIOS

type SingUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}

// CATEGORY
type Category struct {
	CategID   int    `json:"categID"` //inidca como viene en el json
	CategName string `json:"categName"`
	CategPath string `json:"categPath"`
}
