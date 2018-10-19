package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func BenchmarkShort(b *testing.B) {
	b.StopTimer()
	shortReq, _ := json.Marshal(shortReq{LongURL: "http://www.google.com"})
	b.StartTimer()

	for i := 0; i < b.N; i++ {
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

		resp.Body.Close()
	}
}
