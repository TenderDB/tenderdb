package handlers

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/tenderdb/tenderdb/domain"
	"github.com/tenderdb/tenderdb/valid"
)


func (h *handler) Filter(w http.ResponseWriter, r *http.Request) {

	okpd := r.URL.Query().Get("okpd")
	region := r.URL.Query().Get("region")
	innCustomer := r.URL.Query().Get("inncustomer")
	innSupplier := r.URL.Query().Get("innsupplier")

	if !(valid.Okpd(okpd) && valid.RegionNum(region) && valid.Inn(innCustomer) && valid.Inn(innSupplier)) {
		log.Println("Not Valid Request")
		http.Error(w, "Not Valid Request", http.StatusInternalServerError)
		return
	}
	
	var charts []domain.Dataset

	charts, err := h.repo.GetChartOnPrefixFilter(okpd, region, innCustomer, innSupplier)
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
