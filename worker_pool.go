package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Alienvault struct {
	Passive_dns []Domain `json:"Passive_dns"`
	Count       int      `json:"Count"`
}

type Domain struct {
	Address        string `json:"Address"`
	First          string `json:"First"`
	Last           string `json:"Last"`
	Hostname       string `json:"Hostname"`
	Record_type    string `json:"Record_type"`
	Indicator_link string `json:"Indicator_link"`
	Flag_url       string `json:"Flag_url"`
	Flag_title     string `json:"Flag_title"`
	Asset_type     string `json:"Asset_type"`
	Asn            string `json:"Asn"`
}

func Worker(id int, jobs <-chan string, results chan<- []Domain) {

	for job := range jobs {

		r, err := http.Get(fmt.Sprintf("https://otx.alienvault.com/api/v1/indicators/domain/%s/passive_dns", job))

		if err != nil {
			fmt.Println(err)
		}

		url_byte, _ := ioutil.ReadAll(r.Body)

		var Src Alienvault

		json.Unmarshal(url_byte, &Src)

		results <- Src.Passive_dns

	}
}

func main() {

	jobs := make(chan string, 100)

	results := make(chan []Domain, 100)

	for w := 1; w <= 3; w++ {

		go Worker(w, jobs, results)
	}

	joblist := []string{"tesla.com", "apple.com", "facebook.com", "google.com"}

	for index, _ := range joblist {
		jobs <- joblist[index]
	}

	for i := 1; i <= len(joblist); i++ {
		data := <-results

		for index, _ := range data {
			fmt.Println(data[index].Hostname)
		}
	}
}
