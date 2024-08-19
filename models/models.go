package models

type SecretRDSJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Engine string `json:"engine"`
	Host string `json:"host"`
	Port int `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

type TokenJSON struct {
	Sub string
	EventID string `json:"event_id"`
	TokenUse string `json:"token_use"`
	Scope string
	AuthTime int `json:"auth_time"`
	Iss string
	Exp int
	Iat int
	ClientID string `json:"client_id"`
	Username string
}

type User struct {
	Id_User string `json:"ID_User"`
	First_Name string `json:"firstName"`
	Last_Name string `json:"lastName"`
	UserEmail string `json:"userEmail"`
	Age int `json:"age"`
	Weight float32 `json:"weight"`
	Height float32 `json:"height"`
	UserStatus int `json:"userStatus"`
	UserDateAdd string `json:"userDateAdd"`
	UserDateUpd string `json:"userDateUpd"`
}

type Muscle struct {
	Id_Muscle int `json:"ID_Muscle"`
	Name_Muscle string `json:"nameMuscle"`
}