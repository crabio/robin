package helpers

import (
	// External

	"os"
	"os/signal"
	"syscall"
	// Internal
)

// Start waiting terminating signals
func WaitTermSignals() os.Signal {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGTERM)

	// Wait any terminate signal
	return <-signal_chan
}
