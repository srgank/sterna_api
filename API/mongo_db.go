package main

import (
	"database/sql"
	"encoding/base64"
	//	"encoding/json"
	"fmt"
	"time"
	//	"io/ioutil"
	//	"net/http"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sterna_api/API/structs"
)

func CreateConnectionMSSQLDB() *sql.DB {
	var conn *sql.DB
	//conn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sternaX")
	connectionString := "root:office22@tcp(92.53.51.86:3306)/test"
	//	conn, err = sql.Open("mssql", "server=testbaza.mssql.somee.com;user id=srdzank_SQLLogin_1;password=ahxb4ednub;port=1433")
	conn, err := sql.Open("mysql", connectionString)
	if err != nil {
		conn = nil
	}
	return conn
}

func GetArticleListMYSQL(search_query structs.SearchByItem, conn *sql.DB) structs.Article {

	lim := search_query.Limit
	off := search_query.Offset
	src := search_query.SearchBy
	art := search_query.SearchName
	fmt.Println(search_query)
	queryString := ""
	if src == "artikal" {
		queryString = "select id, sifra, artikal, edm, ref, kataloski_broj, ddv, proizvoditel, kategorija  from artikli where artikal like '" + art + "' LIMIT " + off + "," + lim
	} else if src == "sifra" {
		queryString = "select id, sifra, artikal, edm, ref, kataloski_broj, ddv, proizvoditel, kategorija  from artikli where sifra = '" + art + "' LIMIT " + off + "," + lim
	} else if src == "id" {
		queryString = "select id, sifra, artikal, edm, ref, kataloski_broj, ddv, proizvoditel, kategorija  from artikli where id = '" + art + "' LIMIT " + off + "," + lim
	} else {
	}

	statement, err := conn.Prepare(queryString)
	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	rows, err := statement.Query() // execute our select statement
	if err != nil {
		fmt.Println(err)
	}

	var art1 structs.Article
	art1.Properties = nil

	var id *int64
	var sifra *string
	var artikal *string
	var edm *string
	var ref *string
	var kataloski_broj *string
	var proizvoditel *string
	var ddv *int64
	var kategorija *string

	var artItem structs.ArticleItem

	for rows.Next() {
		rows.Scan(&id, &sifra, &artikal, &edm, &ref, &kataloski_broj, &ddv, &proizvoditel, &kategorija)

		if id != nil {
			artItem.Id = *id
		}
		if sifra != nil {
			artItem.Sifra = base64.StdEncoding.EncodeToString([]byte(*sifra))
		}
		if artikal != nil {
			artItem.Artikal = base64.StdEncoding.EncodeToString([]byte(*artikal))
		}
		if edm != nil {
			artItem.Edm = base64.StdEncoding.EncodeToString([]byte(*edm))
		}
		if ref != nil {
			artItem.Ref = base64.StdEncoding.EncodeToString([]byte(*ref))
		}
		if kataloski_broj != nil {
			artItem.Kataloski_broj = base64.StdEncoding.EncodeToString([]byte(*kataloski_broj))
		}
		if ddv != nil {
			artItem.Ddv = *ddv
		}
		if proizvoditel != nil {
			artItem.Proizvoditel = base64.StdEncoding.EncodeToString([]byte(*proizvoditel))
		}
		if kategorija != nil {
			artItem.Kategorija = base64.StdEncoding.EncodeToString([]byte(*kategorija))
		}

		art1.Properties = append(art1.Properties, artItem)
	}
	return art1
}

func InsertNewArticleMYSQL(artNew structs.ArticleItem, conn *sql.DB) bool {
	stat := false
	statement, err := conn.Prepare("INSERT INTO artikli( sifra, artikal, edm, ref, kataloski_broj, ddv, proizvoditel, kategorija) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(artNew.Sifra, artNew.Artikal, artNew.Edm, artNew.Ref, artNew.Kataloski_broj, artNew.Ddv, artNew.Proizvoditel, artNew.Kategorija)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	return stat
}

func UpdateArticleMYSQL(artNew structs.ArticleItem, conn *sql.DB) bool {
	stat := false
	statement, err := conn.Prepare("UPDATE artikli  set sifra = ?, artikal = ?, edm = ?, ref = ?, kataloski_broj = ?, ddv = ?, proizvoditel = ?, kategorija = ? where id = ?")
	fmt.Printf("UpdateArticleMYSQL ------------------------------------------------------------------------------------", artNew)
	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(artNew.Sifra, artNew.Artikal, artNew.Edm, artNew.Ref, artNew.Kataloski_broj, artNew.Ddv, artNew.Proizvoditel, artNew.Kategorija, artNew.Id)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	return stat
}

func DeleteArticleMYSQL(artNew structs.ArticleItem, conn *sql.DB) bool {
	stat := false
	statement, err := conn.Prepare("delete from artikli where id = ?")

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(artNew.Id)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	return stat
}

//komintenti

func GetKomintentListMYSQL(search_query structs.SearchByItem, conn *sql.DB) structs.Komintenti {

	lim := search_query.Limit
	off := search_query.Offset
	src := search_query.SearchBy
	art := search_query.SearchName
	fmt.Println(search_query)
	queryString := ""

	if src == "naziv" {
		queryString = "select id, sifra, naziv, adresa, tel, mobil, zirosmetka, edb, deponent, sifdejnost, mb, zabeleska1," +
			" zabeleska2, rabat, grad from komintenti where naziv like '" + art + "' LIMIT " + off + "," + lim
	} else if src == "sifra" {
		queryString = "select id, sifra, naziv, adresa, tel, mobil, zirosmetka, edb, deponent, sifdejnost, mb, zabeleska1," +
			" zabeleska2, rabat, grad from komintenti where sifra like '" + art + "' LIMIT " + off + "," + lim
	} else {
	}

	statement, err := conn.Prepare(queryString)
	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	rows, err := statement.Query() // execute our select statement
	if err != nil {
		fmt.Println(err)
	}

	var kom1 structs.Komintenti
	kom1.Properties = nil

	var id *int64
	var sifra *string
	var naziv *string
	var adresa *string
	var tel *string
	var mobil *string
	var zirosmetka *string
	var edb *string
	var deponent *string
	var sifdejnost *string
	var mb *string
	var zabeleska1 *string
	var zabeleska2 *string
	var rabat *int64
	var grad *string

	var komItem structs.KomintentiItem

	for rows.Next() {
		rows.Scan(&id, &sifra, &naziv, &adresa, &tel, &mobil, &zirosmetka, &edb, &deponent, &sifdejnost, &mb, &zabeleska1, &zabeleska2, &rabat, &grad)

		if id != nil {
			komItem.Id = *id
		}
		if sifra != nil {
			komItem.Sifra = base64.StdEncoding.EncodeToString([]byte(*sifra))
		}
		if naziv != nil {
			komItem.Naziv = base64.StdEncoding.EncodeToString([]byte(*naziv))
		}
		if adresa != nil {
			komItem.Adresa = base64.StdEncoding.EncodeToString([]byte(*adresa))
		}
		if tel != nil {
			komItem.Tel = base64.StdEncoding.EncodeToString([]byte(*tel))
		}

		if mobil != nil {
			komItem.Mobil = base64.StdEncoding.EncodeToString([]byte(*mobil))
		}

		if zirosmetka != nil {
			komItem.Zirosmetka = base64.StdEncoding.EncodeToString([]byte(*zirosmetka))
		}
		if edb != nil {
			komItem.Edb = base64.StdEncoding.EncodeToString([]byte(*edb))
		}
		if deponent != nil {
			komItem.Deponent = base64.StdEncoding.EncodeToString([]byte(*deponent))
		}
		if sifdejnost != nil {
			komItem.Sifdejnost = base64.StdEncoding.EncodeToString([]byte(*sifdejnost))
		}
		if mb != nil {
			komItem.Mb = base64.StdEncoding.EncodeToString([]byte(*mb))
		}
		if zabeleska1 != nil {
			komItem.Zabeleska1 = base64.StdEncoding.EncodeToString([]byte(*zabeleska1))
		}
		if zabeleska2 != nil {
			komItem.Zabeleska2 = base64.StdEncoding.EncodeToString([]byte(*zabeleska2))
		}
		if rabat != nil {
			komItem.Rabat = *rabat
		}
		if grad != nil {
			komItem.Grad = base64.StdEncoding.EncodeToString([]byte(*grad))
		}

		kom1.Properties = append(kom1.Properties, komItem)
	}
	return kom1
}

func InsertNewKomintentMYSQL(komNew structs.KomintentiItem, conn *sql.DB) bool {
	stat := false

	sql_statement := "INSERT INTO komintenti( sifra, naziv, adresa, tel, mobil, zirosmetka, edb, deponent, " +
		"sifdejnost, mb, zabeleska1, zabeleska2, rabat, grad ) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	statement, err := conn.Prepare(sql_statement)

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(komNew.Sifra, komNew.Naziv, komNew.Adresa, komNew.Tel, komNew.Mobil, komNew.Zirosmetka,
		komNew.Edb, komNew.Deponent, komNew.Sifdejnost, komNew.Mb, komNew.Zabeleska1, komNew.Zabeleska2, komNew.Rabat, komNew.Grad)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	fmt.Println("InsertNewKomintentMYSQL", stat)
	return stat
}

func UpdateKomintentMYSQL(komUpd structs.KomintentiItem, conn *sql.DB) bool {
	stat := false
	sql_statement := "UPDATE komintenti  set sifra = ?, naziv = ?, adresa = ?, tel = ?, mobil = ?, zirosmetka = ?, edb = ?, deponent = ?,  " +
		"sifdejnost = ?, mb = ?, zabeleska1 = ?, zabeleska2 = ?, rabat = ?, grad = ? where id = ?"

	statement, err := conn.Prepare(sql_statement)

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(komUpd.Sifra, komUpd.Naziv, komUpd.Adresa, komUpd.Tel, komUpd.Mobil, komUpd.Zirosmetka, komUpd.Edb, komUpd.Deponent,
		komUpd.Sifdejnost, komUpd.Mb, komUpd.Zabeleska1, komUpd.Zabeleska2, komUpd.Rabat, komUpd.Grad, komUpd.Id)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	return stat
}

func DeleteKomintentMYSQL(komDlt structs.KomintentiItem, conn *sql.DB) bool {
	stat := false
	statement, err := conn.Prepare("delete from komintenti where id = ?")

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(komDlt.Id)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	return stat
}

//dokumenti

func GetDokumentListMYSQL(search_query structs.SearchByItem, conn *sql.DB) structs.Dokumenti {

	lim := search_query.Limit
	off := search_query.Offset

	fmt.Println(search_query)
	queryString := ""

	queryString = " select TID, DOCUMENT_ID, DOCUMENT_TIP, TD, TDS, KOMINTENT_ID,k.naziv as KOMINTENT_NAZIV, PREVOZNIK_ID,p.naziv as PREVOZNIK_NAZIV, VALUTA, KURS, IZNOS_VAL, DDV_VAL, " +
		" RABAT_VAL, IZNOS_PLAKANJE_VAL, IZNOS_DEN, IZNOS_DDV_DEN, RABAT_DEN, IZNOS_PLAKANJE_DEN, TRANSPORT_DEN, CARINA_DEN, " +
		" DDV_DEN, DRUGI_TROSOCI_DEN, DOK_STATUS, USER_ID, KOMENTAR, MAG_ID, OBJECT_ID " +
		" from dokumenti d " +
		" left join komintenti k on d.KOMINTENT_ID = k.id " +
		" left join komintenti p on d.KOMINTENT_ID = p.id LIMIT " + off + "," + lim

	statement, err := conn.Prepare(queryString)
	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	rows, err := statement.Query() // execute our select statement
	if err != nil {
		fmt.Println(err)
	}

	var dok1 structs.Dokumenti
	dok1.Properties = nil

	var TID *int64
	var DOCUMENT_ID *int64
	var DOCUMENT_TIP *int64
	var TD *time.Time
	var TDS *time.Time
	var KOMINTENT_ID *int64
	var KOMINTENT_NAZIV *string
	var PREVOZNIK_ID *int64
	var PREVOZNIK_NAZIV *string
	var VALUTA *string
	var KURS *float64
	var IZNOS_VAL *float64
	var DDV_VAL *float64
	var RABAT_VAL *float64
	var IZNOS_PLAKANJE_VAL *float64
	var IZNOS_DEN *float64
	var IZNOS_DDV_DEN *float64
	var RABAT_DEN *float64
	var IZNOS_PLAKANJE_DEN *float64
	var TRANSPORT_DEN *float64
	var CARINA_DEN *float64
	var DDV_DEN *float64
	var DRUGI_TROSOCI_DEN *float64
	var DOK_STATUS *string
	var USER_ID *int64
	var KOMENTAR *string
	var MAG_ID *int64
	var OBJECT_ID *int64

	var dokItem structs.DokumentiItem

	for rows.Next() {
		rows.Scan(&TID, &DOCUMENT_ID, &DOCUMENT_TIP, &TD, &TDS, &KOMINTENT_ID, &KOMINTENT_NAZIV, &PREVOZNIK_ID, &PREVOZNIK_NAZIV, &VALUTA, &KURS, &IZNOS_VAL, &DDV_VAL,
			&RABAT_VAL, &IZNOS_PLAKANJE_VAL, &IZNOS_DEN, &IZNOS_DDV_DEN, &RABAT_DEN, &IZNOS_PLAKANJE_DEN, &TRANSPORT_DEN, &CARINA_DEN,
			&DDV_DEN, &DRUGI_TROSOCI_DEN, &DOK_STATUS, &USER_ID, &KOMENTAR, &MAG_ID, &OBJECT_ID)

		if TID != nil {
			dokItem.TID = *TID
		}
		if DOCUMENT_ID != nil {
			dokItem.DOCUMENT_ID = *DOCUMENT_ID
		}
		if DOCUMENT_TIP != nil {
			dokItem.DOCUMENT_TIP = *DOCUMENT_TIP
		}
		if TD != nil {
			dokItem.TD = *TD
		}
		if TDS != nil {
			dokItem.TDS = *TDS
		}

		if KOMINTENT_ID != nil {
			dokItem.KOMINTENT_ID = *KOMINTENT_ID
		}

		if PREVOZNIK_ID != nil {
			dokItem.PREVOZNIK_ID = *PREVOZNIK_ID
		}
		if PREVOZNIK_NAZIV != nil {
			dokItem.PREVOZNIK_ID = *PREVOZNIK_ID
		}
		if VALUTA != nil {
			dokItem.VALUTA = *VALUTA
		}
		if KURS != nil {
			dokItem.KURS = *KURS
		}
		if IZNOS_VAL != nil {
			dokItem.IZNOS_VAL = *IZNOS_VAL
		}
		if DDV_VAL != nil {
			dokItem.DDV_VAL = *DDV_VAL
		}
		if RABAT_VAL != nil {
			dokItem.RABAT_VAL = *RABAT_VAL
		}
		if IZNOS_PLAKANJE_VAL != nil {
			dokItem.IZNOS_PLAKANJE_VAL = *IZNOS_PLAKANJE_VAL
		}
		if IZNOS_DEN != nil {
			dokItem.IZNOS_DEN = *IZNOS_DEN
		}
		if IZNOS_DDV_DEN != nil {
			dokItem.IZNOS_DDV_DEN = *IZNOS_DDV_DEN
		}

		if RABAT_DEN != nil {
			dokItem.RABAT_DEN = *RABAT_DEN
		}
		if IZNOS_PLAKANJE_DEN != nil {
			dokItem.IZNOS_PLAKANJE_DEN = *IZNOS_PLAKANJE_DEN
		}
		if TRANSPORT_DEN != nil {
			dokItem.TRANSPORT_DEN = *TRANSPORT_DEN
		}
		if CARINA_DEN != nil {
			dokItem.CARINA_DEN = *CARINA_DEN
		}
		if DDV_DEN != nil {
			dokItem.DDV_DEN = *DDV_DEN
		}
		if DRUGI_TROSOCI_DEN != nil {
			dokItem.DRUGI_TROSOCI_DEN = *DRUGI_TROSOCI_DEN
		}
		if DOK_STATUS != nil {
			dokItem.DOK_STATUS = *DOK_STATUS
		}

		if USER_ID != nil {
			dokItem.USER_ID = *USER_ID
		}
		if KOMENTAR != nil {
			dokItem.KOMENTAR = *KOMENTAR
		}
		if MAG_ID != nil {
			dokItem.MAG_ID = *MAG_ID
		}
		if OBJECT_ID != nil {
			dokItem.OBJECT_ID = *OBJECT_ID
		}

		dok1.Properties = append(dok1.Properties, dokItem)
	}
	return dok1
}

func InsertNewDokumentMYSQL(dokNew structs.DokumentiItem, conn *sql.DB) bool {
	stat := false

	sql_statement := "INSERT INTO dokumenti( " +
		"DOCUMENT_ID, " +
		"DOCUMENT_TIP , " +
		"TD , " +
		"TDS , " +
		"KOMINTENT_ID , " +
		"PREVOZNIK_ID , " +
		"VALUTA , " +
		"KURS , " +
		"IZNOS_VAL , " +
		"DDV_VAL , " +
		"RABAT_VAL , " +
		"IZNOS_PLAKANJE_VAL , " +
		"IZNOS_DEN , " +
		"IZNOS_DDV_DEN , " +
		"RABAT_DEN , " +
		"IZNOS_PLAKANJE_DEN , " +
		"TRANSPORT_DEN , " +
		"CARINA_DEN , " +
		"DDV_DEN , " +
		"DRUGI_TROSOCI_DEN , " +
		"DOK_STATUS , " +
		"USER_ID , " +
		"KOMENTAR , " +
		"MAG_ID , " +
		"OBJECT_ID  )" +
		"VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	statement, err := conn.Prepare(sql_statement)

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()

	_, err = statement.Exec(
		dokNew.TID,
		dokNew.DOCUMENT_ID,
		dokNew.DOCUMENT_TIP,
		dokNew.TD,
		dokNew.TDS,
		dokNew.KOMINTENT_ID,
		dokNew.PREVOZNIK_ID,
		dokNew.VALUTA,
		dokNew.KURS,
		dokNew.IZNOS_VAL,
		dokNew.DDV_VAL,
		dokNew.RABAT_VAL,
		dokNew.IZNOS_PLAKANJE_VAL,
		dokNew.IZNOS_DEN,
		dokNew.IZNOS_DDV_DEN,
		dokNew.RABAT_DEN,
		dokNew.IZNOS_PLAKANJE_DEN,
		dokNew.TRANSPORT_DEN,
		dokNew.CARINA_DEN,
		dokNew.DDV_DEN,
		dokNew.DRUGI_TROSOCI_DEN,
		dokNew.DOK_STATUS,
		dokNew.USER_ID,
		dokNew.KOMENTAR,
		dokNew.MAG_ID,
		dokNew.OBJECT_ID,
	)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	fmt.Println("InsertNewKomintentMYSQL", stat)
	return stat
}

func UpdateDokumentMYSQL(dokUpd structs.DokumentiItem, conn *sql.DB) bool {
	stat := false
	sql_statement := "UPDATE dokumenti  set " +
		"DOCUMENT_ID = ?, " +
		"DOCUMENT_TIP = ?, " +
		"TD = ?, " +
		"TDS = ?, " +
		"KOMINTENT_ID = ?, " +
		"PREVOZNIK_ID = ?, " +
		"VALUTA = ?, " +
		"KURS = ?, " +
		"IZNOS_VAL = ?, " +
		"DDV_VAL = ?, " +
		"RABAT_VAL = ?, " +
		"IZNOS_PLAKANJE_VAL = ?, " +
		"IZNOS_DEN = ?, " +
		"IZNOS_DDV_DEN = ?, " +
		"RABAT_DEN = ?, " +
		"IZNOS_PLAKANJE_DEN = ?, " +
		"TRANSPORT_DEN = ?, " +
		"CARINA_DEN = ?, " +
		"DDV_DEN = ?, " +
		"DRUGI_TROSOCI_DEN = ?, " +
		"DOK_STATUS = ?, " +
		"USER_ID = ?, " +
		"KOMENTAR = ?, " +
		"MAG_ID = ?, " +
		"OBJECT_ID = ? " +
		"where TID = ? "

	statement, err := conn.Prepare(sql_statement)

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(
		dokUpd.TID,
		dokUpd.DOCUMENT_ID,
		dokUpd.DOCUMENT_TIP,
		dokUpd.TD,
		dokUpd.TDS,
		dokUpd.KOMINTENT_ID,
		dokUpd.PREVOZNIK_ID,
		dokUpd.VALUTA,
		dokUpd.KURS,
		dokUpd.IZNOS_VAL,
		dokUpd.DDV_VAL,
		dokUpd.RABAT_VAL,
		dokUpd.IZNOS_PLAKANJE_VAL,
		dokUpd.IZNOS_DEN,
		dokUpd.IZNOS_DDV_DEN,
		dokUpd.RABAT_DEN,
		dokUpd.IZNOS_PLAKANJE_DEN,
		dokUpd.TRANSPORT_DEN,
		dokUpd.CARINA_DEN,
		dokUpd.DDV_DEN,
		dokUpd.DRUGI_TROSOCI_DEN,
		dokUpd.DOK_STATUS,
		dokUpd.USER_ID,
		dokUpd.KOMENTAR,
		dokUpd.MAG_ID,
		dokUpd.OBJECT_ID,
	)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	return stat
}

func DeleteDokumentMYSQL(dokDlt structs.DokumentiItem, conn *sql.DB) bool {
	stat := false
	statement, err := conn.Prepare("delete from dokumenti where id = ?")

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(dokDlt.TID)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	return stat
}

//dokument_detail

func GetDocumentDetailListMYSQL(search_query structs.SearchByItem, conn *sql.DB) structs.DokumentiDetail {

	lim := search_query.Limit
	off := search_query.Offset
	dok_tip := search_query.Dok_TIP
	dok_id := search_query.Dok_ID
	fmt.Println(search_query)
	queryString := ""

	queryString = "select " +
		" TID, DOCUMENT_ID, DOCUMENT_TIP, KOMINTENT_ID, ARTIKAL_ID, ARTIKAL, TIP_ARTIKAL, LINK_ARTIKAL, m.EDM as EDM1, " +
		" V_NAB_CENA_BEZ_DDV, V_NAB_CENA_SO_DDV, V_PREN_DDV, V_PREN_DDV_DEN, V_RABAT, V_NAB_IZNOS_SO_DDV, " +
		" V_MARZA, V_MARZA_DEN, V_PROD_CENA_BEZ_DDV, V_PRESMETAN_DDV, V_PROD_CENA_SO_DDV, V_PROD_IZNOS_SO_DDV, " +
		" I_CENA_BEZ_DDV_KALK, I_CENA_SO_DDV_KALK, I_CENA_SO_DDV_PROD, I_DDV_PROD, KOL, MAG_ID, STATUS " +
		" from  magacin m" +
		" left join artikli a  on m.ARTIKAL_ID =  a.id " +
		" where DOCUMENT_TIP = " + dok_tip + " and DOCUMENT_ID = " + dok_id + " LIMIT " + off + "," + lim

	statement, err := conn.Prepare(queryString)
	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	rows, err := statement.Query() // execute our select statement
	if err != nil {
		fmt.Println(err)
	}

	var dokDetail structs.DokumentiDetail
	dokDetail.Properties = nil

	var TID *int64
	var DOCUMENT_ID *int64
	var DOCUMENT_TIP *int64
	var KOMINTENT_ID *int64
	var ARTIKAL_ID *int64
	var ARTIKAL *string
	var TIP_ARTIKAL *int64
	var LINK_ARTIKAL *int64
	var EDM1 *string
	var V_NAB_CENA_BEZ_DDV *float64
	var V_NAB_CENA_SO_DDV *float64
	var V_PREN_DDV *float64
	var V_PREN_DDV_DEN *float64
	var V_RABAT *float64
	var V_NAB_IZNOS_SO_DDV *float64
	var V_MARZA *float64
	var V_MARZA_DEN *float64
	var V_PROD_CENA_BEZ_DDV *float64
	var V_PRESMETAN_DDV *float64
	var V_PROD_CENA_SO_DDV *float64
	var V_PROD_IZNOS_SO_DDV *float64
	var I_CENA_BEZ_DDV_KALK *float64
	var I_CENA_SO_DDV_KALK *float64
	var I_CENA_SO_DDV_PROD *float64
	var I_DDV_PROD *float64
	var KOL *float64
	var MAG_ID *int64
	var STATUS *int64

	var dokDetailItem structs.DokumentiDetailItem

	for rows.Next() {
		rows.Scan(
			&TID, &DOCUMENT_ID, &DOCUMENT_TIP, &KOMINTENT_ID, &ARTIKAL_ID, &ARTIKAL, &TIP_ARTIKAL, &LINK_ARTIKAL, &EDM1,
			&V_NAB_CENA_BEZ_DDV, &V_NAB_CENA_SO_DDV, &V_PREN_DDV, &V_PREN_DDV_DEN, &V_RABAT, &V_NAB_IZNOS_SO_DDV,
			&V_MARZA, &V_MARZA_DEN, &V_PROD_CENA_BEZ_DDV, &V_PRESMETAN_DDV, &V_PROD_CENA_SO_DDV, &V_PROD_IZNOS_SO_DDV,
			&I_CENA_BEZ_DDV_KALK, &I_CENA_SO_DDV_KALK, &I_CENA_SO_DDV_PROD, &I_DDV_PROD, &KOL, &MAG_ID, &STATUS,
		)

		if TID != nil {
			dokDetailItem.TID = *TID
		}
		if DOCUMENT_ID != nil {
			dokDetailItem.DOCUMENT_ID = *DOCUMENT_ID
		}
		if DOCUMENT_TIP != nil {
			dokDetailItem.DOCUMENT_TIP = *DOCUMENT_TIP
		}
		if KOMINTENT_ID != nil {
			dokDetailItem.KOMINTENT_ID = *KOMINTENT_ID
		}
		if ARTIKAL_ID != nil {
			dokDetailItem.ARTIKAL_ID = *ARTIKAL_ID
		}
		if ARTIKAL != nil {
			dokDetailItem.ARTIKAL = *ARTIKAL
		}

		if TIP_ARTIKAL != nil {
			dokDetailItem.TIP_ARTIKAL = *TIP_ARTIKAL
		}
		if LINK_ARTIKAL != nil {
			dokDetailItem.LINK_ARTIKAL = *LINK_ARTIKAL
		}
		if EDM1 != nil {
			dokDetailItem.EDM = *EDM1
		}
		if V_NAB_CENA_BEZ_DDV != nil {
			dokDetailItem.V_NAB_CENA_BEZ_DDV = *V_NAB_CENA_BEZ_DDV
		}
		if V_NAB_CENA_SO_DDV != nil {
			dokDetailItem.V_NAB_CENA_SO_DDV = *V_NAB_CENA_SO_DDV
		}
		if V_PREN_DDV != nil {
			dokDetailItem.V_PREN_DDV = *V_PREN_DDV
		}
		if V_PREN_DDV_DEN != nil {
			dokDetailItem.V_PREN_DDV_DEN = *V_PREN_DDV_DEN
		}
		if V_RABAT != nil {
			dokDetailItem.V_RABAT = *V_RABAT
		}
		if V_NAB_IZNOS_SO_DDV != nil {
			dokDetailItem.V_NAB_IZNOS_SO_DDV = *V_NAB_IZNOS_SO_DDV
		}
		if V_MARZA != nil {
			dokDetailItem.V_MARZA = *V_MARZA
		}
		if V_MARZA_DEN != nil {
			dokDetailItem.V_MARZA_DEN = *V_MARZA_DEN
		}
		if V_PROD_CENA_BEZ_DDV != nil {
			dokDetailItem.V_PROD_CENA_BEZ_DDV = *V_PROD_CENA_BEZ_DDV
		}
		if V_PRESMETAN_DDV != nil {
			dokDetailItem.V_PRESMETAN_DDV = *V_PRESMETAN_DDV
		}
		if V_PROD_CENA_SO_DDV != nil {
			dokDetailItem.V_PROD_CENA_SO_DDV = *V_PROD_CENA_SO_DDV
		}
		if V_PROD_IZNOS_SO_DDV != nil {
			dokDetailItem.V_PROD_IZNOS_SO_DDV = *V_PROD_IZNOS_SO_DDV
		}
		if I_CENA_BEZ_DDV_KALK != nil {
			dokDetailItem.I_CENA_BEZ_DDV_KALK = *I_CENA_BEZ_DDV_KALK
		}
		if I_CENA_SO_DDV_KALK != nil {
			dokDetailItem.I_CENA_SO_DDV_KALK = *I_CENA_SO_DDV_KALK
		}
		if I_CENA_SO_DDV_PROD != nil {
			dokDetailItem.I_CENA_SO_DDV_PROD = *I_CENA_SO_DDV_PROD
		}
		if I_DDV_PROD != nil {
			dokDetailItem.I_DDV_PROD = *I_DDV_PROD
		}
		if KOL != nil {
			dokDetailItem.KOL = *KOL
		}
		if MAG_ID != nil {
			dokDetailItem.MAG_ID = *MAG_ID
		}
		if STATUS != nil {
			dokDetailItem.STATUS = *STATUS
		}

		dokDetail.Properties = append(dokDetail.Properties, dokDetailItem)
	}
	return dokDetail
}

func InsertNewDokumentDetailMYSQL(docDetailNew structs.DokumentiDetailItem, conn *sql.DB) bool {
	stat := false

	sql_statement := "INSERT INTO komintenti( sifra, naziv, adresa, tel, mobil, zirosmetka, edb, deponent, " +
		"sifdejnost, mb, zabeleska1, zabeleska2, rabat, grad ) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	statement, err := conn.Prepare(sql_statement)

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()

	_, err = statement.Exec(
		docDetailNew.DOCUMENT_ID, docDetailNew.DOCUMENT_TIP, docDetailNew.KOMINTENT_ID, docDetailNew.ARTIKAL_ID, docDetailNew.TIP_ARTIKAL, docDetailNew.LINK_ARTIKAL, docDetailNew.EDM,
		docDetailNew.V_NAB_CENA_BEZ_DDV, docDetailNew.V_NAB_CENA_SO_DDV, docDetailNew.V_PREN_DDV, docDetailNew.V_PREN_DDV_DEN, docDetailNew.V_RABAT, docDetailNew.V_NAB_IZNOS_SO_DDV,
		docDetailNew.V_MARZA, docDetailNew.V_MARZA_DEN, docDetailNew.V_PROD_CENA_BEZ_DDV, docDetailNew.V_PRESMETAN_DDV, docDetailNew.V_PROD_CENA_SO_DDV, docDetailNew.V_PROD_IZNOS_SO_DDV,
		docDetailNew.I_CENA_BEZ_DDV_KALK, docDetailNew.I_CENA_SO_DDV_KALK, docDetailNew.I_CENA_SO_DDV_PROD, docDetailNew.I_DDV_PROD, docDetailNew.KOL, docDetailNew.MAG_ID, docDetailNew.STATUS,
	)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	fmt.Println("InsertNewKomintentMYSQL", stat)
	return stat
}

func UpdateDokumentDetailMYSQL(docDetailUpd structs.DokumentiDetailItem, conn *sql.DB) bool {
	stat := false
	sql_statement := "UPDATE magacin  set " +
		" DOCUMENT_ID = ?, " +
		" DOCUMENT_TIP = ?, " +
		" KOMINTENT_ID = ?, " +
		" ARTIKAL_ID = ?, " +
		" TIP_ARTIKAL = ?, " +
		" LINK_ARTIKAL = ?, " +
		" EDM = ?, " +
		" V_NAB_CENA_BEZ_DDV = ?, " +
		" V_NAB_CENA_SO_DDV = ?, " +
		" V_PREN_DDV = ?, " +
		" V_PREN_DDV_DEN = ?, " +
		" V_RABAT = ?, " +
		" V_NAB_IZNOS_SO_DDV = ?, " +
		" V_MARZA = ?, " +
		" V_MARZA_DEN = ?, " +
		" V_PROD_CENA_BEZ_DDV = ?, " +
		" V_PRESMETAN_DDV = ?, " +
		" V_PROD_CENA_SO_DDV = ?, " +
		" V_PROD_IZNOS_SO_DDV = ?, " +
		" I_CENA_BEZ_DDV_KALK = ?, " +
		" I_CENA_SO_DDV_KALK = ?, " +
		" I_CENA_SO_DDV_PROD = ?, " +
		" I_DDV_PROD = ?, " +
		" KOL = ?, " +
		" MAG_ID = ?, " +
		" STATUS = ? " +
		"where TID = ?"

	statement, err := conn.Prepare(sql_statement)

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(
		docDetailUpd.TID, docDetailUpd.DOCUMENT_ID, docDetailUpd.DOCUMENT_TIP, docDetailUpd.KOMINTENT_ID, docDetailUpd.ARTIKAL_ID, docDetailUpd.TIP_ARTIKAL, docDetailUpd.LINK_ARTIKAL, docDetailUpd.EDM,
		docDetailUpd.V_NAB_CENA_BEZ_DDV, docDetailUpd.V_NAB_CENA_SO_DDV, docDetailUpd.V_PREN_DDV, docDetailUpd.V_PREN_DDV_DEN, docDetailUpd.V_RABAT, docDetailUpd.V_NAB_IZNOS_SO_DDV,
		docDetailUpd.V_MARZA, docDetailUpd.V_MARZA_DEN, docDetailUpd.V_PROD_CENA_BEZ_DDV, docDetailUpd.V_PRESMETAN_DDV, docDetailUpd.V_PROD_CENA_SO_DDV, docDetailUpd.V_PROD_IZNOS_SO_DDV,
		docDetailUpd.I_CENA_BEZ_DDV_KALK, docDetailUpd.I_CENA_SO_DDV_KALK, docDetailUpd.I_CENA_SO_DDV_PROD, docDetailUpd.I_DDV_PROD, docDetailUpd.KOL, docDetailUpd.MAG_ID, docDetailUpd.STATUS,
	)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	return stat
}

func DeleteDokumentDetailMYSQL(docDetailDlt structs.DokumentiDetailItem, conn *sql.DB) bool {
	stat := false
	statement, err := conn.Prepare("delete from magacin where id = ?")

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(docDetailDlt.TID)

	if err != nil {
		stat = false
	} else {
		stat = true
	}
	return stat
}
