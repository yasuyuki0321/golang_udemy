package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// httpクライアント

func main() {

	// GET
	// res, _ := http.Get("https://example.com")

	// fmt.Println(res.Status)
	// fmt.Println(res.Proto)
	// fmt.Println(res.Header["Date"])
	// fmt.Println(res.Header["Content-Type"])
	// fmt.Println(res.Request.Method)
	// fmt.Println(res.Request.URL)

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

	// POST
	vs := url.Values{}

	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs)

	res, err := http.PostForm("https://example.com", vs)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

}
