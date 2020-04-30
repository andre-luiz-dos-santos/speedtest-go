package main

import "net/http"

func allowNetworkWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowFrom.IsSet() && !allowFrom.ContainsTCPAddr(r.RemoteAddr) {
			http.Error(w, "Access denied from "+r.RemoteAddr, http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}
