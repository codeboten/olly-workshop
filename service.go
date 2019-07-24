package main

import (
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	value := rand.Intn(100)
	entry := log.WithFields(log.Fields{
		"value": value,
		"path":  html.EscapeString(r.URL.Path),
	})
	if value < 25 {
		entry.WithField("code", http.StatusInternalServerError).Error("OMG Error!")
		http.Error(w, ":-(", http.StatusInternalServerError)
		return
	}
	time.Sleep(time.Millisecond * 10)
	entry.WithField("code", http.StatusOK).Info("All is well")
	fmt.Fprintf(w, ":-)")
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	http.HandleFunc("/", defaultHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
