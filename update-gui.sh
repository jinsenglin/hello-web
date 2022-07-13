#!/bin/bash

docker cp gui-index.html gui:/usr/share/nginx/html/index.html
docker cp gui-cookie.html gui:/usr/share/nginx/html/cookie.html
docker cp gui-table.html gui:/usr/share/nginx/html/table.html
docker cp gui-table.json gui:/usr/share/nginx/html/table.json