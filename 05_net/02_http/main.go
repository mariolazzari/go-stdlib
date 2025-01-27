package main

import (
	"fmt"
	"net/http"
)

func main() {
	// client
	resp, err := http.Get("https://mariolazzari.it")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// server
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ciao Mario!")
}

// middleware
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		f(w, r)
	}
}
