package structs

type AccountItem struct {
	Id             string `bson:"id" json:"id"`
	UserName       string `bson:"user_name" json:"user_name"`
	Password       string `bson:"password" json:"password"`
	Company        string `bson:"company" json:"company"`
	Email          string `bson:"email" json:"email"`
	Token          string `bson:"token" json:"token"`
	ValidTo        int64  `bson:"valid_to" json:"valid_to"`
	DatabasePrefix string `bson:"database_prefix" json:"database_prefix"`
}
