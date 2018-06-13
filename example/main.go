package main

import (
	"fmt"
	du "github.com/nevermosby/how-to-use-golang/downloadutils"
)

func main() {
	// targetfiles := []string{
	//	"http://cdn5.myfloridabugman.com/wp-content/uploads/2015/03/gopher-300x200.jpg",
	//	"https://blog.golang.org/gopher/usergroups.png",
	// }
	
	du.PrintDefault()
	_, err := du.DownloadHttpFiles(du.GetDefault())
	if err != nil {
		fmt.Println("Got error:", err)
	}
	fmt.Println("Done")
}
