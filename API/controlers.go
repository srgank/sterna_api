package main

import (
	"encoding/json"
	//	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sterna_api/API/structs"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

func Login(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	var sqlData structs.LoginDataRequest
	_ = json.Unmarshal(Respbody, &sqlData)
	//	fmt.Println("sqlData", sqlData)
	_, s := GetLogin(sqlData)

	//	fmt.Println(s)
	w.Write([]byte(s))
}

func GetArticleListMongo(w http.ResponseWriter, r *http.Request) {
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	Respbody, _ := ioutil.ReadAll(r.Body)
	//	fmt.Println("Respbody", Respbody)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)
	//	fmt.Println("sqlData", sqlData)
	_, s := GetArticleListMYSQLMongo(sqlData, dpx, usr)

	//	fmt.Println(s)
	w.Write([]byte(s))
}

func InsertNewArticleMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.ArticleItem
	_ = json.Unmarshal(Respbody, &sqlData)

	var stat bool
	stat = InsertNewArticleMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateArticleMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.ArticleItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = UpdateArticleMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteArticleMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.ArticleItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = DeleteArticleMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//-------------------------------------------------------------------------------------

func GeKomintentiListMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	//	fmt.Println("Respbody", Respbody)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)
	//	fmt.Println("sqlData", sqlData)
	_, s := GetKomintentiListMYSQLMongo(sqlData, dpx, usr)

	//	fmt.Println(s)
	w.Write([]byte(s))
}

func InsertNewKomintentiMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.KomintentiItem
	_ = json.Unmarshal(Respbody, &sqlData)

	var stat bool
	stat = InsertNewKomintentiMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateKomintentiMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.KomintentiItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = UpdateKomintentiMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteKomintentiMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.KomintentiItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = DeleteKomintentiMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//-------------------------------------------------------------------------------------

func GetDokumentListMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	//	fmt.Println("Respbody", Respbody)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)
	//	fmt.Println("sqlData", sqlData)
	_, s := GetDokumentiListMYSQLMongo(sqlData, dpx, usr)

	//	fmt.Println(s)
	w.Write([]byte(s))
}

func InsertNewDokumentMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.DokumentiItem
	_ = json.Unmarshal(Respbody, &sqlData)

	var stat bool
	var s string
	stat, s = InsertNewDokumentMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.Write([]byte(s))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateDokumentMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.DokumentiItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = UpdateDokumentMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteDokumentMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.DokumentiItem
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = DeleteDokumentiMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//-------------------------------------------------------------------------------------

func GetDokumentDetailListMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	//	fmt.Println("Respbody", Respbody)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)
	//	fmt.Println("sqlData", sqlData)
	_, s := GetDokumentiDetailListMYSQLMongo(sqlData, dpx, usr)

	//	fmt.Println(s)
	w.Write([]byte(s))
}

func InsertNewDokumentDetailMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.DokumentiDetail
	_ = json.Unmarshal(Respbody, &sqlData)

	var stat bool
	stat = InsertNewDokumentDetailMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func UpdateDokumentDetailMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.DokumentiDetail
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = UpdateDokumentDetailMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func DeleteDokumentDetailMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	var sqlData structs.DokumentiDetail
	_ = json.Unmarshal(Respbody, &sqlData)
	var stat bool
	stat = DeleteDokumentiDetailMYSQLMongo(sqlData, dpx, usr)
	if stat == true {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

//--------------------------------------------------------------------------------------

func GetAccountListMongo(w http.ResponseWriter, r *http.Request) {
	Respbody, _ := ioutil.ReadAll(r.Body)
	dpx := r.Header.Get("database_prefix")
	usr := r.Header.Get("user_name")

	//	fmt.Println("Respbody", Respbody)
	var sqlData structs.SearchByItem
	_ = json.Unmarshal(Respbody, &sqlData)
	//	fmt.Println("sqlData", sqlData)
	_, s := GetDokumentiDetailListMYSQLMongo(sqlData, dpx, usr)

	//	fmt.Println(s)
	w.Write([]byte(s))
}
