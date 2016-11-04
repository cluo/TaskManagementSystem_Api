docker kill ambassador-mongo ambassador-oauth;docker rm ambassador-mongo ambassador-oauth;


docker kill task-management-mongo oauth-server-db ambassador-mongo ambassador-oauth;docker rm task-management-mongo oauth-server-db ambassador-mongo ambassador-oauth;

docker run -d --name=task-management-mongo 211.157.146.6:5000/mongodb-enterprise
docker run -d --name=oauth-server-db 211.157.146.6:5000/oauth-server-db
docker run -d --link task-management-mongo:mongo -p 27017:27017 --name=ambassador-mongo 211.157.146.6:5000/ambassador:latest
docker run -d --link oauth-server-db:db -p 5432:5432 --name=ambassador-oauth 211.157.146.6:5000/ambassador:latest
