package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	fsServe := http.FileServer(http.Dir("/Users/abtiwary/Development/Golang/GoThreeJS/serve"))
	fsStatic := http.FileServer(http.Dir("/Users/abtiwary/Development/Golang/GoThreeJS/serve/static"))

	http.Handle("/", fsServe)
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	serverIP := "0.0.0.0"
	serverPort := 8088
	log.WithFields(log.Fields{
		"server_ip": serverIP,
		"server_port": serverPort,
	}).Info("server is now listening")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", serverIP, serverPort), nil)
	if err != nil {
		log.WithError(err).Fatal("an error occurred")
	}
}
