package download

import (
	"fmt"
)

var defaultFiles []string

func init() {
	defaultFiles = []string{}
	defaultFiles = append(defaultFiles, "http://cdn5.myfloridabugman.com/wp-content/uploads/2015/03/gopher-300x200.jpg")
	defaultFiles = append(defaultFiles, "https://blog.golang.org/gopher/usergroups.png")
	defaultFiles = append(defaultFiles, "https://static01.nyt.com/images/2017/01/21/opinion/20SIUweb/20SIUweb-articleLarge.jpg")
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
