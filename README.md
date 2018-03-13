### status bar written in Go for i3wm


#### warning 
* please edit `main.go`, and change network interface & disk names according to your system first.

#### install 
* `dep ensure -v`, `dep` as in `https://github.com/golang/dep`
* `go build`
* `sudo ln -s $(pwd)/gostatus /usr/local/bin/gostatus` or if you are using fish `sudo ln -s (pwd)/gostatus /usr/local/bin/gostatus`
* change status_command in  `~/.config/i3/config`


![showcase](https://raw.githubusercontent.com/lsgrep/gostatus/master/screenshot.jpg)


