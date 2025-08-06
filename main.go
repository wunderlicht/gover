package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

type site struct {
	url    string
	regexp string
}

func main() {
	site := site{"https://go.dev/dl/", `<a class="download downloadBox" href="\/dl\/go(\d\.\d*\.\d*)\.darwin`}

	resp, err := http.Get(site.url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't get url. %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Got other status than 200 OK. %v\n", resp.StatusCode)
		os.Exit(1)
	}
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read response Body. %v\n", err)
		os.Exit(1)
	}

	re := regexp.MustCompile(site.regexp)
	version := re.FindStringSubmatch(string(bodyByte))
	fmt.Printf("Version on go.dev: %s\n", version[1])

}
