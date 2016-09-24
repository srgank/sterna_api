package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
	"github.com/sterna_api/API/structs"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Target *string

var session *mgo.Session
var valid_session bool
var database_ string

func InitMongoDB() {
	url_ := "ds033056.mlab.com:33056"
	user_ := "user1"
	pass_ := "pass1"
	database_ = "vodno"
	session, _ = connectMongoDB(url_, user_, pass_, database_)
}

func connectMongoDB(url string, user string, pass string, db string) (*mgo.Session, *mgo.Database) {
	DialInfo := &mgo.DialInfo{
		Addrs:    []string{url},
		Timeout:  60 * time.Second,
		Database: db,
		Username: user,
		Password: pass,
	}

	session, err := mgo.DialWithInfo(DialInfo)
	if err != nil {
		valid_session = false
		panic(err)
	} else {
		valid_session = true
		session.SetMode(mgo.Monotonic, true)
		d := session.DB(db)
		return session, d
	}
}

//-------------Artikli----------------------------------------------------------------------------------------------------------------------

func GetArticleListMYSQLMongo(searchData structs.SearchByItem) (bool, string) {
	var output structs.Article
	var result []structs.ArticleItem
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("artikli")
	i_limit, _ := strconv.Atoi(searchData.Limit)
	i_offset, _ := strconv.Atoi(searchData.Offset)
	var query bson.M
	if searchData.SearchBy == "artikal" {
		if searchData.SearchName != "" {
			query = bson.M{"artikal": bson.M{"$regex": searchData.SearchName}}
		} else {
			query = bson.M{}
		}
	} else {
		if searchData.SearchName != "" {
			query = bson.M{"id": bson.M{"$regex": searchData.SearchName}}
		} else {
			query = bson.M{}
		}
	}

	err := c.Find(query).Limit(i_limit).Skip(i_offset).All(&result)
	if err != nil {
		return false, ""
	} else {
		output.Properties = result
		modules, _ := json.Marshal(output)
		modules_json := string(modules)
		fmt.Println(modules_json)
		return true, modules_json
	}
}

func DeleteArticleMYSQLMongo(data structs.ArticleItem) bool {
	//	var result SisDigitalSignage
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("artikli")

	err := c.Remove(bson.D{{"id", data.Id}})

	if err != nil {
		// panic(err)
		return false
	} else {
		return true
	}
}

func InsertNewArticleMYSQLMongo(data structs.ArticleItem) bool {

	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("artikli")
	data.Id = uuid.NewV4().String()
	err2 := c.Insert(data)

	if err2 != nil {
		return false
	} else {
		return true
	}
}

func UpdateArticleMYSQLMongo(inputData structs.ArticleItem) bool {
	fmt.Println("TTTTT", inputData)
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("artikli")

	colQuerier := bson.M{"id": inputData.Id}
	change := inputData
	err := c.Update(colQuerier, change)
	fmt.Println("errr", err)
	if err != nil {
		return false
	} else {
		return true
	}
}

//-------------Komintenti----------------------------------------------------------------------------------------------------------------------

func GetKomintentiListMYSQLMongo(searchData structs.SearchByItem) (bool, string) {
	var output structs.Komintenti
	var result []structs.KomintentiItem
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("komintenti")
	i_limit, _ := strconv.Atoi(searchData.Limit)
	i_offset, _ := strconv.Atoi(searchData.Offset)
	var query bson.M
	if searchData.SearchBy == "naziv" {
		if searchData.SearchName != "" {
			query = bson.M{"naziv": bson.M{"$regex": searchData.SearchName}}
		} else {
			query = bson.M{}
		}
	} else {
		if searchData.SearchName != "" {
			query = bson.M{"id": bson.M{"$regex": searchData.SearchName}}
		} else {
			query = bson.M{}
		}
	}

	err := c.Find(query).Limit(i_limit).Skip(i_offset).All(&result)
	if err != nil {
		return false, ""
	} else {
		output.Properties = result
		modules, _ := json.Marshal(output)
		modules_json := string(modules)
		fmt.Println(modules_json)
		return true, modules_json
	}
}

func DeleteKomintentiMYSQLMongo(data structs.KomintentiItem) bool {
	//	var result SisDigitalSignage
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("komintenti")

	err := c.Remove(bson.D{{"id", data.Id}})

	if err != nil {
		// panic(err)
		return false
	} else {
		return true
	}
}

func InsertNewKomintentiMYSQLMongo(data structs.KomintentiItem) bool {

	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("komintenti")
	data.Id = uuid.NewV4().String()
	err2 := c.Insert(data)

	if err2 != nil {
		return false
	} else {
		return true
	}
}

func UpdateKomintentiMYSQLMongo(inputData structs.KomintentiItem) bool {
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("komintenti")

	colQuerier := bson.M{"id": inputData.Id}
	change := inputData
	err := c.Update(colQuerier, change)
	fmt.Println("errr", err)
	if err != nil {
		return false
	} else {
		return true
	}
}

//-------------Dokumenti----------------------------------------------------------------------------------------------------------------------

func GetDokumentiListMYSQLMongo(searchData structs.SearchByItem) (bool, string) {
	var output structs.Dokumenti
	var result []structs.DokumentiItem
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("dokumenti")
	i_limit, _ := strconv.Atoi(searchData.Limit)
	i_offset, _ := strconv.Atoi(searchData.Offset)
	var query bson.M
	if searchData.SearchBy == "artikal" {
		if searchData.SearchName != "" {
			query = bson.M{"artikal": bson.M{"$regex": searchData.SearchName}}
		} else {
			query = bson.M{}
		}
	} else {
		if searchData.SearchName != "" {
			query = bson.M{"id": bson.M{"$regex": searchData.SearchName}}
		} else {
			query = bson.M{}
		}
	}

	err := c.Find(query).Sort("-DOCUMENT_ID").Limit(i_limit).Skip(i_offset).All(&result)
	if err != nil {
		return false, ""
	} else {
		output.Properties = result
		modules, _ := json.Marshal(output)
		modules_json := string(modules)
		fmt.Println(modules_json)
		return true, modules_json
	}
}

func DeleteDokumentiMYSQLMongo(data structs.DokumentiItem) bool {
	//	var result SisDigitalSignage
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("dokumenti")

	err := c.Remove(bson.D{{"TID", data.TID}})

	if err != nil {
		// panic(err)
		return false
	} else {
		return true
	}
}

func InsertNewDokumentMYSQLMongo(data structs.DokumentiItem) (bool, string) {
	var output structs.Dokumenti
	var result []structs.DokumentiItem
	var findRes structs.DokumentiItem

	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("dokumenti")
	data.TID = uuid.NewV4().String()
	err := c.Find(bson.M{"DOCUMENT_TIP": data.DOCUMENT_TIP}).Sort("-DOCUMENT_ID").One(&findRes)
	fmt.Println("findRes", findRes, "err", err)
	findRes.DOCUMENT_ID = findRes.DOCUMENT_ID + 1
	data.DOCUMENT_ID = findRes.DOCUMENT_ID
	result = append(result, findRes)
	err2 := c.Insert(data)

	if err2 != nil {
		return false, ""
	} else {
		output.Properties = result
		modules, _ := json.Marshal(output)
		modules_json := string(modules)
		fmt.Println("VVVVVVVVVVVVVVVVVVVVVVVVVVVVVVV", modules_json)
		return true, modules_json
	}
}

func UpdateDokumentMYSQLMongo(inputData structs.DokumentiItem) bool {
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("dokumenti")

	colQuerier := bson.M{"TID": inputData.TID}
	change := inputData
	err := c.Update(colQuerier, change)
	fmt.Println("errr", err)
	if err != nil {
		return false
	} else {
		return true
	}
}

//GetDokumentiDetailListMYSQLMongo
//InsertNewDokumentDetailMYSQLMongo
//UpdateDokumentDetailMYSQLMongo
//DeleteDokumentiDetailMYSQLMongo

//-------------Dokumenti----------------------------------------------------------------------------------------------------------------------

func GetDokumentiDetailListMYSQLMongo(searchData structs.SearchByItem) (bool, string) {
	var output structs.DokumentiDetail
	var result []structs.DokumentiDetailItem
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("dokumenti_detail")
	i_limit, _ := strconv.Atoi(searchData.Limit)
	i_offset, _ := strconv.Atoi(searchData.Offset)
	i_id, _ := strconv.Atoi(searchData.Dok_ID)
	i_tip, _ := strconv.Atoi(searchData.Dok_TIP)

	var query bson.M
	query = bson.M{"DOCUMENT_ID": i_id, "DOCUMENT_TIP": i_tip}

	err := c.Find(query).Sort("SUB_TID").Limit(i_limit).Skip(i_offset).All(&result)
	if err != nil {
		return false, ""
	} else {
		output.Properties = result
		modules, _ := json.Marshal(output)
		modules_json := string(modules)
		fmt.Println(modules_json)
		return true, modules_json
	}
}

func DeleteDokumentiDetailMYSQLMongo(data structs.DokumentiDetail) bool {
	//	var result SisDigitalSignage
	st := data.Properties
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("dokumenti_detail")

	x := c.Bulk()
	for i := range st {
		item := st[i]
		x.Remove(bson.D{{"TID", item.TID}})
	}
	_, err := x.Run()

	if err != nil {
		// panic(err)
		return false
	} else {
		return true
	}
}

func InsertNewDokumentDetailMYSQLMongo(data structs.DokumentiDetail) bool {

	st := data.Properties
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("dokumenti_detail")

	x := c.Bulk()
	var subTid int64
	subTid = 0
	for i := range st {
		item := st[i]
		item.TID = uuid.NewV4().String()
		item.SUB_TID = subTid
		x.Insert(item)
		subTid++
	}
	_, err := x.Run()

	if err != nil {
		return false
	} else {
		return true
	}
}

func UpdateDokumentDetailMYSQLMongo(data structs.DokumentiDetail) bool {
	st := data.Properties
	sessionCopy := session.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(database_).C("dokumenti_detail")

	x := c.Bulk()
	for i := range st {
		item := st[i]
		colQuerier := bson.M{"TID": item.TID}
		change := item
		x.Update(colQuerier, change)
	}
	_, err := x.Run()

	if err != nil {
		return false
	} else {
		return true
	}
}
