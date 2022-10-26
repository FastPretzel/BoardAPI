FROM postgres:14.2-alpine3.15

COPY ./database/*.sql /docker-entrypoint-initdb.d/
