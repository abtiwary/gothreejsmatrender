package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func main() {
	serverIP := flag.String("server_ip", "0.0.0.0", "the IP to serve on")
	serverPort := flag.Int("server_port", 8088, "the Port to listen on")
	dirToServe := flag.String("serve_dir", "./serve", "the directory to serve")
	flag.Parse()

	serverDir := *dirToServe
	staticDir := filepath.Join(serverDir, "/static/")
	fsServe := http.FileServer(http.Dir(serverDir))
	fsStatic := http.FileServer(http.Dir(staticDir))

	http.Handle("/", fsServe)
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	log.WithFields(log.Fields{
		"server_ip":   *serverIP,
		"server_port": *serverPort,
	}).Info("server is now listening")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", *serverIP, *serverPort), nil)
	if err != nil {
		log.WithError(err).Fatal("an error occurred")
	}
}
