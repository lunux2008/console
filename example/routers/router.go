package routers

import (
	"github.com/lunux2008/console/example/commands"
	"github.com/lunux2008/console"
)

func init() {
    console.Router("demo", &commands.MainCommand{})
}
