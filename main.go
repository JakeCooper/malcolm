package main

import (
	"fmt"
	"html"
	"net/http"
	"os"

	malcolm "github.com/JakeCooper/malcolm/src"
)

// []Rule
// Rule {
//   ID: generated
//   From: (IP or addr)
//   To: (IP or addr)
//   Protocol: tcp/udp (default tcp)
// }

func rootGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func rootPost(w http.ResponseWriter, r *http.Request) {

}

func rootDelete(w http.ResponseWriter, r *http.Request) {
	// json.Unmarshal()
}

func fourohfour(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Invalid Request"))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rootGet(w, r)
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

	http.HandleFunc("/", rootHandler)

	malcolm.New()

	fmt.Printf("Listening on port %s\n", PORT)

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}
