FROM envoyproxy/envoy:latest

# Install Go
RUN apt-get update && apt-get -q install -y curl
RUN \
  mkdir -p /goroot && \
  curl https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

COPY main.go .
RUN ls

# Set environment variables.
ENV GOROOT /goroot
ENV GOPATH /
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

# Define working directory.
RUN go run main.go
