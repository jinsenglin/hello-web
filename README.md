```
docker run --name nginx -p 8080:80 -d nginx
docker cp default.conf nginx:/etc/nginx/conf.d/default.conf
docker exec nginx nginx -s reload

docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=demo -d mysql
docker exec mysql mysql -uroot -pdemo -e 'create database demo;'
docker exec mysql mysql -uroot -pdemo -e 'create table demo.demo (id int, name varchar(255));'
docker exec mysql mysql -uroot -pdemo -e 'insert into demo.demo (id, name) values (1, "cclin");'
docker exec mysql mysql -uroot -pdemo -e 'select * from demo.demo;'

docker run --name gui -p 8081:80 -d nginx
docker cp gui-index.html gui:/usr/share/nginx/html/index.html

docker run --name api -p 8082:80 -d nginx
docker cp api-index.html api:/usr/share/nginx/html/index.html

docker network create demo
docker network connect demo nginx
docker network connect demo mysql
docker network connect demo gui
docker network connect demo api

curl localhost:8080
curl localhost:8080/gui
curl localhost:8080/api
open http://localhost:8080/gui
```

```
docker stop api
docker rm api

cd go
docker build -t yourname/api .
docker run --name api -p 8082:80 -d yourname/api

docker network connect demo api
```

```
docker stop api
docker rm api

cd py
docker build -t yourname/api .
docker run --name api -p 8082:80 -d yourname/api

docker network connect demo api
```