package console

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func GetPidFile(appName, routeName string) string {
	pidFile := "/tmp/lunux2008.console."
	
	if appName != "" {
		pidFile += (appName + ".")
	}
	
	return pidFile + strings.Replace(routeName, "/", "_" , -1) + ".pid"
}

func CreatePidFile(pidFile string) error {
	pid := os.Getpid()
	return WriteFile(pidFile, strconv.Itoa(pid))
}

func RemovePidFile(pidFile string) error {
	return RemoveFile(pidFile)
}

func CheckPidFileExists(pidFile string) {
	if FileExists(pidFile) {
		fmt.Printf("pid file: %s exists, application exit\n", pidFile)
		os.Exit(0)
	}
	
	// create pid file
	if err := CreatePidFile(pidFile); err != nil {
		panic(err)
	}
}