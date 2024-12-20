package main

import (
	"fmt"
	"net/http"

	"github.com/t3ddyp1ck3r/Grant-Jihad-DevOps-FinalProject/backend/internal"
)

func main() {
	http.HandleFunc("/ping", internal.PingHandler)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
