package handlers

import (
	"golang.org/x/oauth2"
	"log"
	"html/template"
	"net/http"
	"github.com/tenderdb/tenderdb/repositories"
	"github.com/gorilla/sessions"
)

type EnvAuth struct{
	conf oauth2.Config 
	userDataUrl string
}

type handler struct {
	repo		*repositories.Repository
	tmpl		*template.Template
	cookie		*sessions.CookieStore
	FileServer	http.Handler
	yandexProvider	EnvAuth
	googleProvider	EnvAuth
}
	

func NewHandler(repository *repositories.Repository) *handler {
	return &handler{repo: repository}
}
func (h *handler) AddNewTemplate(dir string){
	tmpl , err := template.ParseGlob(dir)
	h.tmpl = tmpl
	if err != nil {
		log.Fatal(err)
	}
}
func (h *handler) AddNewSession(secret string) {
	h.cookie = sessions.NewCookieStore([]byte(secret))
}
func (h *handler) AddNewYandexProvider(id, secret, url string){
		h.yandexProvider = EnvAuth{
			conf: oauth2.Config{
				ClientID:     id,
				ClientSecret: secret,
				RedirectURL:  url,
				Scopes:       []string{"login:email"},
				//Endpoint:     yandex.Endpoint,
				Endpoint:     oauth2.Endpoint{
					AuthURL:  "https://oauth.yandex.ru/authorize",
					TokenURL: "https://oauth.yandex.ru/token",
				},
			},
			userDataUrl: "https://login.yandex.ru/info",
		}
}
func (h *handler) AddNewGoogleProvider(id, secret, url string){
		h.googleProvider = EnvAuth{
			conf: oauth2.Config{
				ClientID:     id,
				ClientSecret: secret,
				RedirectURL:  url,
				Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
				//Endpoint:     google.Endpoint,
				Endpoint:     oauth2.Endpoint{
					AuthURL:   "https://accounts.google.com/o/oauth2/auth",
					TokenURL:  "https://oauth2.googleapis.com/token",
					AuthStyle: oauth2.AuthStyleInParams,
				},
			},
			userDataUrl: "https://www.googleapis.com/oauth2/v3/userinfo",
		}
}
func (h *handler) AddNewFileServer(dir string){
		h.FileServer = http.FileServer(http.Dir(dir))
}
func (h *handler) RedirectBack(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w,r,"/", http.StatusFound)
}
