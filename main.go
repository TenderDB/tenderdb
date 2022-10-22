package main

import (
	"github.com/tenderdb/tenderdb/handlers"
	"github.com/tenderdb/tenderdb/config"
	"github.com/tenderdb/tenderdb/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"log"
)


func main() {
	conf := config.NewConfig()
	repo := repositories.NewRepository()

	repo.AddUsersDB(conf.UsersDB)
	repo.AddCartsDB(conf.CartsDB)
	repo.AddMinDB(conf.MinDB)
	repo.AddAutocompleteDB(conf.AutocompleteDB)
	repo.AddCodesDB(conf.CodesDB)
	repo.AddTreeDB(conf.TreeDB)
	repo.AddMaxDB(conf.MaxDB)
	repo.AddYears(conf.Years)
	repo.AddFZs(conf.FZs)
	
	h := handlers.NewHandler(repo)
	h.AddNewTemplate(conf.Templates)	
	h.AddNewSession(conf.Session)	
	h.AddNewFileServer(conf.Assets)	
	h.AddNewYandexProvider(conf.YandexID, conf.YandexSecret, conf.YandexRedirect)	
	h.AddNewGoogleProvider(conf.GoogleID, conf.GoogleSecret, conf.GoogleRedirect)	

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", h.Index)
	r.Get("/example", h.Example)

	r.Route("/testmode", func (r chi.Router){
		r.Use(middleware.BasicAuth("X", map[string]string{"admin":"pass"}))
		r.Use(h.CreateTestSession)
		r.Get("/", h.RedirectBack)
	})

	r.Route("/auth/{provider}", func (r chi.Router){
		r.Use(h.CheckAuthResponse)
		r.Get("/", h.RedirectBack)
	})
	r.Route("/exit", func (r chi.Router){
		r.Use(h.Exit)
		r.Get("/", h.RedirectBack)
	})
	r.Get("/autocomplete", h.Autocomplete)
	r.Get("/fetch", h.Fetch)
	r.Route("/authorized", func(r chi.Router) {
		r.Use(h.CheckAuth)
		r.Get("/filter", h.Filter)
		r.Get("/download", h.DownloadCsv)
		r.Route("/charts", func(r chi.Router) {
			r.Get("/", h.ReadCharts)
			r.Put("/", h.UpdateChart)
			r.Delete("/", h.DeleteChart)
		})
	})
	r.Get("/assets/{}", func(w http.ResponseWriter, req *http.Request){
		http.StripPrefix("/assets/", h.FileServer).ServeHTTP(w,req)
	}) 
	log.Fatal(http.ListenAndServe(":" + conf.Port, r))
}
