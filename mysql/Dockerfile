FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=rootpassword
ENV MYSQL_DATABASE=dbo
ENV MYSQL_USER=goappuser
ENV MYSQL_PASSWORD=goapppassword

COPY ./migrations /docker-entrypoint-initdb.d
