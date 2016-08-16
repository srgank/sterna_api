package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sterna_api/API/structs"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

func GetArticleList(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var t_art structs.Article
	if conn != nil {
		t_art = GetArticleListMYSQL(sqlData, conn)
	}
	s, _ := json.Marshal(t_art)
	fmt.Println(string([]byte(s)))
	w.Write([]byte(s))
}

func InsertNewArticle(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.ArticleItem
	_ = json.Unmarshal(Respbody, &sqlData)

	//	{
	//		"id": 10000123,
	//		"sifra":"ABC-0025",
	//		"artikal": "TEST ABC",
	//		"edm":"kom",
	//		"ref":"R-002",
	//		"kataloski_broj": "54-56666-89",
	//		"ddv":5,
	//		"proizvoditel":"ATL",
	//		"kategorija":"KT-5"
	//	}

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = InsertNewArticleMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.ArticleItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = UpdateArticleMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.ArticleItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = DeleteArticleMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func GetKomintentList(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var t_art structs.Komintenti
	if conn != nil {
		t_art = GetKomintentListMYSQL(sqlData, conn)
	}
	s, _ := json.Marshal(t_art)
	fmt.Println(string([]byte(s)))
	w.Write([]byte(s))
}

func InsertNewKomintent(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.KomintentiItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = InsertNewKomintentMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateKomintent(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.KomintentiItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = UpdateKomintentMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteKomintent(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.KomintentiItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = DeleteKomintentMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func GetDokumentList(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var t_art structs.Dokumenti
	if conn != nil {
		t_art = GetDokumentListMYSQL(sqlData, conn)
	}
	s, _ := json.Marshal(t_art)
	fmt.Println(string([]byte(s)))
	w.Write([]byte(s))
}

func InsertNewDokument(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.DokumentiItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = InsertNewDokumentMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateDokument(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.DokumentiItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = UpdateDokumentMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteDokument(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.DokumentiItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = DeleteDokumentMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func GetDokumentDetailList(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var t_art structs.DokumentiDetail
	if conn != nil {
		t_art = GetDocumentDetailListMYSQL(sqlData, conn)
	}
	s, _ := json.Marshal(t_art)
	fmt.Println(string([]byte(s)))
	w.Write([]byte(s))
}

func InsertNewDokumentDetail(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.DokumentiDetailItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = InsertNewDokumentDetailMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateDokumentDetail(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.DokumentiDetailItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = UpdateDokumentDetailMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteDokumentDetail(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.DokumentiDetailItem
	_ = json.Unmarshal(Respbody, &sqlData)

	conn := CreateConnectionMSSQLDB()
	defer conn.Close()
	var stat bool
	if conn != nil {
		stat = DeleteDokumentDetailMYSQL(sqlData, conn)
	}
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
