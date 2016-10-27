package console

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"path/filepath"
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

func CheckOsArgs() {
	if len(os.Args) < 3 {
		fmt.Println("Not Enough Args")
		os.Exit(0)
	}
	
	envArgs := strings.Split(os.Args[1], "::")
	
	if len(envArgs) == 0 || envArgs[0] != "console" {
		fmt.Println("Invalid Console Args")
		os.Exit(0)
	}
	
	if len(envArgs) == 2 {
		beego.BConfig.RunMode = envArgs[1]
	}
}

func CheckPidFileExists() {
	pidFile := GetPidFile(AppName, ModuleName)
	
	if FileExists(pidFile) {
		fmt.Printf("pid file: %s exists, application exit\n", pidFile)
		os.Exit(0)
	}
	
	// create pid file
	if err := CreatePidFile(pidFile); err != nil {
		panic(err)
	}
}

func LoadConfig(moduleName string) error {

	configPath := ""
	if moduleName == "" {
		configPath = "conf/app.conf"
	} else {
		configPath = "modules/" + moduleName + "/conf/app.conf"
	}

	if absConfigPath, err := filepath.Abs(configPath); err == nil {
		return beego.LoadAppConfig("ini", absConfigPath)
	} else {
		return err
	}
}

func GetNewArgs() []string {
	merge   := false
	newArgs := []string{}
	
	for _, v := range os.Args {
		if merge {
			newArgs[len(newArgs)-1] += (" " + v)
		} else {
			newArgs = append(newArgs, v)
		}
	
		if merge && v[0] != '-' && !strings.Contains(v, "=[") && strings.Contains(v, "]") {
			merge = false
		} else if v[0] == '-' && strings.Contains(v, "=[") && !strings.Contains(v, "]") {
			merge = true
		}
	}
	
	return newArgs
}

func GetModuleName(routeName string) string {
	if strings.Contains(routeName, "/") {
		moduleArgs := strings.Split(routeName, "/")
		return moduleArgs[0]
	} else {
		return ""
	}
}

func GetPidFile(appName, moduleName string) string {
	pidFile := PidFilePrefix
	
	if appName != "" {
		pidFile += (appName + ".")
	}
	
	if moduleName != "" {
		pidFile += (moduleName + ".")
	}
	
	return pidFile + "pid"
}

func CreatePidFile(pidFile string) error {
	pid := os.Getpid()
	return WriteFile(pidFile, strconv.Itoa(pid))
}

func RemovePidFile(pidFile string) error {
	return RemoveFile(pidFile)
}