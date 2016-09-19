package structs

type DokumentiItem struct {
	TID                string `bson:"TID" json:"TID"`
	DOCUMENT_ID        string `bson:"DOCUMENT_ID" json:"DOCUMENT_ID"`
	DOCUMENT_TIP       string `bson:"DOCUMENT_TIP" json:"DOCUMENT_TIP"`
	TD                 string `bson:"TD" json:"TD"`
	TDS                string `bson:"TDS" json:"TDS"`
	KOMINTENT_ID       string `bson:"KOMINTENT_ID" json:"KOMINTENT_ID"`
	KOMINTENT_NAZIV    string `bson:"KOMINTENT_NAZIV" json:"KOMINTENT_NAZIV"`
	PREVOZNIK_ID       string `bson:"PREVOZNIK_ID" json:"PREVOZNIK_ID"`
	PREVOZNIK_NAZIV    string `bson:"PREVOZNIK_NAZIV" json:"PREVOZNIK_NAZIV"`
	VALUTA             string `bson:"VALUTA" json:"VALUTA"`
	KURS               string `bson:"KURS" json:"KURS"`
	IZNOS_VAL          string `bson:"IZNOS_VAL" json:"IZNOS_VAL"`
	DDV_VAL            string `bson:"DDV_VAL" json:"DDV_VAL"`
	RABAT_VAL          string `bson:"RABAT_VAL" json:"RABAT_VAL"`
	IZNOS_PLAKANJE_VAL string `bson:"IZNOS_PLAKANJE_VAL" json:"IZNOS_PLAKANJE_VAL"`
	IZNOS_DEN          string `bson:"IZNOS_DEN" json:"IZNOS_DEN"`
	IZNOS_DDV_DEN      string `bson:"IZNOS_DDV_DEN" json:"IZNOS_DDV_DEN"`
	RABAT_DEN          string `bson:"RABAT_DEN" json:"RABAT_DEN"`
	IZNOS_PLAKANJE_DEN string `bson:"IZNOS_PLAKANJE_DEN" json:"IZNOS_PLAKANJE_DEN"`
	TRANSPORT_DEN      string `bson:"TRANSPORT_DEN" json:"TRANSPORT_DEN"`
	CARINA_DEN         string `bson:"CARINA_DEN" json:"CARINA_DEN"`
	DDV_DEN            string `bson:"DDV_DEN" json:"DDV_DEN"`
	DRUGI_TROSOCI_DEN  string `bson:"DRUGI_TROSOCI_DEN" json:"DRUGI_TROSOCI_DEN"`
	DOK_STATUS         string `bson:"DOK_STATUS" json:"DOK_STATUS"`
	USER_ID            string `bson:"USER_ID" json:"USER_ID"`
	KOMENTAR           string `bson:"KOMENTAR" json:"KOMENTAR"`
	MAG_ID             string `bson:"MAG_ID" json:"MAG_ID"`
	OBJECT_ID          string `bson:"OBJECT_ID" json:"OBJECT_ID"`
}

type Dokumenti struct {
	Properties []DokumentiItem `bson:"properties" json:"properties"`
}
