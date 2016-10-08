package console

import (
	"os"
	"strconv"
	"reflect"
	"fmt"
	"flag"
	"strings"
)

type Command struct {
	OptionValues map[string]*OptionValue
	commandName  string
}

type Option struct {
	Name  string
	Value interface{}
	Usage string
}

func NewOption(name string, value interface{}, usage string) *Option {
	return &Option{name, value, usage}
}

type OptionValue struct {
	CommandValue string	 
	DefaultValue interface{}
}

func NewOptionValue(commandValue string, defaultValue interface{}) *OptionValue {
	return &OptionValue{commandValue, defaultValue}
}

// CommandInterface is an interface to uniform all command handler.
type CommandInterface interface {
	Init(commandName string, options []*Option)
	Prepare()
	Handle()
	Finish()
	GetOptions() []*Option
}

// Init generates default values of command operations.
func (c *Command) Init(commandName string, options []*Option) {
	c.commandName = commandName
	c.ParseOptions(options)
}

// Prepare runs after Init before request function execution.
func (c *Command) Prepare() {}

// Finish runs after request function execution.
func (c *Command) Finish() {}

// define command options.
func (c *Command) GetOptions() ([]*Option) {
	panic("GetOptions Method Should Be rewrite")
	return nil
}

// Handle adds a request function to handle request.
func (c *Command) Handle() {
	panic("Handle Method Should Be rewrite")
}

// StopRun makes panic of USERSTOPRUN error and go to recover function if defined.
func (c *Command) StopRun() {
	panic("Stop Run")
}

// ParseOptions maps input data map to obj struct.
func (c *Command) ParseOptions(options []*Option) {
	
	c.OptionValues = make(map[string]*OptionValue)
	
	for _, option := range options {
		c.OptionValues[option.Name] = NewOptionValue("", option.Value)
		flag.StringVar(&c.OptionValues[option.Name].CommandValue, option.Name, "", option.Usage)
	}

	// remove command route
	os.Args = append(os.Args[:1], os.Args[3:]...)
	
	flag.Parse()
}

// GetCommandName gets the executing command name.
func (c *Command) GetCommandName() string {
	return c.commandName
}

// Query
func (c *Command) Query(key string) *OptionValue {
	if v, ok := c.OptionValues[key]; ok {
		return v
	}

	return nil
}

// GetString returns the input value by key string or the default value while it's present and input is blank
func (c *Command) GetString(key string, def ...string) (string, error) {
	v := c.Query(key)
	if v != nil && v.CommandValue != "" {
		return v.CommandValue, nil
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.(string), nil
	}

	return "", nil
}


// GetStrings returns the input string slice by key string or the default value while it's present and input is blank
// it's designed for multi-value input field such as checkbox(input[type=checkbox]), multi-selection.
func (c *Command) GetStrings(key string, def ...[]string) []string {
	var defv []string
	v := c.Query(key)
	
	if v != nil && v.CommandValue != "" {
		return strings.Split(strings.Trim(v.CommandValue, "[] "), " ")
	} else if len(def) > 0 {
		return def[0]
	} else if v.DefaultValue != nil {
		return v.DefaultValue.([]string)
	}

	return defv
}

// GetInt returns input as an int or the default value while it's present and input is blank
func (c *Command) GetInt(key string, def ...int) (int, error) {
	v := c.Query(key)
	
	if v != nil && len(v.CommandValue) > 0 {
		i64, err := strconv.ParseInt(v.CommandValue, 10, 10)
		return int(i64), err
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.(int), nil
	}

	return 0, nil
}

// GetInt8 return input as an int8 or the default value while it's present and input is blank
func (c *Command) GetInt8(key string, def ...int8) (int8, error) {
	v := c.Query(key)

	if v != nil && len(v.CommandValue) > 0 {
		i64, err := strconv.ParseInt(v.CommandValue, 10, 8)
		return int8(i64), err
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.(int8), nil
	}

	return 0, nil
}


// GetInt16 returns input as an int16 or the default value while it's present and input is blank
func (c *Command) GetInt16(key string, def ...int16) (int16, error) {
	v := c.Query(key)

	if v != nil && len(v.CommandValue) > 0 {
		i64, err := strconv.ParseInt(v.CommandValue, 10, 16)
		return int16(i64), err
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.(int16), nil
	}

	return 0, nil
}

// GetInt32 returns input as an int32 or the default value while it's present and input is blank
func (c *Command) GetInt32(key string, def ...int32) (int32, error) {
	v := c.Query(key)

	if v != nil && len(v.CommandValue) > 0 {
		i64, err := strconv.ParseInt(v.CommandValue, 10, 32)
		return int32(i64), err
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.(int32), nil
	}

	return 0, nil
}

// GetInt64 returns input value as int64 or the default value while it's present and input is blank.
func (c *Command) GetInt64(key string, def ...int64) (int64, error) {
	v := c.Query(key)

	if v != nil && len(v.CommandValue) > 0 {
		i64, err := strconv.ParseInt(v.CommandValue, 10, 64)
		return int64(i64), err
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.(int64), nil
	}

	return 0, nil
}

// GetBool returns input value as bool or the default value while it's present and input is blank.
func (c *Command) GetBool(key string, def ...bool) (bool, error) {
	v := c.Query(key)

	if v != nil && len(v.CommandValue) > 0 {
		return strconv.ParseBool(v.CommandValue)
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.(bool), nil
	}

	return false, nil
}

// GetFloat returns input value as float64 or the default value while it's present and input is blank.
func (c *Command) GetFloat(key string, def ...float64) (float64, error) {
	v := c.Query(key)

	if v != nil && len(v.CommandValue) > 0 {
		return strconv.ParseFloat(v.CommandValue, 64)
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		switch v.DefaultValue.(type) {
		case int:
			return float64(v.DefaultValue.(int)), nil
		case float64:
			return v.DefaultValue.(float64), nil
		default:
			return 0.0, nil
		}
	}

	return 0.0, nil
}

//
type commandInfo struct {
	pattern        string
	commandType	   reflect.Type
}

func (p *commandInfo) Handle() {

	vc := reflect.New(p.commandType)
	execCommand, ok := vc.Interface().(CommandInterface)
	
	if !ok {
		panic("Command Is Not CommandInterface")
	}
	
	execCommand.Init(p.pattern, execCommand.GetOptions())
	execCommand.Handle()
}

// CommandRegister containers registered router rules, command handlers and filters.
type CommandRegister struct {
	routers map[string]*commandInfo
}

// NewCommandRegister returns a new CommandRegister.
func NewCommandRegister() *CommandRegister {
	cr := &CommandRegister {
		routers: make(map[string]*commandInfo),
	}

	return cr
}

// Add command handler and pattern rules to CommandRegister.
// usage:
//	default methods is the same name as method
//	Add("report:daily", &ReportDailyCommand{})
func (p *CommandRegister) Add(pattern string, c CommandInterface) {

	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()

	route := &commandInfo{}
	route.pattern = pattern
	route.commandType = t
	
	p.routers[pattern] = route
}


// Dispatch Event
func (p *CommandRegister) Dispatch() {

	if len(os.Args) < 3 {
		fmt.Println("Not Enough Args")
		os.Exit(0)
	}
	
	if os.Args[1] != "console" {
		fmt.Println("Invalid Console Args")
		os.Exit(0)
	}

	c, ok := p.routers[os.Args[2]]
	if !ok {
		fmt.Println("Not Command Route Find For " + os.Args[2])
		os.Exit(0)
	}
	
	merge   := false
	newArgs := []string{}
	
	for _, v := range os.Args {
		if merge {
			newArgs[len(newArgs)-1] += (" " + v)
		} else {
			newArgs = append(newArgs, v)
		}
	
		if merge && v[0] != '-' && !strings.Contains(v, "=[") && strings.Contains(v, "]") {
			merge = false
		} else if v[0] == '-' && strings.Contains(v, "=[") && !strings.Contains(v, "]") {
			merge = true
		}
	}
	
	os.Args = newArgs

	c.Handle()
}