// Show some content that was found in each URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		normalizedUrl, prefix := url, "http://"
		if !strings.HasPrefix(url, prefix) {
			normalizedUrl = prefix + url
		}
		resp, err := http.Get(normalizedUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		status := resp.Status
		fmt.Printf("Status code %s\n", status)
		// b, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch - reading %s: %v\n", url, err)
			os.Exit(1)
		}

		// fmt.Printf("%s\n", b)
	}
}
