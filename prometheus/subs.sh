#!/bin/sh

# Install envsubst
wget -O envsubst https://github.com/a8m/envsubst/releases/download/v1.4.2/envsubst-Linux-x86_64
chmod +x envsubst

# Perform variable substitution on prometheus.yml
export CONSUMER_HOST="${CONSUMER_HOST:-localhost}"
envsubst < /etc/prometheus/prometheus.yml > /etc/prometheus/prometheus_substituted.yml

# Start Prometheus server with the substituted configuration
/bin/prometheus --config.file=/etc/prometheus/prometheus_substituted.yml
