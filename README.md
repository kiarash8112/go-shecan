# go-shecan
go-shecan is dns changer command line tool (for linux and mac) that change config.resolv file according to configs/config.json 

by default this tool change dns servers to [shecan](https://shecan.ir/).

you can use this tool for getting golang packages that cannot be download with normal net in iran

## Install
you can clone this repo and use Makefile.
```bash
git clone https://github.com/kiarash8112/go-shecan.git && cd go-shecan && make
```
this will copy configs to /dev so it can be used in any directory if you want you can change path in Makefile and Init

## Usage
You can change your dns server with these commands
```bash
go-shecan on
```
and change your config to last state with
```bash
go-shecan off
```
