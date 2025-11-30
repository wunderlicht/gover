package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"regexp"
	"runtime"
)

type site struct {
	url    string
	regexp string
}

func main() {
	site := site{"https://go.dev/dl/", `<a class="download downloadBox" href="\/dl\/go(\d*\.\d*\.\d*)\.src`}

	installedVersion, err := exec.Command("go", "version").Output()
	if err != nil {
		panic("Can't get installed go version " + err.Error())
	}
	fmt.Printf("%s", installedVersion)

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

	fmt.Println("Download:", getDlURL(version[1]))
}

func getDlURL(version string) (dlURL string) {
	const srcExt = ".src.tar.gz"
	osExt := map[string]string{
		"darwin":  ".pkg",
		"linux":   ".tar.gz",
		"windows": ".msi",
	}

	baseURL := "https://go.dev/dl/"
	file := "go" + version + "." + runtime.GOOS + "-" + runtime.GOARCH

	ext, ok := osExt[runtime.GOOS]
	if ok {
		dlURL = baseURL + file + ext
	} else {
		dlURL = baseURL + file + srcExt
	}

	return
}
