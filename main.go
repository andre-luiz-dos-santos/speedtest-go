package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	webBindAddr   string
	webDir        string
	redirBindAddr string
	redirURL      string
)

func startRedirector() {
	s := &http.Server{
		Addr:    redirBindAddr,
		Handler: http.RedirectHandler(redirURL, http.StatusPermanentRedirect),
	}
	log.Fatal(s.ListenAndServe())
}

func main() {
	flag.StringVar(&webBindAddr, "web-bind", ":8080", "web bind address")
	flag.StringVar(&webDir, "web-root", "static", "web root directory")
	flag.StringVar(&redirBindAddr, "redir-bind", "", "redirector bind address")
	flag.StringVar(&redirURL, "redir-url", "", "redirector target URL")
	flag.Parse()

	if redirBindAddr != "" && redirURL != "" {
		go startRedirector()
	}

	http.HandleFunc("/empty.php", emptyHandler)
	http.HandleFunc("/garbage.php", garbageHandler)

	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(webBindAddr, nil))
}
