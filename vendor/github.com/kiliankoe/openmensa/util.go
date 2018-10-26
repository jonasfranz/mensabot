package openmensa

import (
	"io/ioutil"
	"net/http"
)

func get(url string) (response []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return respData, err
}
