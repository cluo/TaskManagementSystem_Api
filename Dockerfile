FROM debian:latest

COPY ./conf.docker /task_management_api/conf
COPY ./main /task_management_api/main
ENV TZ=Asia/Shanghai
CMD ["/task_management_api/main"]

EXPOSE 80