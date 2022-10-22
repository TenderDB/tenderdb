package handlers

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/tenderdb/tenderdb/domain"
	"github.com/tenderdb/tenderdb/valid"
)

func (h *handler) Fetch(w http.ResponseWriter, r *http.Request) {

	okpd := r.URL.Query().Get("okpd")

	if !valid.OkpdForChart(okpd) {
		log.Println("Not Valid Request")
		http.Error(w, "Not Valid Request", http.StatusInternalServerError)
		return
	}
	var charts []domain.Dataset

	charts, err := h.repo.GetChartOnPrefix(okpd)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// No data for request
	if len(charts) == 0 {
		charts = []domain.Dataset{{Name: "Нет данных", Data: map[string]string{"0": "0"}}}
	}
	result, err := json.Marshal(charts)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
	return
}
