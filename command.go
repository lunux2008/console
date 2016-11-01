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
	fmt.Println("GetOptions Method Should Be rewrite")
	os.Exit(0)
	
	return nil
}

// Handle adds a request function to handle request.
func (c *Command) Handle() {
	fmt.Println("Handle Method Should Be rewrite")
	os.Exit(0)
}

// StopRun makes panic of USERSTOPRUN error and go to recover function if defined.
func (c *Command) StopRun() {
	panic("StopRun")
}

// ParseOptions maps input data map to obj struct.
func (c *Command) ParseOptions(options []*Option) {
	
	c.OptionValues = make(map[string]*OptionValue)
	
	for _, option := range options {
		c.OptionValues[option.Name] = NewOptionValue("", option.Value)
		flag.StringVar(&c.OptionValues[option.Name].CommandValue, option.Name, "", option.Usage)
	}

	// remove command route
	tmp := []string{}
	tmp = append(tmp, os.Args[:1]...)
	tmp = append(tmp, os.Args[3:]...)
	os.Args = tmp
	
	flag.Parse()
	
	os.Args = OsArgs
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
func (c *Command) GetStrings(key string, def ...[]string) ([]string, error) {
	var defv []string
	v := c.Query(key)
	
	if v != nil && v.CommandValue != "" {
		return strings.Split(strings.Trim(v.CommandValue, "[] "), " "), nil
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.([]string), nil
	}

	return defv, nil
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

func (c *Command) GetInts(key string, def ...[]int) ([]int, error) {
	var defv []int
	v := c.Query(key)
	
	if v != nil && v.CommandValue != "" {
		strs := strings.Split(strings.Trim(v.CommandValue, "[] "), " ")
		ints := []int{}
		if len(strs) > 0 {
			for _, str := range strs {
				if i64, err := strconv.ParseInt(str, 10, 10); err == nil {
					ints = append(ints, int(i64))
				} else {
					return ints, err
				}
			}	
		}
		
		return ints, nil
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.([]int), nil
	}

	return defv, nil
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

func (c *Command) GetInt8s(key string, def ...[]int8) ([]int8, error) {
	var defv []int8
	v := c.Query(key)
	
	if v != nil && v.CommandValue != "" {
		strs := strings.Split(strings.Trim(v.CommandValue, "[] "), " ")
		int8s := []int8{}
		if len(strs) > 0 {
			for _, str := range strs {
				if i64, err := strconv.ParseInt(str, 10, 8); err == nil {
					int8s = append(int8s, int8(i64))
				} else {
					return int8s, err
				}
			}	
		}
		
		return int8s, nil
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.([]int8), nil
	}

	return defv, nil
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

func (c *Command) GetInt16s(key string, def ...[]int16) ([]int16, error) {
	var defv []int16
	v := c.Query(key)
	
	if v != nil && v.CommandValue != "" {
		strs := strings.Split(strings.Trim(v.CommandValue, "[] "), " ")
		int16s := []int16{}
		if len(strs) > 0 {
			for _, str := range strs {
				if i64, err := strconv.ParseInt(str, 10, 16); err == nil {
					int16s = append(int16s, int16(i64))
				} else {
					return int16s, err
				}
			}	
		}
		
		return int16s, nil
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.([]int16), nil
	}

	return defv, nil
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

func (c *Command) GetInt32s(key string, def ...[]int32) ([]int32, error) {
	var defv []int32
	v := c.Query(key)
	
	if v != nil && v.CommandValue != "" {
		strs := strings.Split(strings.Trim(v.CommandValue, "[] "), " ")
		int32s := []int32{}
		if len(strs) > 0 {
			for _, str := range strs {
				if i64, err := strconv.ParseInt(str, 10, 32); err == nil {
					int32s = append(int32s, int32(i64))
				} else {
					return int32s, err
				}
			}	
		}
		
		return int32s, nil
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.([]int32), nil
	}

	return defv, nil
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

func (c *Command) GetInt64s(key string, def ...[]int64) ([]int64, error) {
	var defv []int64
	v := c.Query(key)
	
	if v != nil && v.CommandValue != "" {
		strs := strings.Split(strings.Trim(v.CommandValue, "[] "), " ")
		int64s := []int64{}
		if len(strs) > 0 {
			for _, str := range strs {
				if i64, err := strconv.ParseInt(str, 10, 64); err == nil {
					int64s = append(int64s, int64(i64))
				} else {
					return int64s, err
				}
			}	
		}
		
		return int64s, nil
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.([]int64), nil
	}

	return defv, nil
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

func (c *Command) GetBools(key string, def ...[]bool) ([]bool, error) {
	var defv []bool
	v := c.Query(key)
	
	if v != nil && v.CommandValue != "" {
		strs := strings.Split(strings.Trim(v.CommandValue, "[] "), " ")
		bools := []bool{}
		if len(strs) > 0 {
			for _, str := range strs {
				if b, err := strconv.ParseBool(str); err == nil {
					bools = append(bools, b)
				} else {
					return bools, err
				}
			}	
		}
		
		return bools, nil
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.([]bool), nil
	}

	return defv, nil
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

func (c *Command) GetFloats(key string, def ...[]float64) ([]float64, error) {
	var defv []float64
	v := c.Query(key)
	
	if v != nil && v.CommandValue != "" {
		strs := strings.Split(strings.Trim(v.CommandValue, "[] "), " ")
		float64s := []float64{}
		if len(strs) > 0 {
			for _, str := range strs {
				if f64, err := strconv.ParseFloat(str, 64); err == nil {
					float64s = append(float64s, f64)
				} else {
					return float64s, err
				}
			}	
		}
		
		return float64s, nil
	} else if len(def) > 0 {
		return def[0], nil
	} else if v.DefaultValue != nil {
		return v.DefaultValue.([]float64), nil
	}

	return defv, nil
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
		fmt.Println("Command Is Not CommandInterface")
		os.Exit(0)
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

	c, ok := p.routers[RouteName]
	if !ok {
		fmt.Println("Not Registered Route: " + RouteName)
		os.Exit(0)
	}

	c.Handle()

	if err := RemovePidFile(PidFile); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}