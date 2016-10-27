package console

import (
	"os"
	"os/signal"
	"syscall"
	"strconv"
)

func SignalNotify() {
	sigs:= make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGUSR1, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

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
		case syscall.SIGHUP:
			Restart()
		}
	}
}

func SendSignal(sig syscall.Signal) {
	var err error
	var pid int64
	var process *os.Process
	
	content := ReadFile(PidFile)
	if pid, err = strconv.ParseInt(content, 10, 32); err != nil {
		panic(err)
	}
	
	if process, err = os.FindProcess(int(pid)); err != nil {
		panic(err)
	}
	
	if err = process.Signal(sig); err != nil {
		panic(err)
	}
}

// todo 提供HOOKS