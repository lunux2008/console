package console

import (
	"os"
	"os/signal"
	"syscall"
)

func SignalNotify() {
	sigs:= make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGTERM)

	for {
		msg := <-sigs

		switch msg {
		default:
		case syscall.SIGINT:
			Stop()
		case syscall.SIGTERM:
			Terminate()
		case syscall.SIGUSR1:
			Reload()
		case syscall.SIGUSR2:
			Restart()
		}
	}
}

// todo 提供HOOKS