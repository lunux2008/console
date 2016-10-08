package commands

import (
	"github.com/lunux2008/console"
	"github.com/astaxie/beego"
	"fmt"
)

type MainCommand struct {
	console.Command
}

func (c *MainCommand) Handle() {
	
	// 字符
	stringArg, _ := c.GetString("string", "def")
	fmt.Println(stringArg)
	
	// 整形
	intArg, _ := c.GetInt("age", 200)
	fmt.Println(intArg)

	// 布尔
	boolArg, _ := c.GetBool("bool", false)
	fmt.Println(boolArg)
	
	// 浮点
	floatArg, _ := c.GetFloat("float", 1.1)
	fmt.Println(floatArg)
	
	// SLICE
	sliceArg := c.GetStrings("slice", []string{"c", "d", "e"})
	fmt.Println(sliceArg)
	
	// 实用BEEGO读取CONF
	fmt.Println(beego.AppConfig.String("appname"))
	
	// 查看当前运行环境
	fmt.Println(beego.AppConfig.String("env"))
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
