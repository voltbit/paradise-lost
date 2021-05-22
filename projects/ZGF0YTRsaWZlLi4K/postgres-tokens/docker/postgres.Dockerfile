FROM postgres:13.3

COPY migrations/*.sql /docker-entrypoint-initdb.d/
