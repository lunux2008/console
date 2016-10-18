package routers

import (
	"github.com/lunux2008/console/example/commands"
	_ "github.com/lunux2008/console/example/modules/module1/routers"
	"github.com/lunux2008/console"
)

func init() {
    console.Router("demo", &commands.MainCommand{})
}
