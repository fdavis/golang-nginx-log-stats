#!/bin/sh
sleep 1s
# load graphite data backend
# https://github.com/grafana/grafana/issues/1789
curl 'http://admin:admin@127.0.0.1:80/api/datasources' -X POST \
    -H 'Content-Type: application/json;charset=UTF-8' \
    --data-binary '{"name":"localGraphite","type":"graphite","url":"http://127.0.0.1:8000","access":"proxy","isDefault":true,"database":"statsd"}'
