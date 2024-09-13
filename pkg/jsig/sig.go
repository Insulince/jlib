package jsig

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var defaultSigs = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
}

func Trap(sigs ...os.Signal) chan struct{} {
	if len(sigs) == 0 {
		sigs = defaultSigs
	}

	done := make(chan struct{})

	sigLis := make(chan os.Signal, 1)
	signal.Notify(sigLis, sigs...)

	go func() {
		sig := <-sigLis
		log.Printf("Received signal %q, exiting...\n", sig)
		close(done)
	}()

	return done
}
