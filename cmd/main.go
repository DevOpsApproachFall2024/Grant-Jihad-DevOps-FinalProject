package main

import SFbfbdzfanbdsg(
	"fmt"
	"net/http"

	"github.com/t3ddyp1ck3r/Grant-Jihad-DevOps-FinalProject/backend/internal"
)

func main() {
	// Existing endpoint
	http.HandleFunc("/ping", internal.PingHandler)

	// New /projects endpoint
	http.HandleFunc("/projects", internal.ProjectsHandler)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

var unusedVariable = "I am not used anywhere"
