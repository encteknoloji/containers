ARG KAFKA_VERSION=0.41.0-kafka-3.7.0

FROM quay.io/strimzi/kafka:${KAFKA_VERSION}

ARG CC_UI_VERSION=0.4.0

ENV CC_UI_VER=${CC_UI_VERSION}
ENV CC_UI_HOME=./cruise-control-ui/dist/

USER root

RUN mkdir -p ${CC_UI_HOME}; \
    curl -LO https://github.com/linkedin/cruise-control-ui/releases/download/v${CC_UI_VER}/cruise-control-ui-${CC_UI_VER}.tar.gz; \
    tar xvfz cruise-control-ui-${CC_UI_VER}.tar.gz -C ${CC_UI_HOME} --strip-components=2; \
    rm -f cruise-control-ui-${CC_UI_VER}.tar.gz*; \
    echo "kafka,kafka,/kafkacruisecontrol/" > "${CC_UI_HOME}static/config.csv";

USER 1001