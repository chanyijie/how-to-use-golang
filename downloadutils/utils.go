package downloadutils

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

// DownloadHttpFiles using the http lib to download the files which are represented as plain http link.
// It accepts the array of http url strings for target files.
// It returns the code and any download error encountered.
func DownloadHttpFiles(urls []string) (int, error) {
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			tokens := strings.Split(url, "/")
			fileName := tokens[len(tokens)-1]
			fmt.Println("Downloading", url, "to", fileName)

			// use the last token of target url as the name of downloaded file
			output, err := os.Create(fileName)
			if err != nil {
				log.Fatal("Error while creating", fileName, "-", err)
			}
			defer output.Close()

			// handle https request
			// create new http transport
			transPort := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			// create new http client with the defined transport above
			client := &http.Client{Transport: transPort}
			res, err := client.Get(url)
			if err != nil {
				log.Fatal("http get error: ", err)
			} else {
				defer res.Body.Close()
				_, err = io.Copy(output, res.Body)
				if err != nil {
					log.Fatal("Error while downloading", url, "-", err)
				} else {
					fmt.Println("Downloaded", fileName)
				}
			}
		}(url)
	}
	wg.Wait()
	fmt.Println("Done")
	return 0, nil
}
