package structs

type LoginDataRequest struct {
	UserName string `bson:"user_name" json:"user_name"`
	Password string `bson:"password" json:"password"`
}

type LoginRespond struct {
	Properties []LoginDataRespond `bson:"properties" json:"properties"`
}
type LoginDataRespond struct {
	Token          string `bson:"token" json:"token"`
	ValitTo        string `bson:"valid_to" json:"valid_to"`
	DataBasePrefix string `bson:"database_prefix" json:"database_prefix"`
	Company        string `bson:"company" json:"company"`
}

type Accounts struct {
	UserName       string `bson:"user_name" json:"user_name"`
	Pass           string `bson:"password" json:"password"`
	Company        string `bson:"company" json:"company"`
	Email          string `bson:"email" json:"email"`
	Token          string `bson:"token" json:"token"`
	ValitTo        string `bson:"valid_to" json:"valid_to"`
	DataBasePrefix string `bson:"database_prefix" json:"database_prefix"`
}
