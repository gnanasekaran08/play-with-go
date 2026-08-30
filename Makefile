VERSION = 1.0.0
hello:
	echo "Its not working!"
hi:
	echo "Its working!"
run: 
	go run .
compile:
	echo "Compiling for every platform"
	GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows_amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin_amd64 main.go


