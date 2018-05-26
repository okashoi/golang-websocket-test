package main

import (
	"github.com/trevex/golem"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", createRouter().Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createRouter() *golem.Router {
	r := golem.NewRouter()
	r.On("echo", echo)
	return r
}

func echo(conn *golem.Connection, data *message) {
	conn.Emit("echo", data)
}

type message struct {
	Msg string `json:"msg"`
}
