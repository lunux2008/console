package routers

import (
	"github.com/lunux2008/console/example/commands"
	module1_commands "github.com/lunux2008/console/example/modules/module1/commands"
	"github.com/lunux2008/console"
)

func init() {
    console.Router("demo", &commands.MainCommand{})
	console.Router("module1/demo", &module1_commands.MainCommand{})
}
