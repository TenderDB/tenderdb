package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/tenderdb/tenderdb/valid"
)

func (h *handler) Autocomplete(w http.ResponseWriter, r *http.Request) {
	list := make(map[string]string)
	word := r.URL.Query().Get("word")

	if !valid.Rus(word) {
		//As it search in russian this call to not use latin letters and signs 
		list["Исключите цифры, знаки, латинские буквы"] = ""
		result, _ := json.Marshal(list)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	}

	list, err := h.repo.Autocomplete(word)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(list) == 0 {
		//Warning: No data for request 
		list["Нет соответствующих записей"] = ""
	}
	result, err := json.Marshal(list)
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
