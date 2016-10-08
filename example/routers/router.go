package routers

import (
	"home/luxu-so/test/t3/commands"
	"github.com/lunux2008/console"
)

func init() {
    console.Router("demo", &commands.MainCommand{})
}
