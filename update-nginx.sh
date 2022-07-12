#!/bin/bash

docker cp default.conf nginx:/etc/nginx/conf.d/default.conf
docker exec nginx nginx -s reload