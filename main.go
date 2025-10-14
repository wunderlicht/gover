package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type site struct {
	url    string
	regexp string
}

func main() {
	site := site{"https://go.dev/dl/", `<a class="download downloadBox" href="\/dl\/go(\d\.\d*\.\d*)\.src`}

	resp, err := http.Get(site.url)
	if err != nil {
		panic("Can't get url. " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("Got other status than 200 OK: " + resp.Status)
	}
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Can't read response Body. " + err.Error())
	}

	re := regexp.MustCompile(site.regexp)
	version := re.FindStringSubmatch(string(bodyByte))
	fmt.Printf("Version on go.dev: %s\n", version[1])
}
