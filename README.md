## gostatus [![CircleCI](https://circleci.com/gh/lsgrep/gostatus.svg?style=svg)](https://circleci.com/gh/lsgrep/gostatus)

> status bar written in Go for i3wm

![showcase](https://raw.githubusercontent.com/lsgrep/gostatus/master/gostatus.jpg)

#### warning 
* please edit `config.yml`, and change network interface & disk mount paths according to your system first.

#### build & install  
check out [releases](https://github.com/lsgrep/gostatus/releases) or build manually.

* `dep ensure -v`, `dep` as in `https://github.com/golang/dep`
* `go build`
* `sudo mv gostatus /usr/local/bin/gostatus`
* edit & copy/create `config.yml`
* change status_command in  `~/.config/i3/config`, e.g. `status_command gostatus --config config.yml` 

#### multiple monitor setup
* displays can be queried via `xrandr -q`

```
bar {
    # The display is connected either via HDMI or via DisplayPort
    output DP-0
    status_command gostatus --config ~/.config/i3/gostatus.yml
}

bar {
    output HDMI-0
    status_command gostatus --config ~/.config/i3/gostatus.min.yml
}
```

#### logs
* default log location `/tmp/gostatus.log`

#### debugging
Just run it and read the `stdout`. 
Except the first JSON object line, each line should be a JSON array with a leading comma(`,`), or i3bar failed to parse the line.

```
$ ./gostatus --config config.yml
{ "version": 1, "stop_signal": 10, "cont_signal": 12, "click_events": true }[[],[{"full_text":"14 Mar 18 15:46 CST"}]          
,[{"full_text":"14 Mar 18 15:46 CST"}]                         
,[{"full_text":" 16.28%"},{"full_text":" 7.71GB / 15.58GB"},{"full_text":"14 Mar 18 15:46 CST"}]                             
,[{"full_text":" 16.28%"},{"full_text":" 7.71GB / 15.58GB"},{"full_text":"14 Mar 18 15:46 CST"}]                             
,[{"full_text":" 16.28%"},{"full_text":"10.30.7.13","color":"#00ff00"},{"full_text":" 7.71GB / 15.58GB"},{"full_text":"  / 71.18GB / 109.53GB"},{"full_text":"  /data 169.68GB / 228.23GB"},{"full_text":"14 Mar 18 15:46 CST"}]                          
,[{"full_text":" 16.28%"},{"full_text":"10.30.7.13","color":"#00ff00"},{"full_text":" 7.70GB / 15.58GB"},{"full_text":"  / 71.18GB / 109.53GB"},{"full_text":"  /data 169.68GB / 228.23GB"},{"full_text":"14 Mar 18 15:46 CST"}]                          
,[{"full_text":" 16.28%"},{"full_text":"10.30.7.13","color":"#00ff00"},{"full_text":" 7.70GB / 15.58GB"},{"full_text":"  / 71.18GB / 109.53GB"},{"full_text":"  /data 169.68GB / 228.23GB"},{"full_text":"14 Mar 18 15:46 CST"}]  
```

#### Inspired by
* https://github.com/burik666/yagostatus
* https://github.com/davidscholberg/goblocks
