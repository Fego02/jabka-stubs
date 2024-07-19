#!/bin/bash
sudo apt update
sudo apt install curl
sudo rm -rf /usr/local/go
tar -C /usr/local -xzf <(curl -OL https://go.dev/dl/go1.22.5.linux-amd64.tar.gz)
export PATH=$PATH:/usr/local/go/bin
go version