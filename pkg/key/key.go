package key

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// ApiKey: Generates a device API Key using the provided credentials
func ApiKey(device, username, password string) string {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore invalid SSL certs
	}
	client := &http.Client{Transport: tr}

	path := "/api/?type=keygen&user=" + username + "&password=" + string(password) //path to retrieve API Key

	resp, err := client.Get("https://" + device + path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// reads the http body from the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// grep only the part with the API Key string
	re := regexp.MustCompile("<key>.*</key>")
	keyTag := re.FindString(string(body))
	APIKEY := strings.TrimRight(strings.TrimLeft(keyTag, "<key>"), "</key>")

	return APIKEY
}
