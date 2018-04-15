package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var m module

	var cpp cppcheck
	m = cpp

	warns := m.parse("../cppcheck.log")

	url := "localhost:8080"

	j, _ := json.Marshal(warns)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "applcation/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
