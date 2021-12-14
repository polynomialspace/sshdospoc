all: windows darwin linux

windows:
	GOOS=windows GOARCH=amd64 go build -o sshdos.exe
linux:
	GOOS=linux GOARCH=amd64 go build -o sshdos_linux
darwin:
	GOOS=darwin GOARCH=amd64 go build -o sshdos_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o sshdos_darwin_arm64
	lipo -create -output sshdos_darwin sshdos_darwin_amd64 sshdos_darwin_arm64
	rm sshdos_darwin_amd64 sshdos_darwin_arm64
clean:
	rm sshdos.exe sshdos_linux sshdos_darwin
