package console

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"github.com/astaxie/beego"
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

func Stop() {
	RemoveFile(GetPidFile(AppName, ModuleName))
	os.Exit(0)
}

func Terminate() {
	RemoveFile(GetPidFile(AppName, ModuleName))
	fmt.Println("signal: terminated")
	os.Exit(0)
} 

func Reload() {
	pidFile := GetPidFile(AppName, ModuleName)
	if err := RemovePidFile(pidFile); err != nil {
		panic(err)
	}

	if err := LoadConfig(ModuleName); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	
	AppName = beego.AppConfig.String("appname")	
	
	CheckPidFileExists()
	
	fmt.Println("reload success")
}

func Restart() {
	fmt.Println("restart")
}