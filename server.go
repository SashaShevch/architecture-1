package main

import (
	"time"
	"net/http"
	"encoding/json"
)

type Message struct {
	Time string `json:"time"`
}


func main() {
	http.HandleFunc("/time", func (w http.ResponseWriter, req *http.Request) {
		t := time.Now();
		m := Message{ t.Format(time.RFC3339) }
		b, _ := json.Marshal(m)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	http.ListenAndServe(":8795", nil)
}