dependencies:
	@go mod tidy


run:
	@clear
	@echo 'Running Server...'
	@go run main.go s


test:
	@go test -tags testing ./...	


build-linux:
	@GOOS=linux
	@go build -o stripe.bin main.go


build-windows:
	@GOOS=windows
	@go build -o stripe.exe main.go 

build-mac:
	@GOOS=darwin
	@go build -o stripe main.go     	