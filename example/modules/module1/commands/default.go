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
	// stringArg, _ := c.GetString("string", "def")
	// fmt.Println(stringArg)
	
	// 整形
	// intArg, _ := c.GetInt("int", 200)
	// fmt.Println(intArg)

	// 布尔
	// boolArg, _ := c.GetBool("bool", false)
	// fmt.Println(boolArg)
	
	// 浮点
	// floatArg, _ := c.GetFloat("float", 1.1)
	// fmt.Println(floatArg)
	
	// SLICE
	// sliceArg, _ := c.GetStrings("slice", []string{"c", "d", "e"})
	// fmt.Println(sliceArg)
	
	// intsArg, _ := c.GetInts("ints", []int{2, 2, 2})
	// intsArg, _ := c.GetInts("ints")
	// fmt.Println(intsArg)
	
	// int8sArg, _ := c.GetInt8s("int8s", []int8{4, 4, 4})
	// fmt.Println(int8sArg)
	
	// f64s, _ := c.GetFloats("f64s", []float64{4.1, 4.1, 4.1})
	// fmt.Println(f64s)
	
	// bools, _ := c.GetBools("bools", []bool{true, true, false})
	// fmt.Println(bools)
	
	// 实用BEEGO读取CONF
	fmt.Println(beego.AppConfig.String("appname"))
	
	// 查看当前运行环境
	fmt.Println(beego.AppConfig.String("env"))
	
	for {}
}

func (c *MainCommand) GetOptions() []*console.Option {
	return []*console.Option {
		console.NewOption("string", "default", "字符串"),
		console.NewOption("int", 100, "整形"),
		console.NewOption("bool", true, "布尔"),
		console.NewOption("float", 10.1, "浮点"),
		console.NewOption("slice", []string{"a", "b", "c"}, "数组"),
		console.NewOption("ints", []int{1, 1, 1}, "数组"),
		console.NewOption("int8s", []int{1, 1, 1}, "数组"),
		console.NewOption("f64s", []float64{1, 1, 1}, "数组"),
		console.NewOption("bools", []bool{true, true, true}, "数组"),
	}
}
