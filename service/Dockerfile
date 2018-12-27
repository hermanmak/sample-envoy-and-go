FROM envoyproxy/envoy:latest

COPY ./config ./config
CMD /usr/local/bin/envoy -c ./config/service-envoy.yaml
