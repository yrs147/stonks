FROM confluentinc/cp-kafka:7.3.0

ENV KAFKA_BROKER_ID=1
ENV KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181
ENV KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_INTERNAL:PLAINTEXT,PLAINTEXT_EXTERNAL:PLAINTEXT
ENV KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092,PLAINTEXT_INTERNAL://localhost:9093
ENV KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092,PLAINTEXT_INTERNAL://0.0.0.0:9093
ENV KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1
ENV KAFKA_TRANSACTION_STATE_LOG_MIN_ISR=1
ENV KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR=1

CMD bash -c "/etc/confluent/docker/run & sleep 10 && kafka-topics --create --topic nifty --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1 && kafka-topics --create --topic infosys --bootstrap-server localhost:9093 --partitions 1 --replication-factor 1 && tail -f /dev/null"
