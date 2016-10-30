FROM debian:latest

COPY ./conf.docker /task_management_api/conf
COPY ./main /task_management_api/main

CMD ["/task_management_api/main"]

EXPOSE 6001