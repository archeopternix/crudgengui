package main

import (
	. "crudgengui/internal"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelInfo)

	server := NewGuiServer()
	server.Init()
	server.StartServer(1323)
}
