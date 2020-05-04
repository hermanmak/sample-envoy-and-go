#! /bin/bash
sudo apt-get update
sudo apt-get -y upgrade
wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz
sudo tar -xvf go1.13.3.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export GOPATH=$HOME/Projects/Proj1
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

sudo apt-get install git -y


git clone https://github.com/hermanmak/sample-envoy-and-go.git
cd sample-envoy-and-go/app/src
go run main.go
EOF
