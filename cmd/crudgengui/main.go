package main

import (
	. "crudgengui/internal"
)

func main() {
	server := NewGuiServer()
	server.Init()
	server.StartServer(1323)
}
