package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

/*
Time function
is a middleware to measure execution times
will be deleted in production
*/
var f *os.File
var count int = 0

func init() {
	var err error
	f, err = os.OpenFile("metrics.urlencoded", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalln(err)
	}
}

func Time() Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer stopTimer(start, w, r)
			h.ServeHTTP(w, r)
		})
	}
}

func stopTimer(start time.Time, w http.ResponseWriter, r *http.Request) {
	total := time.Since(start)
	writeToFile(total)
	log.Println(total, r.URL.String())
}

func writeToFile(total time.Duration) {
	n, err := f.WriteString(total.String() + "\n")
	if err != nil {
		log.Println(err, n)
	}
}

func CloseLogFile() {
	log.Println("Closing file")
	f.Close()
}
