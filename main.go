package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ftl/digimodes/cw"
	"github.com/ftl/patrix/pa"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go handleCancelation(signals, cancel)

	for {
		play(ctx, "paris is the capital of france.")
	}
}

func play(ctx context.Context, text string) {
	modulator := cw.NewModulator(float64(700), 31)
	modulator.AbortWhenDone(ctx.Done())

	oscillator, err := pa.NewOscillator()
	if err != nil {
		log.Fatal(err)
	}

	oscillator.Modulator = modulator
	oscillator.Start()

	_, err = fmt.Fprintln(modulator, text)
	if err != nil {
		log.Fatal(err)
	}

	oscillator.Stop(ctx)
	modulator.Close()
	oscillator.Close()
}

func handleCancelation(signals <-chan os.Signal, cancel context.CancelFunc) {
	count := 0
	for {
		select {
		case <-signals:
			count++
			if count == 1 {
				cancel()
			} else {
				log.Fatal("hard shutdown")
			}
		}
	}
}
