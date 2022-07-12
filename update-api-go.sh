#!/bin/bash

docker stop api
docker rm api

cd go
docker build -t yourname/api .
docker run --name api -p 8082:80 -d yourname/api

docker network connect demo api
cd ..