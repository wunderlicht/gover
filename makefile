dir=release
release: macosarm macosamd linuxarm linuxamd windowsamd windows386

macosarm:
	GOOS=darwin GOARCH=arm64 go build -o $(dir)/gover.darwin-arm64 .

macosamd:
	GOOS=darwin GOARCH=amd64 go build -o $(dir)/gover.darwin-amd64 .

linuxarm:
	GOOS=linux GOARCH=arm64 go build -o $(dir)/gover.linux-arm64 .

linuxamd:
	GOOS=linux GOARCH=amd64 go build -o $(dir)/gover.linux-amd64 .

windowsamd:
	GOOS=windows GOARCH=amd64 go build -o $(dir)/gover.windows-amd64.exe .

windows386:
	GOOS=windows GOARCH=386 go build -o $(dir)/gover.windows-386.exe .
