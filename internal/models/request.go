package models

type Request struct {
	Url     string `json:"url"`
	Timeout int    `json:"timeout"`
}
