docker kill task-management-api oauth-server-service task-management-frontend ambassador-mongo ambassador-oauthdb;docker rm task-management-api oauth-server-service task-management-frontend ambassador;
docker rmi 211.157.146.6:5000/task-management-api 211.157.146.6:5000/task-management-frontend 211.157.146.6:5000/oauth-server-service 211.157.146.6:5000/ambassador 

docker run -d --name=ambassador --expose 27017 --expose 5432 --expose 4379 -e MONGO_PORT_27017_TCP=tcp://172.16.2.25:27017  -e DB_PORT_5432_TCP=tcp://172.16.2.25:5432 -e REDIS_PORT_4379_TCP=tcp://172.16.2.25:4379 211.157.146.6:5000/ambassador:latest
docker run -d -p 6009:80 --name=task-management-frontend 211.157.146.6:5000/task-management-frontend
docker run -d -p 6002:80 --link ambassador:db --name=oauth-server-service 211.157.146.6:5000/oauth-server-service
docker run -d -p 6001:6001 --link ambassador:mongo --link ambassador:redis --link oauth-server-service:oauth --name=task-management-api 211.157.146.6:5000/task-management-api


docker run -it -p 6002:80 --link ambassador:db --name=oauth-server-service 211.157.146.6:5000/oauth-server-service -bash

docker run -it --link ambassador:db --name=oauth-server-service 211.157.146.6:5000/oauth-server-service -init
docker rm oauth-server-service

service docker restart
docker restart ambassador-mongo task-management-api task-management-frontend

#!/bin/bash
docker kill oauth-server-service;docker rm oauth-server-service
docker rmi 211.157.146.6:5000/oauth-server-service
docker run -d -p 6002:80 --link ambassador:db --name=oauth-server-service 211.157.146.6:5000/oauth-server-service


http://oauth.hisign.top:6002/o/authorize/?response_type=code&client_id=hAHln3ZKrnPf8odTUdkizuSSbIP3CvRzNY0zBZXD