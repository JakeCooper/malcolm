package main

import (
	"fmt"
	"net/http"
	"os"

	malcolm "github.com/JakeCooper/malcolm/pkg/client"
)

// []Rule
// Rule {
//   ID: generated
//   From: (IP or addr)
//   To: (IP or addr)
//   Protocol: tcp/udp (default tcp)
// }

func rootGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func rootPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
}

func rootDelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
}

func fourohfour(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Invalid Request"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rootGet(w, r)
	case http.MethodPost:

	default:
		fourohfour(w, r)
	}
}

func bootstrap() {
	// Get shit from Redis and add listeners
}

func main() {
	// Create http service
	// /
	// GET - Returns all Ingress'
	// POST - Creates a new ingress
	// DELETE - Removes an ingress {""}

	PORT, exists := os.LookupEnv("PORT")
	if !exists {
		PORT = "1337"
	}

	http.HandleFunc("/rule", rootHandler)

	malcolm.New()

	fmt.Printf("Listening on port %s\n", PORT)

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}
