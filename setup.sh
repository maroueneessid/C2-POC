#!/bin/sh

sudo apt-get update -y

## Install Redis ##
sudo apt-get install lsb-release curl gpg -y 
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
sudo chmod 644 /usr/share/keyrings/redis-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list
sudo apt-get update --allow-unauthenticated --allow-insecure-repositories
sudo apt-get install redis -y

sudo systemctl enable redis-server
sudo systemctl start redis-server

## Install Proto Compiler ##

sudo apt install -y protobuf-compiler

## Install Go plugins for ProtBuf ##

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest




