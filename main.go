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
	pingURL       string
	allowFrom     IPNetList
)

func startRedirector() {
	s := &http.Server{
		Addr:    redirBindAddr,
		Handler: http.RedirectHandler(redirURL, http.StatusTemporaryRedirect),
	}
	log.Fatal(s.ListenAndServe())
}

func main() {
	var err error

	flag.StringVar(&webBindAddr, "web-bind", ":8080", "web bind address")
	flag.StringVar(&webDir, "web-root", "static", "web root directory")
	flag.StringVar(&redirBindAddr, "redir-bind", "", "redirector bind address")
	flag.StringVar(&redirURL, "redir-url", "", "redirector target URL")
	flag.StringVar(&pingURL, "ping-url", "", "ping API URL")
	allowFromStr := flag.String("allow-from", "", "limit permitted IPs")
	flag.Parse()

	err = allowFrom.ParseArg(*allowFromStr)
	if err != nil {
		log.Fatal(err)
	}

	if redirBindAddr != "" && redirURL != "" {
		go startRedirector()
	}

	http.HandleFunc("/ping.php", pingHandler)
	http.HandleFunc("/empty.php", emptyHandler)
	http.HandleFunc("/garbage.php", garbageHandler)

	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(webBindAddr, allowNetworkWrapper(http.DefaultServeMux)))
}
