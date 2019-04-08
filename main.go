package main

import (
	_ "endpoints"
	"log"
	"middleware"
	"mymux"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		middleware.CloseLogFile()
		log.Println("Bye bye")
		os.Exit(0)
	}()
	log.Println("Server up and running, port 8000")
	http.ListenAndServe(":8000", mymux.GetMux())
}
