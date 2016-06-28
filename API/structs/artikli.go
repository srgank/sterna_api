package structs

type ArticleItem struct {
	Id             int64  `bson:"id" json:"id"`
	Sifra          string `bson:"sifra" json:"sifra"`
	Artikal        string `bson:"artikal" json:"artikal"`
	Edm            string `bson:"edm" json:"edm"`
	Ref            string `bson:"ref" json:"ref"`
	Kataloski_broj string `bson:"kataloski_broj" json:"kataloski_broj"`
	Ddv            int64  `bson:"ddv" json:"ddv"`
	Proizvoditel   string `bson:"proizvoditel" json:"proizvoditel"`
	Kategorija     string `bson:"kategorija" json:"kategorija"`
}
type Article struct {
	Properties []ArticleItem `bson:"properties" json:"properties"`
}

type SearchByItem struct {
	Limit      string `bson:"limit" json:"limit"`
	Offset     string `bson:"offset" json:"offset"`
	SearchBy   string `bson:"search_by" json:"search_by"`
	SearchName string `bson:"search_name" json:"search_name"`
}
