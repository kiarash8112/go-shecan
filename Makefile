all: build copy clean
build:
	go build -o go-shecan main.go
copy:
	sudo cp go-shecan /usr/local/bin
clean:
	rm -rf go-shecan