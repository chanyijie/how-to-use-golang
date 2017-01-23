package downloadutils

import (
	"fmt"
)

var defaultFiles []string

func init() {
	defaultFiles := make([]string, 2)
	defaultFiles[0] = "http://cdn5.myfloridabugman.com/wp-content/uploads/2015/03/gopher-300x200.jpg"
	defaultFiles[1] = "https://blog.golang.org/gopher/usergroups.png"
}

func Add(newFile string) {
	defaultFiles = append(defaultFiles, newFile)
}

func GetDefault() []string {
	return defaultFiles
}

func PrintDefault() {
	for _, v := range defaultFiles {
		fmt.Println(v)
	}
}
