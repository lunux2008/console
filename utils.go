package console

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"path/filepath"
	"github.com/astaxie/beego"
)

func Usage() {
	fmt.Println("Usage:")
	fmt.Printf("%s start   <route> [-arg1=your_arg1] [-args1=[arg1 arg2 arg3]]\n", os.Args[0])
	fmt.Printf("%s reload  <route> \n", os.Args[0])
	fmt.Printf("%s grace   <route> \n", os.Args[0])
	fmt.Printf("%s restart <route> \n", os.Args[0])
	fmt.Printf("%s stop    <route> \n", os.Args[0])
}

func CheckOsArgs() {
	if len(os.Args) < 3 {
		Usage()
		os.Exit(0)
	}
}

func GetControl() string {
	envArgs := strings.Split(os.Args[1], "::")
	
	if len(envArgs) == 0 {
		Usage()
		os.Exit(0)
	}
	
	if len(envArgs) == 2 {
		beego.BConfig.RunMode = envArgs[1]
	}
	
	return envArgs[0]
}

func LoadConfig(moduleName string) error {

	var configPath string
	
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

func SetRunArgs() {
	OsArgs  = os.Args
	runArgs := []string{}
	runArgs = append(runArgs, os.Args[:1]...)
	for _, v := range os.Args {
		if strings.Contains(v, "-runid=") {
			runArgs = append(runArgs, v)
		}
	}
	os.Args = runArgs

	var runid = flag.String("runid", "", "app run id")
	flag.Parse()
	
	os.Args = OsArgs
	
	AppName    = beego.AppConfig.String("appname")
	RouteName  = os.Args[2]
	ModuleName = GetModuleName(RouteName)
	RunId      = *runid
	PidFile    = GetPidFile()
}

func GetModuleName(routeName string) string {
	if strings.Contains(routeName, "/") {
		moduleArgs := strings.Split(routeName, "/")
		return moduleArgs[0]
	} else {
		return ""
	}
}