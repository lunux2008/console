package commands

import (
	"github.com/lunux2008/console"
	"fmt"
)

type MainCommand struct {
	console.Command
}

func (c *MainCommand) Handle() {
	// fmt.Println("ok")
	name, _ := c.GetString("name", "ok")
	fmt.Println(name)
	age, _ := c.GetInt("age", 30)
	fmt.Println(age)
	high, _ := c.GetBool("bool")
	fmt.Println(high)
	fl, _ := c.GetFloat("float", 11)
	fmt.Println(fl)
	
	ss := c.GetStrings("ss", []string{"c", "d"})
	fmt.Println(ss)
}

func (c *MainCommand) GetOptions() []*console.Option {
	return []*console.Option {
		console.NewOption("name", "default", "姓名"),
		console.NewOption("age", 10, "年龄"),
		console.NewOption("bool", true, "BOOL"),
		console.NewOption("float", 10, "float"),
		console.NewOption("ss", []string{"a", "b"}, "ss"),
	}
}
