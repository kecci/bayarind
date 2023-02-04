package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/generate-authcode", func(w http.ResponseWriter, r *http.Request) {
		// params
		transactionNo := r.URL.Query().Get("transactionNo")
		transactionAmount := r.URL.Query().Get("transactionAmount")
		channelId := r.URL.Query().Get("channelId")
		secretKey := r.URL.Query().Get("secretKey")

		// transactionNo+transactionAmount+channelId+secretKey
		authCode := transactionNo + transactionAmount + channelId + secretKey
		
		// Hash the authCode
		hashed := sha256.Sum256([]byte(authCode))
		
		w.Write([]byte(fmt.Sprintf("%x", hashed[:])))
	})

	http.ListenAndServe(":8080", mux)
}
