### status bar written in Go for i3wm

![showcase](https://raw.githubusercontent.com/lsgrep/gostatus/master/screenshot.jpg)

#### ROADMAP
* TODO support YAML config file
* TODO add necessary addons
    * Disk IO throughput
    * Network throughput
    * Volume Up/Down
    * Spotify
    * Bitcoin/Ethereum Ticker
* TODO Better Design

#### warning 
* please edit `main.go`, and change network interface & disk mount paths according to your system first.

#### install 
* `dep ensure -v`, `dep` as in `https://github.com/golang/dep`
* `go build`
* `sudo ln -s $(pwd)/gostatus /usr/local/bin/gostatus` or if you are using fish `sudo ln -s (pwd)/gostatus /usr/local/bin/gostatus`
* change status_command in  `~/.config/i3/config`


#### debugging
Just run it and read the `stdout`  . Except the first JSON object row,
 each row should be a JSON array with a leading comma(`,`), or i3bar failed to parse the line.
  

```
╰─λ ./gostatus                                                                                                                                                                                                                          fish-0 | 127 < 15:46:44
{ "version": 1, "stop_signal": 10, "cont_signal": 12, "click_events": true }[[],[{"full_text":"14 Mar 18 15:46 CST"}]          
,[{"full_text":"14 Mar 18 15:46 CST"}]                         
,[{"full_text":" 16.28%"},{"full_text":" 7.71GB / 15.58GB"},{"full_text":"14 Mar 18 15:46 CST"}]                             
,[{"full_text":" 16.28%"},{"full_text":" 7.71GB / 15.58GB"},{"full_text":"14 Mar 18 15:46 CST"}]                             
,[{"full_text":" 16.28%"},{"full_text":"10.30.7.13","color":"#00ff00"},{"full_text":" 7.71GB / 15.58GB"},{"full_text":"  / 71.18GB / 109.53GB"},{"full_text":"  /data 169.68GB / 228.23GB"},{"full_text":"14 Mar 18 15:46 CST"}]                          
,[{"full_text":" 16.28%"},{"full_text":"10.30.7.13","color":"#00ff00"},{"full_text":" 7.70GB / 15.58GB"},{"full_text":"  / 71.18GB / 109.53GB"},{"full_text":"  /data 169.68GB / 228.23GB"},{"full_text":"14 Mar 18 15:46 CST"}]                          
,[{"full_text":" 16.28%"},{"full_text":"10.30.7.13","color":"#00ff00"},{"full_text":" 7.70GB / 15.58GB"},{"full_text":"  / 71.18GB / 109.53GB"},{"full_text":"  /data 169.68GB / 228.23GB"},{"full_text":"14 Mar 18 15:46 CST"}]  
```



