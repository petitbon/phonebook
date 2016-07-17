package main

import phonebook "github.com/petitbon/phonebook"

func main() {
	host := "localhost:9090"
	server := phonebook.NewPhonebookServer(host)
	server.Run()
}
