package handlers

import (
	"net/http"
	"github.com/tenderdb/tenderdb/helpers"
	"log"
)

func (h *handler) Example(w http.ResponseWriter, r *http.Request) {
	

	type Content struct {
		Title          string
		Subtitle       string
		Authorized     bool
		Domain         string
		AuthGoogleLink string
		AuthYandexLink string
	}

	var content Content	

	content.Title = "TenderDB"
	content.Domain = "tenderdb.ru"
	
	session, err := h.cookie.Get(r ,"store")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	val := session.Values["user-id"]
	email, ok := val.(string)

	if !ok {
		state, _ := helpers.RandToken(32)
		session.Values["state"] = state
		session.Save(r, w)

		content.Subtitle = "Анализ торгов"
		content.Authorized = false
		content.AuthGoogleLink = h.googleProvider.conf.AuthCodeURL(state)
		content.AuthYandexLink = h.yandexProvider.conf.AuthCodeURL(state)
	} else {
		content.Subtitle = "Анализ торгов  для " + email
		content.Authorized = true

	}

	h.tmpl.ExecuteTemplate(w, "example.html", content)
}
