all: build createConfig copy clean
build:
	go build -o go-shecan main.go
createConfig:
	sudo mkdir -p /etc/goShecan
	sudo cp configs/config.json /etc/goShecan/
copy:
	sudo cp go-shecan /usr/local/bin
clean:
	rm -rf go-shecan