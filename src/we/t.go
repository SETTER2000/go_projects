package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://setter.auth0.com/oauth/token"

	payload := strings.NewReader("{\"client_id\":\"1kMLKoxmMtvvGQHeGm4LkpF3jN1R2tD1\",\"client_secret\":\"Dqkhl5upab2ZUHi0gbPqHinuDwmiIFjhIPy4-W5CP_nLJsc7yl_Egolyebo7QTw_\",\"audience\":\"angular\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")


	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}