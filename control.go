package console

import (
	"os"
	"os/exec"
	"fmt"
	"github.com/astaxie/beego"
)

var reloading  bool = false
var restarting bool = false
var stopping   bool = false

func Stop() {
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
	fmt.Println("Stop Success")
	os.Exit(0)
} 

func Reload() (err error) {
	if reloading {
		return
	} else {
		reloading = true
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
	PidFile = GetPidFile(AppName, RouteName)
	
	CheckPidFileExists(PidFile)
	
	fmt.Println("Reload Success")
	
	reloading = false
	
	return
}

func Restart() (err error) {
	if restarting {
		return 
	} else {
		restarting = true
	}

	return
}