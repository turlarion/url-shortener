package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"turlarion.ru/url-shortener/internal/api"
	"turlarion.ru/url-shortener/internal/config"
)

func main() {

	cfgPath := flag.String("cfg", "config.yaml", "path to the config file")
	flag.Parse()

	cfg, err := config.FromFile(*cfgPath)

	if err != nil {
		fmt.Println("could not initialize config:" + err.Error())
	}

	r := chi.NewRouter()

	h := api.NewHandler(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password, cfg.Server.Host)

	r.Post("/save", h.Save)
	r.Get("/{id}", h.Get)

	r.Get("/hw", func(w http.ResponseWriter, r *http.Request) {
		println("HELLO WORLD!")
	})
	http.ListenAndServe(strconv.Itoa(cfg.Server.Port), r)
}
