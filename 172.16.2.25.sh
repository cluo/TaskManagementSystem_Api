docker kill ambassador;docker rm ambassador;


# docker kill redis task-management-mongo oauth-server-db ambassador; \
#     docker rm redis task-management-mongo oauth-server-db ambassador;

# docker run -d --name=task-management-mongo 211.157.146.6:5000/mongodb-enterprise
# docker run -d --name=oauth-server-db 211.157.146.6:5000/oauth-server-db
# docker run -d --name=redis 211.157.146.6:5000/redis
# docker run -d --link task-management-mongo:mongo \
#     --link redis:redis --link oauth-server-db:db \
#     -p 27017:27017 -p 5432:5432 -p 6379:6379 \
#     --name=ambassador 211.157.146.6:5000/ambassador:latest

docker rmi 211.157.146.6:5000/mongodb-enterprise

docker kill redis task-management-mongo ambassador; \
    docker rm redis task-management-mongo ambassador;

docker run -d -e TZ="Asia/Shanghai" --name=task-management-mongo 211.157.146.6:5000/mongodb-enterprise
docker run -d -e TZ="Asia/Shanghai" --name=redis 211.157.146.6:5000/redis
docker run -d -e TZ="Asia/Shanghai" --link task-management-mongo:mongo --link redis:redis \
    -p 27017:27017 -p 6379:6379 \
    --name=ambassador 211.157.146.6:5000/ambassador:latest

docker kill api2 ambassador; \
    docker rm api2 ambassador;
docker rmi 211.157.146.6:5000/task-management-api;
docker run -d  -e TZ="Asia/Shanghai" --link task-management-mongo:mongo --link redis:redis --name=api2 211.157.146.6:5000/task-management-api
docker run -d  -e TZ="Asia/Shanghai" --link task-management-mongo:mongo --link redis:redis  --link api2:api2 \
    -p 27017:27017 -p 6379:6379  -p 80:80 \
    --name=ambassador 211.157.146.6:5000/ambassador:latest

