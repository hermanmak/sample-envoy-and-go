FROM envoyproxy/envoy:latest

RUN apt-get update && apt-get -q install -y \
    curl
COPY ./config ./config
CMD /usr/local/bin/envoy -c ./config/service-envoy.yaml
