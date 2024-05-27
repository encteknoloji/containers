# strimzi-kafka-cruise-control-ui

A container image based on [Strimzi Kafka](https://strimzi.io/) with [cruise-control-ui](https://github.com/linkedin/cruise-control-ui)

Use following commands to build & push to Quay.io

```
export KAFKA_VERSION=0.41.0-kafka-3.7.0
export CC_UI_VERSION=0.4.0

docker build --pull --push -t quay.io/encteknoloji/kafka:${KAFKA_VERSION}-ui-${CC_UI_VERSION} --build-arg="KAFKA_VERSION=${KAFKA_VERSION}" --build-arg="CC_UI_VERSION=${CC_UI_VERSION}" .
```