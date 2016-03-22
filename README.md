# golang-nginx-log-stats

If you want the daemons to keep running you have to run docker-compose in the background and view the logs manually https://github.com/docker/compose/issues/1751
```bash
docker-compose up -d
docker-compose logs
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

Then 


To create requests from your laptop to Nginx on the command line you first have to find the docker-engine host's IP address.
Its easier to run 

```bash
curl -sk https://192.168.99.100:32785 -H'X-Forwarded-For: 127.7.7.7,192.168.91.121'
```


## Testing changes

After changing parts of this repo run the following to try it out

```bash
docker-compose stop; docker-compose rm -f; docker-compose up -d; docker-compose logs
```

## My setup

```bash
$ go version
go version go1.6 darwin/amd64

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
