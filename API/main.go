package main

import (
	"net/http"

	"github.com/go-martini/martini"
)

func Logger() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		c.Next()
	}
}

func ClassicA() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(Logger())
	m.Use(martini.Recovery())
	m.Use(martini.Static("public"))
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &martini.ClassicMartini{m, r}
}

func main() {
	InitMongoDB()
	m := ClassicA()
	m.Post("/get_article_list", GetArticleListMongo)
	m.Post("/insert_article", InsertNewArticleMongo)
	m.Post("/update_article", UpdateArticleMongo)
	m.Post("/delete_article", DeleteArticleMongo)

	m.Post("/get_komintent_list", GeKomintentiListMongo)
	m.Post("/insert_komintent", InsertNewKomintentiMongo)
	m.Post("/update_komintent", UpdateKomintentiMongo)
	m.Post("/delete_komintent", DeleteKomintentiMongo)

	m.Post("/get_dokumenti_list", GetDokumentListMongo)
	m.Post("/insert_dokumenti", InsertNewDokumentMongo)
	m.Post("/update_dokumenti", UpdateDokumentMongo)
	m.Post("/delete_dokumenti", DeleteDokumentMongo)

	//	m.Post("/get_dokumenti_detail_list", GetDokumentDetailListMongo)
	//	m.Post("/insert_dokumenti_detail", InsertNewDokumentDetailMongo)
	//	m.Post("/update_dokumenti_detail", UpdateDokumentDetailMongo)
	//	m.Post("/delete_dokumenti_detail", DeleteDokumentDetailMongo)

	http.ListenAndServe(":5002", m)
}
