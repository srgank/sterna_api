package main

import (
	"database/sql"
	"encoding/base64"
	//	"encoding/json"
	"fmt"
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
	statement, err := conn.Prepare("UPDATE artikli  set sifra = ?, artikal = ?, edm = ?, ref = ?, kataloski_broj = ?, ddv = ?, proizvoditel = ?, kategorija = ? where sifra = ?")

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(artNew.Sifra, artNew.Artikal, artNew.Edm, artNew.Ref, artNew.Kataloski_broj, artNew.Ddv, artNew.Proizvoditel, artNew.Kategorija, artNew.Sifra)

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
	//	src := search_query.SearchBy
	art := search_query.SearchName
	fmt.Println(search_query)
	queryString := ""

	queryString = "select id, sifra, naziv, adresa, tel, mobil, zirosmetka, edb, deponent, sifdejnost, mb, zabeleska1," +
		" zabeleska2, rabat, grad from komintenti where naziv like '" + art + "' LIMIT " + off + "," + lim

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
	return stat
}

func UpdateKomintentMYSQL(komUpd structs.KomintentiItem, conn *sql.DB) bool {
	stat := false
	sql_statement := "UPDATE artikli  set sifra = ?, naziv = ?, adresa = ?, tel = ?, mobil = ?, zirosmetka = ?, edb = ?, deponent = ?,  " +
		"sifdejnost = ?, mb = ?, zabeleska1 = ?, zabeleska2 = ?, rabat = ?, grad = ? where sifra = ?"

	statement, err := conn.Prepare(sql_statement)

	if err != nil {
		fmt.Println(err)
	}
	conn.Ping()
	_, err = statement.Exec(komUpd.Sifra, komUpd.Naziv, komUpd.Adresa, komUpd.Tel, komUpd.Mobil, komUpd.Zirosmetka, komUpd.Edb, komUpd.Deponent,
		komUpd.Sifdejnost, komUpd.Mb, komUpd.Zabeleska1, komUpd.Zabeleska2, komUpd.Rabat, komUpd.Grad, komUpd.Sifra)

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
