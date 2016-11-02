package console

import (
	"os"
	"fmt"
	"syscall"
	"github.com/astaxie/beego"
)

var stopping bool = false

func SwitchControl(control string) {
	switch control {
	default:
		Usage()
	case "reload":
		SendSignal(syscall.SIGUSR1)
		fmt.Println("Reload Success")
	case "grace":
		SendSignal(syscall.SIGUSR2)
		fmt.Println("Grace-Restart Success")
	case "restart":
		SendSignal(syscall.SIGHUP)
		fmt.Println("Restart Success")
	case "stop":
		SendSignal(syscall.SIGTERM)
		fmt.Println("Stop Success")
	}
}

func Interrupt() {
	if stopping {
		return 
	} else {
		stopping = true
	}

	RemoveFile(PidFile)
	os.Exit(0)
} 

func Terminate() {
	if stopping {
		return 
	} else {
		stopping = true
	}

	RemoveFile(PidFile)
	os.Exit(0)
} 

func Reload() (err error) {
	if stopping {
		return
	} else {
		stopping = true
	}

	if err = RemovePidFile(PidFile); err != nil {
		fmt.Println(err)
		return
	}

	if err = LoadConfig(ModuleName); err != nil {
		fmt.Println(err)
		return
	}
	
	AppName = beego.AppConfig.String("appname")
	PidFile = GetPidFile()
	
	CheckPidFileExists(PidFile)
	
	stopping = false
	
	return
}

func Restart() (err error) {
	if stopping {
		return 
	} else {
		stopping = true
	}

	if err = RemovePidFile(PidFile); err != nil {
		fmt.Println(err.Error())
		return
	}
	
	files := make([]*os.File, 3, 6)
	nullDev, err := os.OpenFile("/dev/null", 0, 0)
	if err != nil {
		return err
	}
	files[0], files[1], files[2] = nullDev, nullDev, nullDev

	dir, _   := os.Getwd()
	sysattrs := syscall.SysProcAttr{Setsid: true}
	attrs    := os.ProcAttr{Dir: dir, Env: os.Environ(), Files: files, Sys: &sysattrs}
	
	proc, err := os.StartProcess(os.Args[0], os.Args, &attrs)
	if err != nil {
		return
	}
	proc.Release()
	
	os.Exit(0)
	
	return
}

func Grace() (err error) {
	if stopping {
		return
	} else {
		stopping = true
	}
	
	for _, f := range SignalHooks[FireSignal][syscall.SIGUSR2]  {
		f()
	}

	stopping = false

	return
}