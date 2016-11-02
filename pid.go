package console

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func GetPidFile() string {
	pidFile := "/tmp/lunux2008.console."

	if AppName != "" {
		pidFile += (AppName + ".")
	}

	pidFile += strings.Replace(RouteName, "/", "_" , -1) + "."

	if RunId != "" {
		pidFile += (RunId + ".")
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

func CheckPidFileExists(pidFile string) {
	if FileExists(pidFile) {
		fmt.Printf("Pid File: %s Exists, Application Exit\n", pidFile)
		os.Exit(0)
	}
	
	// create pid file
	if err := CreatePidFile(pidFile); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}