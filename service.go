package main

import (
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	value := rand.Intn(100)
	status := http.StatusOK

	entry := log.WithFields(log.Fields{
		"path":   html.EscapeString(r.URL.Path),
		"method": r.Method,
		"app":    "service-a",
	})

	time.Sleep(time.Duration(value) * time.Millisecond)
	defer func(t time.Time, entry *log.Entry) {
		if status == http.StatusOK {
			entry.WithFields(log.Fields{
				"code": status,
				"tts":  fmt.Sprintf("%2.3f", time.Since(start).Seconds()),
			}).Info("All is well")
		} else {
			entry.WithFields(log.Fields{
				"code": status,
				"tts":  fmt.Sprintf("%2.3f", time.Since(start).Seconds()),
			}).Error("OMG Error!")
		}
	}(start, entry)
	if value < 25 {
		status = http.StatusInternalServerError
	}
	if status == http.StatusOK {
		fmt.Fprintf(w, ":-)")
	} else {
		http.Error(w, ":-(", http.StatusInternalServerError)
	}
}

func main() {
	log.SetFormatter(&log.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			parts := strings.Split(f.File, "/")
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", parts[len(parts)-1], f.Line)
		},
	})
	log.SetReportCaller(true)
	http.HandleFunc("/", defaultHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
