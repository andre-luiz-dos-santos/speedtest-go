package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	webDir      string
	bindAddress string
)

func main() {
	flag.StringVar(&bindAddress, "bind", ":8080", "bind address")
	flag.StringVar(&webDir, "root", "static", "root web directory")
	flag.Parse()

	http.HandleFunc("/empty.php", emptyHandler)
	http.HandleFunc("/garbage.php", garbageHandler)

	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(bindAddress, nil))
}
