package structs

type KomintentiItem struct {
	Id         int64  `bson:"id" json:"id"`
	Sifra      string `bson:"sifra" json:"sifra"`
	Naziv      string `bson:"naziv" json:"naziv"`
	Adresa     string `bson:"adresa" json:"adresa"`
	Tel        string `bson:"tel" json:"tel"`
	Mobil      string `bson:"mobil" json:"mobil"`
	Zirosmetka string `bson:"zirismetka" json:"zirismetka"`
	Edb        string `bson:"edb" json:"edb"`
	Deponent   string `bson:"deponent" json:"deponent"`
	Sifdejnost string `bson:"sifradejnost" json:"sifradejnost"`
	Mb         string `bson:"mb" json:"mb"`
	Zabeleska1 string `bson:"zabeleska1" json:"zabeleska1"`
	Zabeleska2 string `bson:"zabeleska2" json:"zabeleska2"`
	Rabat      int64  `bson:"rabat" json:"rabat"`
	Grad       string `bson:"grad" json:"grad"`
}

type Komintenti struct {
	Properties []KomintentiItem `bson:"properties" json:"properties"`
}
