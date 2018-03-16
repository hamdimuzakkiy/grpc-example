package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/plus", plus)

	fmt.Println(http.ListenAndServe(":9010", nil))
}

type PlusInput struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

func plus(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var p PlusInput
	json.Unmarshal(body, &p)

	fmt.Fprint(w, map[string]interface{}{
		"response": p.Number1+p.Number2,
	})
	return
}