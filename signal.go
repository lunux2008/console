package console

import (
	"os"
	"fmt"
	"os/signal"
	"syscall"
	"strconv"
)

func SignalNotify() {
	sigs:= make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	for {
		sig := <-sigs
		RunSignalHooks(PreSignal, sig)
		switch sig {
		default:
		case syscall.SIGHUP:
			Restart()
		case syscall.SIGINT:
			Interrupt()
		case syscall.SIGTERM:
			Terminate()
		case syscall.SIGUSR1:
			Reload()
		case syscall.SIGUSR2:
			Grace()
		}
		RunSignalHooks(PostSignal, sig)
	}
}

func SendSignal(sig syscall.Signal) {
	var err error
	var pid int64
	var process *os.Process
	
	content, err := ReadFile(PidFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	
	if pid, err = strconv.ParseInt(content, 10, 32); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	
	if process, err = os.FindProcess(int(pid)); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	
	if err = process.Signal(sig); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func InitSignalHooks() {
	SignalHooks = map[int]map[os.Signal][]func(){
		PreSignal: {
			syscall.SIGHUP:  {},
			syscall.SIGINT:  {},
			syscall.SIGTERM: {},
			syscall.SIGUSR1: {},
			syscall.SIGUSR2: {},
		},
		FireSignal: {
			syscall.SIGHUP:  {},
			syscall.SIGINT:  {},
			syscall.SIGTERM: {},
			syscall.SIGUSR1: {},
			syscall.SIGUSR2: {},
		},
		PostSignal: {
			syscall.SIGHUP:  {},
			syscall.SIGINT:  {},
			syscall.SIGTERM: {},
			syscall.SIGUSR1: {},
			syscall.SIGUSR2: {},
		},
	}
}

func RunSignalHooks(ppFlag int, sig os.Signal) {

	if _, notSet := SignalHooks[ppFlag][sig]; !notSet {
		return
	}
	for _, f := range SignalHooks[ppFlag][sig] {
		f()
	}
	return
}

func AddPreReloadHook(f func()) {
	SignalHooks[PreSignal][syscall.SIGUSR1] = append(SignalHooks[PreSignal][syscall.SIGUSR1], f)
}

func AddPostReloadHook(f func()) {
	SignalHooks[PostSignal][syscall.SIGUSR1] = append(SignalHooks[PostSignal][syscall.SIGUSR1], f)
}

func AddPreRestartHook(f func()) {
	SignalHooks[PreSignal][syscall.SIGHUP] = append(SignalHooks[PreSignal][syscall.SIGHUP], f)
}

func AddPostRestartHook(f func()) {
	SignalHooks[PostSignal][syscall.SIGHUP] = append(SignalHooks[PreSignal][syscall.SIGHUP], f)
}

func AddPreInterruptHook(f func()) {
	SignalHooks[PreSignal][syscall.SIGINT] = append(SignalHooks[PreSignal][syscall.SIGINT], f)
}

func AddPostInterruptHook(f func()) {
	SignalHooks[PostSignal][syscall.SIGINT] = append(SignalHooks[PreSignal][syscall.SIGINT], f)
}

func AddPreTerminateHook(f func()) {
	SignalHooks[PreSignal][syscall.SIGTERM] = append(SignalHooks[PreSignal][syscall.SIGTERM], f)
}

func AddPostTerminateHook(f func()) {
	SignalHooks[PostSignal][syscall.SIGTERM] = append(SignalHooks[PreSignal][syscall.SIGTERM], f)
}

func AddPreGraceHook(f func()) {
	SignalHooks[PreSignal][syscall.SIGUSR2] = append(SignalHooks[PreSignal][syscall.SIGUSR2], f)
}

func AddFireGraceHook(f func()) {
	SignalHooks[FireSignal][syscall.SIGUSR2] = append(SignalHooks[FireSignal][syscall.SIGUSR2], f)
}

func AddPostGraceHook(f func()) {
	SignalHooks[PostSignal][syscall.SIGUSR2] = append(SignalHooks[PostSignal][syscall.SIGUSR2], f)
}
