package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/prigotti/cargo/portsservice/server"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-signalCh

		cancel()
	}()

	c, err := server.DefaultConfiguration().Merge(server.GetEnvironmentConfiguration())
	if err != nil {
		fmt.Println(err)

		return -1
	}

	if _, err := server.NewServer(ctx, c); err != nil {
		fmt.Println(err)

		return -1
	}

	fmt.Println("server started")

	<-ctx.Done()

	fmt.Println("shutting down")

	return 0
}
