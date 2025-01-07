package main

import (
	. "crudgengui/internal"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	server := NewGuiServer()
	server.Init()
	server.StartServer(1323)
}
