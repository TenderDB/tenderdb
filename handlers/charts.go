package handlers

import (
	"log"
	"net/http"
	"github.com/tenderdb/tenderdb/valid"
	"time"
	"encoding/json"
	"strings"
)


func (h *handler) UpdateChart(w http.ResponseWriter, r *http.Request){
	
	email := r.Context().Value("user-id").(string)
	link := r.URL.RequestURI()
	paramStr := strings.Split(link,"?")[1]

	log.Println(paramStr)
	title := r.URL.Query().Get("title")

	if !(valid.Title(title) && valid.Link(paramStr)) {
		log.Println("Not Valid Request")
		http.Error(w, "Not Valid Request", http.StatusInternalServerError)
		return
	}

	err := h.repo.UpdateChart(email, time.Now().Format("2006-01-02 15:04:05")+" "+title, paramStr)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	//Msg for frontend: Record was saved
	msg := map[string]string{"message": "Запись " + title + " сохранена"}
	result, _ := json.Marshal(msg)
	w.Write(result)
	return

}


func (h *handler)DeleteChart(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("user-id").(string)
	title := r.URL.Query().Get("title")
	
	if !valid.Title(title) {
		log.Println("Not Valid Request")
		http.Error(w, "Not Valid Request", http.StatusInternalServerError)
	}

	err := h.repo.DeleteChart(email, title)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Msg for frontend: Record was deleted
	msg := map[string]string{"message": "Запись " + title + " удалена"}
	result, _ := json.Marshal(msg)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	return

}

func (h *handler)ReadCharts(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("user-id").(string)

	carts, err := h.repo.ReadCharts(email)
	
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(carts)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
	return
}
