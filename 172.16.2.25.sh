docker kill ambassador;docker rm ambassador;


docker kill task-management-mongo oauth-server-db ambassador;docker rm task-management-mongo oauth-server-db ambassador;

docker run -d --name=task-management-mongo 211.157.146.6:5000/mongodb-enterprise
docker run -d --name=oauth-server-db 211.157.146.6:5000/oauth-server-db
docker run -d --name=redis 211.157.146.6:5000/redis
docker run -d --link task-management-mongo:mongo --link redis:redis --link oauth-server-db:db -p 27017:27017 -p 5432:5432 -p 6379:6379 --name=ambassador 211.157.146.6:5000/ambassador:latest
