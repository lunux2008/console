package console

import (
	"os"
	"bufio"
	"github.com/astaxie/beego/utils"
)

func FileExists(filePath string) bool {
	return utils.FileExists(filePath)
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