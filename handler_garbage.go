package main

import (
	"crypto/rand"
	"log"
	"net/http"
	"strconv"
)

var (
	garbageBytes [1048576]byte
)

func init() {
	rand.Read(garbageBytes[:])
}

func garbageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Description", "File Transfer")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=random.dat")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Add("Cache-Control", "post-check=0, pre-check=0")
	w.Header().Set("Pragma", "no-cache")
	w.WriteHeader(http.StatusOK)

	chunks := 4
	ckSize := r.FormValue("ckSize")
	if ckSize != "" {
		i, err := strconv.Atoi(ckSize)
		if err != nil {
			log.Printf("garbage.php: invalid ckSize: %s: %s", ckSize, err)
		} else if i < 0 || i > 100 {
			log.Printf("garbage.php: 0 > ckSize > 100: %d", chunks)
		} else {
			chunks = i
		}
	}

	for i := 0; i < chunks; i++ {
		_, err := w.Write(garbageBytes[:])
		if err != nil {
			log.Printf("garbage.php: %s", err)
			return
		}
	}
}
