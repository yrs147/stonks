FROM prom/prometheus:latest


COPY prometheus.yml /etc/prometheus/prometheus.yml
COPY entrypoint.sh /subs.sh

ENTRYPOINT ["/subs.sh"]
CMD []

