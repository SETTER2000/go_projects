package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://lphp.ru")
	if err != nil {
		log.Fatalf("Encountered the following error: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
