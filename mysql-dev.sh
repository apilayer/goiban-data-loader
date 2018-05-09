#!/usr/bin/env bash
docker stop some-mysql
docker rm some-mysql
docker run --name some-mysql -p 3306:3306 -e MYSQL_DATABASE=goiban -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
