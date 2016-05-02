package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.RequestURI)
	})
}

func main() {
	var msg string
	var port int

	flag.StringVar(&msg, "msg", "", "the message to respond with")
	flag.IntVar(&port, "port", 8080, "port on which to bind")
	flag.Parse()

	bindAddr := fmt.Sprintf(":%v", port)

	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if msg != "" {
			fmt.Fprintf(w, "Message: %s\n\n", msg)
		}

		dump, _ := httputil.DumpRequest(r, true)
		fmt.Fprintf(w, "%s", dump)
	})

	log.Printf("Listening on port %d.", port)
	log.Fatalln(http.ListenAndServe(bindAddr, loggingMiddleware(r)))
}
