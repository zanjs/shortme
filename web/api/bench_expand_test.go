package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// BenchmarkExpand may consume all open files in server. When running this
// test, you should configure max open files server can open. This can
// commonly be done through "ulimit -n 10000" shell command.
func BenchmarkExpand(b *testing.B) {
	b.StopTimer()
	shortReq, _ := json.Marshal(shortReq{LongURL: "http://www.google.com"})
	req, err := http.NewRequest(
		"POST",
		"http://127.0.0.1:3030/short",
		bytes.NewBuffer(shortReq))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		b.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		b.Fatalf("http response status: %v", resp.StatusCode)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var shortResp shortResp
	json.Unmarshal(body, &shortResp)
	shortURL := shortResp.ShortURL
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("%v", shortURL),
			nil,
		)
		if err != nil {
			b.Fatal(err)
		}

		transport := http.Transport{}
		resp, err := transport.RoundTrip(req)

		if err != nil {
			b.Fatal(err)
		}

		if resp.StatusCode != http.StatusTemporaryRedirect {
			b.Log(shortURL)
			b.Log(resp.Request.URL)
			b.Fatalf("http response status: %v", resp.StatusCode)
		}
	}
}
