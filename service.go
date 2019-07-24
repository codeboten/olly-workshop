package main

import (
	"errors"
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func doStuff() (int, error) {
	value := rand.Intn(100)
	var err error
	time.Sleep(time.Duration(value) * time.Millisecond)
	if value < 25 {
		err = errors.New("below threshold")
	} else {
		err = nil
	}
	return value, err
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	status := http.StatusOK

	entry := log.WithFields(log.Fields{
		"path":   html.EscapeString(r.URL.Path),
		"method": r.Method,
		"app":    "service-a",
	})

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
	_, err := doStuff()
	if err == nil {
		fmt.Fprintf(w, ":-)")
	} else {
		status = http.StatusInternalServerError
		http.Error(w, ":-(", status)
	}
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			parts := strings.Split(f.File, "/")
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", parts[len(parts)-1], f.Line)
		},
		DisableColors: true,
	})
	log.SetReportCaller(true)
	http.HandleFunc("/", defaultHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
