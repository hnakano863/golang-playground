package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	resp, _ := http.Get("https://httpbin.org/get")
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	s := string(bytes)
	fmt.Println(s)
}
