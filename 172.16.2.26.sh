# docker kill task-nginx api1 api2 api3 oauth1 oauth2 oauth3 frontend1 frontend2 frontend3 ambassador; \
#     docker rm task-nginx api1 api2 api3 oauth1 oauth2 oauth3 frontend1 frontend2 frontend3 ambassador;
# docker rmi 211.157.146.6:5000/task-management-api \
#     211.157.146.6:5000/task-management-frontend \
#     211.157.146.6:5000/oauth-server-service \
#     211.157.146.6:5000/task-nginx

# docker run -d --name=ambassador --expose 27017 --expose 5432 --expose 6379 -e MONGO_PORT_27017_TCP=tcp://172.16.2.25:27017  -e DB_PORT_5432_TCP=tcp://172.16.2.25:5432 -e REDIS_PORT_6379_TCP=tcp://172.16.2.25:6379 211.157.146.6:5000/ambassador:latest
# docker run -d --name=frontend1 211.157.146.6:5000/task-management-frontend
# docker run -d --name=frontend2 211.157.146.6:5000/task-management-frontend
# docker run -d --name=frontend3 211.157.146.6:5000/task-management-frontend
# docker run -d --link ambassador:db --name=oauth1 211.157.146.6:5000/oauth-server-service
# docker run -d --link ambassador:db --name=oauth2 211.157.146.6:5000/oauth-server-service
# docker run -d --link ambassador:db --name=oauth3 211.157.146.6:5000/oauth-server-service
# docker run -d --link ambassador:mongo --link ambassador:redis --name=api1 211.157.146.6:5000/task-management-api
# docker run -d --link ambassador:mongo --link ambassador:redis --name=api2 211.157.146.6:5000/task-management-api
# docker run -d --link ambassador:mongo --link ambassador:redis --name=api3 211.157.146.6:5000/task-management-api
# docker run -d --name=task-nginx -p 6001:6001 -p 6002:6002 -p 6009:6009\
#     --link frontend1:frontend1 --link frontend2:frontend2 --link frontend3:frontend3 \
#     --link api1:api1 --link api2:api2 --link api3:api3 \
#     --link oauth1:oauth1 --link oauth2:oauth2 --link oauth3:oauth3 \
#     211.157.146.6:5000/task-nginx



docker kill task-nginx ambassador api1 api2 api3 frontend1 frontend2 frontend3; \
    docker rm task-nginx ambassador api1 api2 api3 frontend1 frontend2 frontend3;
docker rmi 211.157.146.6:5000/task-management-api \
    211.157.146.6:5000/task-management-frontend \
    211.157.146.6:5000/oauth-server-service \
    211.157.146.6:5000/task-nginx
docker run -d --name=ambassador --expose 27017 --expose 6379 \
    -e MONGO_PORT_27017_TCP=tcp://172.16.2.25:27017 \
    -e REDIS_PORT_6379_TCP=tcp://172.16.2.25:6379 \
    211.157.146.6:5000/ambassador:latest
docker run -d --name=frontend1 211.157.146.6:5000/task-management-frontend
docker run -d --link ambassador:mongo --link ambassador:redis --name=api1 211.157.146.6:5000/task-management-api
docker run -d --name=task-nginx -p 6001:6001 -p 6009:6009 \
    --link frontend1:frontend1 --link api1:api1 \
    211.157.146.6:5000/task-nginx


docker run -d --name=ambassador --expose 80 --expose 27017 --expose 6379 \
    -e MONGO_PORT_27017_TCP=tcp://172.16.2.25:27017 \
    -e REDIS_PORT_6379_TCP=tcp://172.16.2.25:6379 \
    -e API2_PORT_80_TCP=tcp://172.16.2.25:80 \
    211.157.146.6:5000/ambassador:latest
docker run -d --name=task-nginx -p 6001:6001 -p 6009:6009 \
    --link frontend1:frontend1 --link api1:api1 --link ambassador:api2 \
    211.157.146.6:5000/task-nginx
# docker run -d --name=task-nginx -p 6001:6001 -p 6009:6009\
    # --link frontend1:frontend1 --link frontend2:frontend2 --link frontend3:frontend3 \
    # --link api1:api1 --link api2:api2 --link api3:api3 \
    # 211.157.146.6:5000/task-nginx















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