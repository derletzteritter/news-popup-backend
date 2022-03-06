package main

import (
	"encoding/xml"
	"fmt"
	xml2json "github.com/basgys/goxml2json"
	"itschip/news-popup-backend/internal"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/nrkFeed", getNrkFeed)

	http.ListenAndServe(":5000", nil)
}

func getNrkFeed(writer http.ResponseWriter , response *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
    writer.Header().Set("Access-Control-Allow-Headers", "*")

	res, err := http.Get("https://www.nrk.no/nyheter/siste.rss")
	if err != nil {
		fmt.Println(err.Error())
	}

	xmlFeed := internal.NRK_Reader(res)
	b, err := xml.Marshal(&xmlFeed)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Convert from xml to json
	feed := strings.NewReader(string(b))
	json, err := xml2json.Convert(feed)

	writer.Write([]byte(json.String()))
}