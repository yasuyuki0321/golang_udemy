package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Get

func main() {
	base, _ := url.Parse("https://httpbin.org")
	// fmt.Println(base)

	ref, _ := url.Parse("/get")

	endpoint := base.ResolveReference(ref).String()
	// fmt.Printf("%T\n", endpoint)
	// fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Content-Type", `"applicaiton/json"`)

	// q := req.URL.Query()
	// q.Add("name", "test")

	// fmt.Println(q)
	// fmt.Println(q.Encode())

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
