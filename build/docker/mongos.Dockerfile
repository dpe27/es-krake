FROM mongo:latest

RUN apt-get update && apt-get -q install -y

COPY --chown=mongodb:mongodb ../../scripts/mongo/mongos-start.sh /usr/local/bin/
COPY --chown=mongodb:mongodb ../../scripts/mongo/init-cluster.sh /usr/local/bin/

RUN chmod u+x /usr/local/bin/mongos-start.sh /usr/local/bin/init-cluster.sh

ENTRYPOINT ["mongos-start.sh"]

CMD ["mongos", "--port", "27017", "--bind_ip_all"]
