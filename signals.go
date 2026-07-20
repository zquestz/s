package main

import (
	"os"
	"os/signal"
)

func setupSignalHandlers() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go interruptSignal(c)
}

func interruptSignal(c <-chan os.Signal) {
	<-c
	os.Exit(0)
}
