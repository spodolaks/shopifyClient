package network

import (
	"net/http"
)

var Client *http.Client

func init() {
	Client = &http.Client{}
}

func Initialize() {
	Client = &http.Client{}
}