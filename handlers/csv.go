package handlers

import (
	"fmt"
	"log"
	"net/http"
	"github.com/tenderdb/tenderdb/valid"
	"time"
	"encoding/json"
)

func (h *handler) DownloadCsv(w http.ResponseWriter, r *http.Request) {
	
	email := r.Context().Value("user-id").(string)
	
	lastDownloadStr, err := h.repo.CheckUserAction(email, "download")

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lastDownloadTime, err := time.Parse(time.RFC3339, lastDownloadStr)
	log.Println(lastDownloadTime) 


	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	nextTime := lastDownloadTime.Add(time.Second * 15)

	if time.Now().Before(nextTime) {
		//Msg for frontend: Limit one downloading per 15 sec
		message := fmt.Sprintf("Установлено ограничение один раз в 15 секунд. Попробуйте через %.0f секунд", nextTime.Sub(time.Now()).Seconds())

		msg := map[string]string{"message": message}
		result, _ := json.Marshal(msg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}

	okpd := r.URL.Query().Get("okpd")
	region := r.URL.Query().Get("region")
	innCustomer := r.URL.Query().Get("inncustomer")
	innSupplier := r.URL.Query().Get("innsupplier")
	title := r.URL.Query().Get("title")

	if !(valid.Okpd(okpd) && valid.RegionNum(region) && valid.Inn(innCustomer) && valid.Inn(innSupplier) && valid.Title(title)) {
		log.Println(err)
		http.Error(w, "Not Valid Request", http.StatusInternalServerError)
		return
	}

	byteFile, err := h.repo.CsvDownload(okpd, region, innCustomer, innSupplier)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.repo.UpdateUser(email, "download")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(byteFile)

}
