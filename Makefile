all: build createConfig copy clean
build:
	go build -o go-shecan main.go
createConfig:
	sudo mkdir /dev/goShecan
	sudo cp configs/config.json /dev/goShecan/
copy:
	sudo cp go-shecan /usr/local/bin
clean:
	rm -rf go-shecan