package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	value := rand.Intn(100)
	if value < 25 {
		log.WithFields(log.Fields{
			"value": value,
		}).Error("OMG Error!")
		http.Error(w, ":-(", http.StatusInternalServerError)
		return
	}
	time.Sleep(time.Millisecond * 10)
	log.WithFields(log.Fields{
		"value": value,
	}).Info("All is well")
	fmt.Fprintf(w, ":-)")
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	http.HandleFunc("/", defaultHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
