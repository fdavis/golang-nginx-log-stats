# golang-nginx-log-stats

This README.md is about the docker-compose environment. For the nginx parsing program please see the folder [golangstats](golangstats)

## about

This docker-compose environment starts an nginx web server, an nginx tcp proxy,
a gatling test simulation, a golang nginx stats collector (this project), and
a statsd/grafana/graphite container.

## fire it up!

Open up two terminals and a web browser

1. In a terminal use `docker-compose up -d; docker-compose logs` to launch the containers,
and follow the logs. docker-compose up may be sufficient but if any container (such as gatling.io) exits prematurely
docker-compose tears down the whole environment.
1. In the other terminal use `docker-compose ps` to get the forwarded ports for 80 on the statsd container and the web container
1. Visit these in your browser (the ip you will need depends on your setup and where your docker-engine is running, if you are using docker-tools `docker-machine ls` should tell you the IP address otherwise you can use localhost on Linux)
1. On Grafana login with the username and password admin (that's admin for both fields). Navigate to the "http status" dashboard to see the results of the program output to graphite
1. Feel free to cause more results by visiting the web server (routes used in tests, /, /403mepls, /404mepls, /500mepls)
1. Gatling will publish results of its tests when finished in a folder like this: gatling/results/mystatuscodesim-1458704153189/index.html


Example grafana screen shot of script reporting during gatling.io tests (200 spike is me running `while :; do curl -Is 192.168.99.100:32811 > /dev/null; done` for a minute or so)

![grafana screenshot](img/grafana.png?raw=true "Grafana In Action")

## Typical development cycle

```
$ change something
docker-compose stop # if comtainers already running
docker-compose rm -fv # if containers already exist
docker-compose up -d # allows your term to detach without stopping all containers
docker-compose logs # allows you to keep tailing all the logs anyway
```

I like to make that workflow very accessible via shell alias

```bash
alias dcompcycle='docker-compose stop; docker-compose rm -fv; docker-compose up -d; docker-compose logs'
```

## To send requests to the web server on the command line

Just attach to the gatling container and curl the web or proxy host. First find the container name using `docker-compose ps`


```bash
$ docker-compose ps
            Name                           Command               State                        Ports                     
-----------------------------------------------------------------------------------------------------------------------
golangnginxlogstats_gatling_1   sh -c gatling.sh -sf /opt/ ...   Exit 1                                                 
golangnginxlogstats_proxy_1     nginx -g daemon off;             Up       443/tcp, 0.0.0.0:32785->80/tcp                
golangnginxlogstats_web_1       nginx -g daemon off; -c /e ...   Up       0.0.0.0:32783->443/tcp, 0.0.0.0:32784->80/tcp 
```

To create requests from your laptop to Nginx on the command line you first have to find the docker-engine host's IP address.
Its easier to run 

```bash
curl -sk https://192.168.99.100:32785 -H'X-Forwarded-For: 127.7.7.7,192.168.91.121'
```

## My installed versions

```bash
OSX Yosemite 10.10.5

installed go via homebrew
$ go version
go version go1.6 darwin/amd64

installed docker via docker-tools using a VBox backed docker-machine
$ docker info
Containers: 95
Images: 408
Storage Driver: aufs
 Root Dir: /mnt/sda1/var/lib/docker/aufs
 Backing Filesystem: extfs
 Dirs: 598
 Dirperm1 Supported: true
Execution Driver: native-0.2
Logging Driver: json-file
Kernel Version: 4.1.10-boot2docker
Operating System: Boot2Docker 1.8.3 (TCL 6.4); master : af8b089 - Mon Oct 12 18:56:54 UTC 2015
CPUs: 7
Total Memory: 1.955 GiB
Name: notdefault
ID: OV7L:NLXN:64BA:XYQA:MBDR:NH7Z:W2MB:W6V6:5JX6:J4G2:2AAJ:KZ6W
Debug mode (server): true
File Descriptors: 42
Goroutines: 118
System Time: 2016-03-22T05:47:29.914823369Z
EventsListeners: 0
Init SHA1: 
Init Path: /usr/local/bin/docker
Docker Root Dir: /mnt/sda1/var/lib/docker
Labels:
 provider=virtualbox
```
