#!/bin/sh

# wait until graphite is responding
while ! curl -sI --fail localhost:8000 > /dev/null; do
  sleep 1s
done
# wait until grafana is reponding
while ! curl -s --fail localhost > /dev/null; do
  sleep 1s
done
# wait a little longer because we keep receiving bad gateway
sleep 2s

# load graphite data backend
# https://github.com/grafana/grafana/issues/1789
# try again because it keeps failing
for x in {1..3}; do
  curl 'http://admin:admin@127.0.0.1:80/api/datasources' -X POST \
    -H 'Content-Type: application/json;charset=UTF-8' \
    --data-binary '{"name":"localGraphite","type":"graphite","url":"http://127.0.0.1:8000","access":"proxy","isDefault":true,"database":"statsd"}'
  sleep 0.5s
done
