# golang-nginx-log-stats

If you want the daemons to keep running you have to run docker-compose in the background and view the logs manually https://github.com/docker/compose/issues/1751
```bash
docker-compose up -d
docker-compose logs
```

To send requests to the web server on the command line

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
