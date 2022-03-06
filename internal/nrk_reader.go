package internal

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Feed struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Item []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	Date string `xml:"pubDate"`
}

func NRK_Reader(res *http.Response) Feed {
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var feed Feed

	e := xml.Unmarshal(b, &feed)
	if e != nil {
		fmt.Println(e.Error())
	}

	return feed
}