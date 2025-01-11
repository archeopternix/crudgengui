package main

import (
	. "crudgengui/internal"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelWarn)

	server := NewGuiServer()
	server.Init()
	server.StartServer(1323)
}
