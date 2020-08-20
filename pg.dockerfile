FROM postgres:12.4-alpine

COPY pg/init.sh /docker-entrypoint-initdb.d/00_init.sh
COPY pg/*.sql /docker-entrypoint-initdb.d/