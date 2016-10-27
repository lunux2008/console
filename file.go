package console

import (
	"os"
	"bufio"
	"io/ioutil"  
	"github.com/astaxie/beego/utils"
)

func FileExists(filePath string) bool {
	return utils.FileExists(filePath)
}

func ReadFile(filePath string) string {
    file, err := os.Open(filePath)  
    if err != nil {
		panic(err)
	}  
    defer file.Close()  
    content, err := ioutil.ReadAll(file)  

    return string(content)  
}

func WriteFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
 
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.Flush()

	return err
}

func RemoveFile(filePath string) error {
	return os.Remove(filePath)
}