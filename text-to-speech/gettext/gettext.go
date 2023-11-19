package gettext

import (
	"io"
	"log"
	"net/http"
)

func Gettex() string {
	url := "http://goapi.apps.svc.cluster.local/books"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
