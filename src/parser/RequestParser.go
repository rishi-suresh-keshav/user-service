package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseRequest(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error while reading request")
		return
	}

	if err := json.Unmarshal(body, x); err != nil {
		return
	}
}
