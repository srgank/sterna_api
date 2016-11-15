package structs

type DokumentiDetailItem struct {
	TID                 string `bson:"TID" json:"TID"`
	SUB_TID             int64  `bson:"SUB_TID" json:"SUB_TID"`
	DOCUMENT_ID         int64  `bson:"DOCUMENT_ID" json:"DOCUMENT_ID"`
	DOCUMENT_TIP        int64  `bson:"DOCUMENT_TIP" json:"DOCUMENT_TIP"`
	KOMINTENT_ID        string `bson:"KOMINTENT_ID" json:"KOMINTENT_ID"`
	ARTIKAL_ID          string `bson:"ARTIKAL_ID" json:"ARTIKAL_ID"`
	ARTIKAL             string `bson:"ARTIKAL" json:"ARTIKAL"`
	TIP_ARTIKAL         string `bson:"TIP_ARTIKAL" json:"TIP_ARTIKAL"`
	LINK_ARTIKAL        string `bson:"LINK_ARTIKAL" json:"LINK_ARTIKAL"`
	EDM                 string `bson:"EDM" json:"EDM"`
	V_NAB_CENA_BEZ_DDV  string `bson:"V_NAB_CENA_BEZ_DDV" json:"V_NAB_CENA_BEZ_DDV"`
	V_NAB_CENA_SO_DDV   string `bson:"V_NAB_CENA_SO_DDV" json:"V_NAB_CENA_SO_DDV"`
	V_PREN_DDV          string `bson:"V_PREN_DDV" json:"V_PREN_DDV"`
	V_PREN_DDV_DEN      string `bson:"V_PREN_DDV_DEN" json:"V_PREN_DDV_DEN"`
	V_RABAT             string `bson:"V_RABAT" json:"V_RABAT"`
	V_NAB_IZNOS_SO_DDV  string `bson:"V_NAB_IZNOS_SO_DDV" json:"V_NAB_IZNOS_SO_DDV"`
	V_MARZA             string `bson:"V_MARZA" json:"V_MARZA"`
	V_MARZA_DEN         string `bson:"V_MARZA_DEN" json:"V_MARZA_DEN"`
	V_PROD_CENA_BEZ_DDV string `bson:"V_PROD_CENA_BEZ_DDV" json:"V_PROD_CENA_BEZ_DDV"`
	V_PRESMETAN_DDV     string `bson:"V_PRESMETAN_DDV" json:"V_PRESMETAN_DDV"`
	V_PROD_CENA_SO_DDV  string `bson:"V_PROD_CENA_SO_DDV" json:"V_PROD_CENA_SO_DDV"`
	V_PROD_IZNOS_SO_DDV string `bson:"V_PROD_IZNOS_SO_DDV" json:"V_PROD_IZNOS_SO_DDV"`
	I_CENA_BEZ_DDV_KALK string `bson:"I_CENA_BEZ_DDV_KALK" json:"I_CENA_BEZ_DDV_KALK"`
	I_CENA_SO_DDV_KALK  string `bson:"I_CENA_SO_DDV_KALK" json:"I_CENA_SO_DDV_KALK"`
	I_CENA_SO_DDV_PROD  string `bson:"I_CENA_SO_DDV_PROD" json:"I_CENA_SO_DDV_PROD"`
	I_DDV_PROD          string `bson:"I_DDV_PROD" json:"I_DDV_PROD"`
	I_PROD_IZNOS_SO_DDV string `bson:"I_PROD_IZNOS_SO_DDV" json:"I_PROD_IZNOS_SO_DDV"`
	KOL                 string `bson:"KOL" json:"KOL"`
	MAG_ID              string `bson:"MAG_ID" json:"MAG_ID"`
	STATUS              string `bson:"STATUS" json:"STATUS"`
}

type DokumentiDetail struct {
	Properties []DokumentiDetailItem `bson:"properties" json:"properties"`
}
