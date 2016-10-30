docker kill ambassador;docker rm ambassador;


docker run -d --expose 27017 --name=task-management-mongo 211.157.146.6:5000/mongodb-enterprise
docker run -d --link task-management-mongo:mongo -p 27017:27017 --name=ambassador 211.157.146.6:5000/ambassador:latest
