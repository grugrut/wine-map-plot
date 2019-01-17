package main

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Hello, world")
}

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}
