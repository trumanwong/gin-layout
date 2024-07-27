FROM docker.io/debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        supervisor  \
        libc6 \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

WORKDIR /services

COPY cmd/server/chat-roleplay-api-server /services/server
COPY init.sh /services/init.sh
COPY supervisor /etc/supervisor

EXPOSE 8000

CMD ["/bin/bash", "/services/init.sh"]