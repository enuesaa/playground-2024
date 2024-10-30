FROM ubuntu:latest

RUN apt-get update && \
    apt-get install -y wget nginx unzip

COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
RUN mkdir -p /var/log/nginx

RUN wget "https://github.com/grafana/loki/releases/download/v2.7.0/promtail-linux-arm.zip" \
    && unzip promtail-linux-arm.zip \
    && mv promtail-linux-arm /usr/local/bin/promtail \
    && chmod +x /usr/local/bin/promtail

COPY nginx/promtail.yml /etc/promtail.yml

CMD ["sh", "-c", "nginx && promtail -config.file=/etc/promtail.yml"]
