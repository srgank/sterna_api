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
	m := ClassicA()
	m.Post("/get_article_list", GetArticleList)
	m.Post("/insert_article", InsertNewArticle)
	m.Post("/update_article", UpdateArticle)
	m.Post("/delete_article", DeleteArticle)

	m.Post("/get_komintent_list", GetKomintentList)
	m.Post("/insert_komintent", InsertNewKomintent)
	m.Post("/update_komintent", UpdateKomintent)
	m.Post("/delete_komintent", DeleteKomintent)

	m.Post("/get_dokumenti_list", GetDokumentList)
	m.Post("/insert_dokumenti", InsertNewDokument)
	m.Post("/update_dokumenti", UpdateDokument)
	m.Post("/delete_dokumenti", DeleteDokument)

	m.Post("/get_dokumenti_detail_list", GetDokumentDetailList)
	m.Post("/insert_dokumenti_detail", InsertNewDokumentDetail)
	m.Post("/update_dokumenti_detail", UpdateDokumentDetail)
	m.Post("/delete_dokumenti_detail", DeleteDokumentDetail)

	http.ListenAndServe(":5002", m)
}
