package handlers
import (
	"net/http"
	"context"
	"strings"
	"golang.org/x/oauth2"
	"io"
	"io/ioutil"
	"encoding/json"
	"github.com/tenderdb/tenderdb/valid"
	"log"
)


// AuthHandler handles authentication of a user and initiates a session.
func (h *handler) CheckAuthResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		session, _ := h.cookie.Get(r ,"store")
		
		//Check state from OAuth Provider
		
		retrievedState := session.Values["state"]
		queryState := r.URL.Query().Get("state")
			
		if retrievedState != queryState {
			http.Error(w, "Wrong state", http.StatusInternalServerError)
			return
		}
	
		//Check code from OAuth Provider

		code := r.URL.Query().Get("code")

		var oauthProvider EnvAuth

		switch strings.TrimPrefix(r.URL.Path, "/auth/"){
		case "yandex":
			oauthProvider = h.yandexProvider
		case "google":
			oauthProvider = h.googleProvider
		default:	
			http.Error(w, "Page not found", 404)
			return
		}

		tok, err := oauthProvider.conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := oauthProvider.conf.Client(oauth2.NoContext, tok)

		userinfo, err := client.Get(oauthProvider.userDataUrl)
		if err != nil {
			log.Println(userinfo)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer userinfo.Body.Close()

		data, err := ioutil.ReadAll(userinfo.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		type UserData    struct {
			UserGoogleEmail       string `json:"email"`
			UserYandexEmail string `json:"default_email"`
		}

		var u UserData

		err = json.Unmarshal(data, &u) 
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		session, _ = h.cookie.Get(r ,"store")
		
		var email string

		switch strings.TrimPrefix(r.URL.Path, "/auth/"){
		case "yandex": email = u.UserYandexEmail
		case "google": email = u.UserGoogleEmail
		default:	
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w,"{'message': 'Please Login'}")
			return
		}

		if email == "" {
			http.Error(w, "No user data" , http.StatusInternalServerError)
			return
		}

		session.Values["user-id"] = email

		err = session.Save(r,w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = h.repo.CreateUser(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		next.ServeHTTP(w, r)
		return
	})

}
func (h *handler) CreateTestSession (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		session, _ := h.cookie.Get(r ,"store")
		email := "test@test.com"
		session.Values["user-id"] = email 

		err := session.Save(r,w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		err = h.repo.CreateUser(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = h.repo.CreateChartsUser(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		next.ServeHTTP(w, r)
		return
	})
}

func (h *handler) CheckAuth (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		session, err := h.cookie.Get(r ,"store")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		email, ok := session.Values["user-id"].(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w,"{'message': 'Please Login'}")
			return
		}
		
		if !valid.Email(email) {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w,"{'message': 'Not Valid Email'}")
		return
		}
		ctx := context.WithValue(r.Context(), "user-id", email )
		err = h.repo.UpdateUser(email, "enter")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func (h *handler) Exit (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := h.cookie.Get(r ,"store")
		delete(session.Values, "user-id")
		err = session.Save(r,w)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}	

		next.ServeHTTP(w, r)
		return

	})
}
