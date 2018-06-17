FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD mypassword
EXPOSE 3306
COPY setup.sql /docker-entrypoint-initdb.d/setup.sql

