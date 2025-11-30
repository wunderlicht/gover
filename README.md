<img src="assets/Govers.webp" width="500" alt="old and grey gopher, new and shiny gopher">

[![Build linux, windows, macOS](https://github.com/wunderlicht/gover/actions/workflows/release-artifacts.yml/badge.svg)](https://github.com/wunderlicht/gover/actions/workflows/release-artifacts.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/wunderlicht/gover)](https://goreportcard.com/report/github.com/wunderlicht/gover)

# gover
Quickly check for a new go version on go.dev. Plain and simple.
For doing that it needs a working internet connection.

As a convenience `gover` displays the download link for the new version
tailored to the current system and architecture.

Special thanks ❤️ to the wonderful human being contributing the gopher drawing to this (micro) project.

## Usage
```
> gover
go version go1.25.4 darwin/arm64
Version on go.dev: 1.25.4
Download: https://go.dev/dl/go1.25.4.darwin-arm64.pkg
```
Example shows `gover` running on MacOS on an M-chip

## How it works
1. display `go version`
2. Load https://go.dev/dl/
3. perform a regex search for the newest version number
4. get my process' GOOS and GOARCH
5. display the the download link
6. done

## Get it
Download the latest built executables or build it yourself.

When you get the executables on macOS you need to run
`xattr -d com.apple.quarantine <Binary>` to start it on the command line.