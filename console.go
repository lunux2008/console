package console

import (
	"os"
	"fmt"
	"syscall"
)

var (
	OsArgs	  	 []string
	AppName		 = ""
	ModuleName   = ""
	RouteName	 = ""
	RunId		 = ""
	PidFile		 = ""
	SignalHooks  map[int]map[os.Signal][]func()
)

const (
	// PreSignal is the position to add filter before signal
	PreSignal = iota
	// FireSignal is the position to add filter fire signal
	FireSignal
	// PostSignal is the position to add filter after signal
	PostSignal
)

func init() {

	if syscall.Getppid() == 1 {
		syscall.Umask(0)
	}
	
	os.Args = GetNewArgs()

	CheckOsArgs()
	
	SetRunArgs()
	
	control := GetControl()

	if control != "console" && control != "start" {
		SwitchControl(control)
		os.Exit(0)
	}
	
	CheckPidFileExists(PidFile)
	
	if err := LoadConfig(ModuleName); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	InitSignalHooks()
	
	go SignalNotify()
}
