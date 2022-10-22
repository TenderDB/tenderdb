package handlers
import (
	"net/http"
	//"log"
	"github.com/tenderdb/tenderdb/helpers"
	"github.com/tenderdb/tenderdb/valid"
)

type Content struct {
	Title          string
	Subtitle       string
	Authorized     bool
	Domain         string
	AuthGoogleLink string
	AuthYandexLink string
	Code           string
	Name           string
	MenuCode       string
	MenuName       string
	MenuList       map[string]string
	ShowChart      bool
}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	
	var content Content

	content.Title = "TenderDB"
	content.Domain = "tenderdb.ru"

	session, _ := h.cookie.Get(r ,"store")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

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
	okpd := r.URL.Query().Get("okpd")
	content.Code = okpd
	content.ShowChart = valid.OkpdForChart(okpd)
	content.Name, _ = h.repo.GetValue(okpd)
	content.Name = helpers.SeoMap(content.Name)

	//Make okpd short for Menu
	if len(okpd) > 8 {
		okpd = okpd[:8]
	}

	content.MenuCode = okpd
	content.MenuName, _ = h.repo.GetValue(okpd)
	content.MenuList, _ = h.repo.GetList(okpd)

	h.tmpl.ExecuteTemplate(w, "index.html", content)
}
