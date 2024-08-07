package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"turlarion.ru/url-shortener/internal/models"
	"turlarion.ru/url-shortener/internal/services"
)

type handler struct {
	service    services.ShortenerService
	serverHost string
}

func NewHandler(redisHost string, redisPort int, redisPwd string, serverUrl string) handler {
	return handler{services.New(redisHost, redisPort, redisPwd), serverUrl}
}

func (h handler) Save(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(fmt.Errorf("could not parse request body"))
		return
	}

	newLink, err := h.service.SaveNewUrl(req.Url, req.Timeout)

	if err != nil {
		http.Error(w, "Could not create short url", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(models.NewUrl{ShortUrl: h.serverHost + "/" + newLink})

	if err != nil {
		http.Error(w, "Could not create short url", http.StatusInternalServerError)
		return
	}

	w.Write(response)

}

func (h handler) Redirect(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "linkId")
	if id == "" {
		http.Error(w, "link is bad", http.StatusBadRequest)
		return
	}

	response, err := h.service.GetFullLink(id)
	if err != nil {

		if err == redis.Nil {
			http.Error(w, "url not found", http.StatusBadRequest)
			return
		}

		http.Error(w, "Error occured", http.StatusInternalServerError)
		return
	}

	fmt.Println("redirecting to " + response)

	http.Redirect(w, r, response, http.StatusSeeOther)

}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "linkId")
	if id == "" {
		http.Error(w, "link is bad", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteShortLink(id)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
