#!/bin/bash

SCRIPT_DIR=$(dirname "$(realpath "$0")")
SERVER_DIR="$SCRIPT_DIR/server"

SERVER_KEY="$SERVER_DIR/server.key"
SERVER_CERT="$SERVER_DIR/server.crt"
SERVER_CSR="$SERVER_DIR/server.csr"
CONFIG_FILE="$SCRIPT_DIR/openssl.cnf"


if [ ! -d $SERVER_DIR ]
then
    mkdir $SERVER_DIR
fi

openssl genrsa -out $SERVER_KEY 4096

openssl req -new -key $SERVER_KEY -out $SERVER_CSR -config $CONFIG_FILE

openssl x509 -req -days 365 -in $SERVER_CSR -signkey $SERVER_KEY -out $SERVER_CERT -extensions v3_req -extfile $CONFIG_FILE

