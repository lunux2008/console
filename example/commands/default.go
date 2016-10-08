package commands

import (
	"github.com/lunux2008/console"
	"fmt"
)

type MainCommand struct {
	console.Command
}

func (c *MainCommand) Handle() {

	stringArg, _ := c.GetString("string", "def")
	fmt.Println(stringArg)

	intArg, _ := c.GetInt("age", 200)
	fmt.Println(intArg)

	boolArg, _ := c.GetBool("bool", false)
	fmt.Println(boolArg)

	floatArg, _ := c.GetFloat("float", 1.1)
	fmt.Println(floatArg)
	
	sliceArg := c.GetStrings("slice", []string{"c", "d", "e"})
	fmt.Println(sliceArg)
}

func (c *MainCommand) GetOptions() []*console.Option {
	return []*console.Option {
		console.NewOption("string", "default", "字符串"),
		console.NewOption("int", 100, "整形"),
		console.NewOption("bool", true, "布尔"),
		console.NewOption("float", 10.1, "浮点"),
		console.NewOption("slice", []string{"a", "b", "c"}, "数组"),
	}
}
