package routers

import (
	"github.com/lunux2008/console/example/modules/module1/commands"
	"github.com/lunux2008/console"
)

func init() {
	console.Router("module1/demo", &commands.MainCommand{})
}
