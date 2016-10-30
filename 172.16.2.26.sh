docker kill task-management-api task-management-frontend ambassador-mongo;docker rm task-management-api task-management-frontend ambassador-mongo;
docker rmi 211.157.146.6:5000/task-management-api:0.01 211.157.146.6:5000/task-management-frontend:0.01 211.157.146.6:5000/ambassador:latest


docker run -d -p 6009:80 --name=task-management-frontend 211.157.146.6:5000/task-management-frontend:0.01
docker run -d --name=ambassador-mongo --expose 27017 -e MONGO_PORT_27017_TCP=tcp://172.16.2.25:27017 211.157.146.6:5000/ambassador:latest
docker run -d -p 6001:6001 --link ambassador-mongo:mongo --name=task-management-api 211.157.146.6:5000/task-management-api:0.01
