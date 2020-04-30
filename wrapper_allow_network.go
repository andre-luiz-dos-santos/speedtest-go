package main

import (
	"log"
	"net/http"
)

func allowNetworkWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowFrom.IsSet() && !allowFrom.ContainsTCPAddr(r.RemoteAddr) {
			message := "Access denied from " + r.RemoteAddr
			log.Print(message)
			http.Error(w, message, http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}
