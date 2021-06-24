package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//RSS struct
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

//Channel struct
type Channel struct {
	Title    string `xml:"title"`
	ItemList []Item `xml:"item"`
}

//Item struct
type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

//News struct
type News struct {
	Headline string `xml:"news_item_title"`
}

func main() {
	var r RSS

	data := readGoogleTrends(getGoogleTrends())
	err := xml.Unmarshal([]byte(data), &r)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(r.Channel.ItemList[0].Title)
	fmt.Println(r.Channel.ItemList[0].Traffic)
	fmt.Println(r.Channel.ItemList[0].NewsItems[0].Headline)

}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}

func readGoogleTrends(*http.Response) []byte {
	resp := getGoogleTrends()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(string(data))
	return data
}
