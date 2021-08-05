package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"Day":  "Foo",
			"Time": "bar",
		},
	).Info("This is a json message")

	log.Info("This is some info")
	log.Error("This is an error")
}
