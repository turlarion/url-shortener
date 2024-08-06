package api

import (
	"encoding/json"
	"net/http"

	"turlarion.ru/url-shortener/internal/models"
	"turlarion.ru/url-shortener/internal/services"
)

func Save(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	log.Println(fmt.Errorf("could not parse request body"))
	// 	return
	// }

	services.SaveNewUrl(req.Url, req.Timeout)

	w.WriteHeader(http.StatusPermanentRedirect)
	w.Header().Set("Location", "https://www.google.com")

}
