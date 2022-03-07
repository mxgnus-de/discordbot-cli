package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetDataFromURL(module string, url string) []byte {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Hmmm, something went wrong (" + module + ")")
		os.Exit(1)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Hmmm, something went wrong (" + module + ")")
		os.Exit(1)
	}

	return body
}
