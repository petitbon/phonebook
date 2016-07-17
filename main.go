package main

import (
	"github.com/petitbon/phonebook/phonebook"
)

func main() {
	host := "localhost:9090"
	server := phonebook.NewPhonebookServer(host)
	server.Run()
}
