package structs

import (
	"time"
)

type DokumentiItem struct {
	TID                int64     `bson:"TID" json:"TID"`
	DOCUMENT_ID        int64     `bson:"DOCUMENT_ID" json:"DOCUMENT_ID"`
	DOCUMENT_TIP       int64     `bson:"DOCUMENT_TIP" json:"DOCUMENT_TIP"`
	TD                 time.Time `bson:"TD" json:"TD"`
	TDS                time.Time `bson:"TDS" json:"TDS"`
	KOMINTENT_ID       int64     `bson:"KOMINTENT_ID" json:"KOMINTENT_ID"`
	KOMINTENT_NAZIV    string    `bson:"KOMINTENT_NAZIV" json:"KOMINTENT_NAZIV"`
	PREVOZNIK_ID       int64     `bson:"PREVOZNIK_ID" json:"PREVOZNIK_ID"`
	PREVOZNIK_NAZIV    string    `bson:"PREVOZNIK_NAZIV" json:"PREVOZNIK_NAZIV"`
	VALUTA             string    `bson:"VALUTA" json:"VALUTA"`
	KURS               float64   `bson:"KURS" json:"KURS"`
	IZNOS_VAL          float64   `bson:"IZNOS_VAL" json:"IZNOS_VAL"`
	DDV_VAL            float64   `bson:"DDV_VAL" json:"DDV_VAL"`
	RABAT_VAL          float64   `bson:"RABAT_VAL" json:"RABAT_VAL"`
	IZNOS_PLAKANJE_VAL float64   `bson:"IZNOS_PLAKANJE_VAL" json:"IZNOS_PLAKANJE_VAL"`
	IZNOS_DEN          float64   `bson:"IZNOS_DEN" json:"IZNOS_DEN"`
	IZNOS_DDV_DEN      float64   `bson:"IZNOS_DDV_DEN" json:"IZNOS_DDV_DEN"`
	RABAT_DEN          float64   `bson:"RABAT_DEN" json:"RABAT_DEN"`
	IZNOS_PLAKANJE_DEN float64   `bson:"IZNOS_PLAKANJE_DEN" json:"IZNOS_PLAKANJE_DEN"`
	TRANSPORT_DEN      float64   `bson:"TRANSPORT_DEN" json:"TRANSPORT_DEN"`
	CARINA_DEN         float64   `bson:"CARINA_DEN" json:"CARINA_DEN"`
	DDV_DEN            float64   `bson:"DDV_DEN" json:"DDV_DEN"`
	DRUGI_TROSOCI_DEN  float64   `bson:"DRUGI_TROSOCI_DEN" json:"DRUGI_TROSOCI_DEN"`
	DOK_STATUS         string    `bson:"DOK_STATUS" json:"DOK_STATUS"`
	USER_ID            int64     `bson:"USER_ID" json:"USER_ID"`
	KOMENTAR           string    `bson:"KOMENTAR" json:"KOMENTAR"`
	MAG_ID             int64     `bson:"MAG_ID" json:"MAG_ID"`
	OBJECT_ID          int64     `bson:"OBJECT_ID" json:"OBJECT_ID"`
}

type Dokumenti struct {
	Properties []DokumentiItem `bson:"properties" json:"properties"`
}
