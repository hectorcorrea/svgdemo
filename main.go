package main

import (
	"fmt"
	"log"
	"net/http"
)

var address = "localhost:9001"

func home(resp http.ResponseWriter, req *http.Request) {
	log.Printf("Home: %v", req.URL)
	html := "<html><body>"
	html += "<p>Sample web server for static files.</p>"
	html += "<p>Serves anything under "
	html += "<a href=http://" + address + "/public/>http://" + address + "/public/</a></p>"
	html += "</body></html>"
	fmt.Fprint(resp, html)
}

func main() {
	log.Printf("Listening for requests at: http://%s", address)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", home)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Failed to start the web server: ", err)
	}
}
