// Usage:
//    ./main -p="8181"
//
// Navigating to http://localhost:8181 will display the request that was made

package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	port := flag.String("p", "8180", "port to serve on")
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if str, err := httputil.DumpRequest(r, true); err == nil {
		w.Write(str)
	}
}
