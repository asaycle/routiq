FROM postgres:15

RUN apt-get update && \
    apt-get install -y postgresql-15-postgis-3 postgresql-15-postgis-3-scripts && \
    apt-get clean

COPY scripts/db/init-postgis.sh /docker-entrypoint-initdb.d/

CMD ["postgres"]