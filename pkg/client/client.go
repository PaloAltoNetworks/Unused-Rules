package client

import (
	"crypto/tls"
	"log"
	"net/http"
)

func FwClient(url string) *http.Response {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore invalid SSL certs

	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	return resp

}
