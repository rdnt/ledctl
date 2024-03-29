package main

import (
	"fmt"
	"os"
	"os/signal"

	"ledctl3/internal/server"
	"ledctl3/internal/server/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	ctl, err := application.New(cfg)
	if err != nil {
		panic(err)
	}

	err = ctl.Start()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server started.")

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit

	fmt.Println("exit")
}
