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

func GetArticleListMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Respbody", Respbody)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)
	fmt.Println("sqlData", sqlData)
	_, s := GetArticleListMYSQLMongo(sqlData)

	fmt.Println(s)
	w.Write([]byte(s))
}

func InsertNewArticleMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.ArticleItem
	_ = json.Unmarshal(Respbody, &sqlData)

	var stat bool
	stat = InsertNewArticleMYSQLMongo(sqlData)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateArticleMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.ArticleItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = UpdateArticleMYSQLMongo(sqlData)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteArticleMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.ArticleItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = DeleteArticleMYSQLMongo(sqlData)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//-------------------------------------------------------------------------------------

func GeKomintentiListMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Respbody", Respbody)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)
	fmt.Println("sqlData", sqlData)
	_, s := GetKomintentiListMYSQLMongo(sqlData)

	fmt.Println(s)
	w.Write([]byte(s))
}

func InsertNewKomintentiMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.KomintentiItem
	_ = json.Unmarshal(Respbody, &sqlData)

	var stat bool
	stat = InsertNewKomintentiMYSQLMongo(sqlData)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateKomintentiMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.KomintentiItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = UpdateKomintentiMYSQLMongo(sqlData)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteKomintentiMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.KomintentiItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = DeleteKomintentiMYSQLMongo(sqlData)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//-------------------------------------------------------------------------------------

func GetDokumentListMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Respbody", Respbody)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)
	fmt.Println("sqlData", sqlData)
	_, s := GetDokumentiListMYSQLMongo(sqlData)

	fmt.Println(s)
	w.Write([]byte(s))
}

func InsertNewDokumentMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.DokumentiItem
	_ = json.Unmarshal(Respbody, &sqlData)

	var stat bool
	stat = InsertNewDokumentMYSQLMongo(sqlData)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateDokumentMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.DokumentiItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = UpdateDokumentMYSQLMongo(sqlData)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteDokumentMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.DokumentiItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = DeleteDokumentiMYSQLMongo(sqlData)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//-------------------------------------------------------------------------------------
