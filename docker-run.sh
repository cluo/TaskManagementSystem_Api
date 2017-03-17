#!/bin/bash
docker kill api;docker rm api;
docker rmi api-image 211.157.146.6:5000/task-management-api
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build *.go
docker restart TM-mongo
docker build -t api-image .
docker tag api-image 211.157.146.6:5000/task-management-api
docker push 211.157.146.6:5000/task-management-api
docker rmi api-image
