package main

import (
	"fmt"
	"net/url"
)

func main() {

	// [protocol://][userinfo@]host[:port][/path][?query][#fragment] -> userinfo & port are optional

	rawUrl := "https://example.com:8080/path?query=param#fragment"

	parsedUrl, _ := url.Parse(rawUrl)
	fmt.Println("Scheme :", parsedUrl.Scheme)
	fmt.Println("Host :", parsedUrl.Host)
	fmt.Println("Port :", parsedUrl.Port())
	fmt.Println("Path :", parsedUrl.Path)
	fmt.Println("Queryparams :", parsedUrl.RawQuery)
	fmt.Println("Fragments :", parsedUrl.Fragment)

	rawUrl1 := "https://example.com/path?name=john&age=30"
	parsedUrl1, _ := url.Parse(rawUrl1)
	queryParams := parsedUrl1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name :", queryParams.Get("name"))
	fmt.Println("Age :", queryParams.Get("age"))

	//building URL
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/axs",
	}
	query := baseUrl.Query()
	query.Set("name", "abinash")
	baseUrl.RawQuery = query.Encode()
	fmt.Println(baseUrl)

	// Another way to build a URL
	values := url.Values{}

	//k-v pairs to the values object
	values.Add("name", "Howard")
	values.Add("age", "30")
	values.Add("role", "SE")
	//Encode
	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)
	//Build a URL
	baseUrl1 := "https://example.com/search"
	fullUrl := baseUrl1 + "?" + encodedQuery
	fmt.Println(fullUrl)
}
