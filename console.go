package console

import (
	"os"
	"fmt"
	"syscall"
	"github.com/astaxie/beego"
)

var (
	OsArgs	  	 []string
	AppName		 string = ""
	ModuleName   string = ""
	RouteName	 string = ""
	PidFile		 string = ""
)

func init() {

	CheckOsArgs()

	AppName    = beego.AppConfig.String("appname")
	RouteName  = os.Args[2]
	ModuleName = GetModuleName(RouteName)
	PidFile    = GetPidFile(AppName, RouteName)

	control := GetControl()
	if control != "console" {
		SwitchControl(control)
		os.Exit(0)
	}
	
	CheckPidFileExists(PidFile)
	
	if err := LoadConfig(ModuleName); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	OsArgs  = os.Args
	os.Args = GetNewArgs()

	go SignalNotify()
}

func SwitchControl(control string) {
	switch control {
	default:
		fmt.Println("Invalid Control Arg")
	case "reload":
		SendSignal(syscall.SIGUSR1)
	case "restart":
		SendSignal(syscall.SIGUSR2)
	case "stop":
		SendSignal(syscall.SIGTERM)
	}
}
