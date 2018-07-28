package main

import (
	"fmt"
	"net/http"
	"os"
	// "strings"
	// "bytes"
)

func main() {
	// http.HandleFunc("/", indexRoute)
	currDir, _ := os.Getwd()
	http.ErrNoLocation
	http.Handle("/", http.FileServer(http.Dir(currDir)))
	http.ListenAndServe(":1234", nil)
}
func indexRoute(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello bitches")
}
