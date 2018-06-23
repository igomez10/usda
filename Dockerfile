FROM mysql:5.7

ENV MYSQL_ROOT_PASSWORD mypassword
EXPOSE 3306
COPY setup.sql /docker-entrypoint-initdb.d/setup.sql

