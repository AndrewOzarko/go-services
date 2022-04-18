package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

type Card struct {
	Title       string
	Description string
}

var list []*Card

var encodedConn *nats.EncodedConn

func init() {
	list = append(list, &Card{Title: "test", Description: "test"})

	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatalln("can't connect to nats: ", err)
	}
	encodedConn, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	log.Println("Connected to nats")
}

func main() {
	defer encodedConn.Close()

	encodedConn.Subscribe("foo", func(c *Card) {
		log.Println("it is work")
		list = append(list, c)
	})

	http.HandleFunc("/list", mainHandler)
	log.Println("Started :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(list)
}
