# sauron3
third approach to Sauron - a real time eye on your network - this time in Go Language 

![demo](https://sc-cdn.scaleengine.net/i/22e315caecf77506e50be0619e57303e.png)

## features
* defined in config hosts have ports (tcp/~~udp~~/icmp) that are pinged on demand and displayed in self refreshing webpage - ideal to be put on plasma tv in Networks Operation Command Center or other similar facilites or accessed from mobile phone 
* if you provide location href `2col` (like `http://sauron/#2col`) sauron will switch to 2 column mode (experimental)


## run
 - config.yml file 
 - browser -> by default http://localhost:8888
```
./sauron3 --help
usage: sauron3 [-h|--help] [-c|--config "<value>"] [-p|--port "<value>"] [-b|--bind "<value>"]

               a real time eye on your network

Arguments:

   -h   --help      Print help information
   -c   --config    Path to config file (default is ./config.yml)
   -p   --port      Port for webgui (default is 8888)
   -b   --bind      IP to bind (default is 0.0.0.0)
```

## build - dependencies 
 - `go get github.com/danielskowronski/sauron3_core` # <- core code now in separate package
 - `go get gopkg.in/yaml.v2`
 - `go get golang.org/x/net/icmp`
 - `go get golang.org/x/net/ipv4`
 - `go get github.com/akamensky/argparse`

## build - compilation
`go generate && go build -o sauron.o && ./sauron.o` 

## running
 - `sysctl -w net.ipv4.ping_group_range="0 65535"` - setup Linux OS for using ping
 - `setcap cap_net_raw+ep /path/to/bin` - give capability to use raw sockets
 - `CapabilityBoundingSet=CAP_NET_RAW` - systemd 