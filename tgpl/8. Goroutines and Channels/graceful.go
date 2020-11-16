package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ch := make(chan os.Signal, 1)
	log.Println(1)
	SetNotify(ch)
	log.Println(2)

	sig := <- ch
	log.Println(sig)
}

func SetNotify(c chan<- os.Signal) {
	signal.Notify(c, syscall.SIGUSR2, syscall.SIGTERM, syscall.SIGINT)
}