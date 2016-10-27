package console

import (
	"os"
	"fmt"
	"github.com/astaxie/beego"
)

var (
	OsArgs	  	  []string
	AppName		  string = ""
	ModuleName    string = ""
	RouteName	  string = ""
	PidFilePrefix string = "/tmp/lunux2008.console."
)

func init() {

	CheckOsArgs()
	
	AppName    = beego.AppConfig.String("appname")
	RouteName  = os.Args[2]
	ModuleName = GetModuleName(RouteName)
	
	CheckPidFileExists()
	
	if err := LoadConfig(ModuleName); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	OsArgs  = os.Args
	os.Args = GetNewArgs()

	go SignalNotify()
}

