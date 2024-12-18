## LazyRAT ##


### Intro ###
* **POC** Command & Control Server/Manager/Client using gRPC and Protocol Buffers in Golang


* Prerequisites:
    * go version 1.23.3 or later
    * protoc compiler
    * go protobuf plugins
    * unused redis instance on default port 6379 (make all deletes all keys)

--> ensure required go version is intalled then run `setup.sh` (targetting apt based system , and no WSL)

### Usage ###

* Ensure you're not using a Redis instance with unbacked keys.

* Add your server IP or DNS name to `utils/cert/openssl.cnf`.
* Start `./server/bin/server`.
* Start manager. use `help` to get command manual
* Start `asset` on however many Linux/Windows endpoints you want , should be handled.


### Project Structure ###

* Self Explanatory

### Notes ###

* It is assumed for now that `Server` and `Manager` will be running on the same host (even though multiple remote Managers are handled) , because of elements of the *TODO* section below

### TODO ###

* Add Log Propagation/Sync between Server and Managers. `download` command will download file into the Logging directory that is only available on the Server for now.

* Add Manager Authentication. maybe mTLS
 



