package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		output, err := makeRequest(url)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		_, err = fmt.Fprintf(w, "%s", output)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func makeRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body: %w", err)
	}
	return body, nil
}
