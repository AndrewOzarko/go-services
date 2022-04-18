package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nats-io/nats.go"
)

var encodedConn *nats.EncodedConn

func init() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatalln("can't connect to nats: ", err)
	}
	encodedConn, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	log.Println("Connected to nats")
}

func main() {

	defer encodedConn.Close()

	http.HandleFunc("/tasks/create", taskCreate)
	log.Println("Started :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func taskCreate(w http.ResponseWriter, r *http.Request) {

	type Card struct {
		Title       string
		Description string
	}

	card := &Card{Title: "tester21", Description: "new card"}
	encodedConn.Publish("foo", card)

	json.NewEncoder(w).Encode(card)
}
